//go:build integrate
// +build integrate

package repositories_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/pkg/errors"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/interfaces"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/repositories"
	"github.com/t-ash0410/tdd-sample/backend/pkg/infrastructure"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	"google.golang.org/grpc/codes"
)

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	m.Run()
}

var repository interfaces.IRepository
var ctx context.Context
var cancel context.CancelFunc
var spannerContext *infrastructure.SpannerContext
var connectionString string

func setup() {
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Minute)
	connectionString = os.Getenv("SPANNER_CONNECTION_STRING")
	spannerContext = infrastructure.NewSpannerContext(ctx, connectionString)
	migration()
	repository = repositories.NewRepository(spannerContext)
}

func migration() {
	// split connection strings
	matches := regexp.MustCompile("^(.*)/databases/(.*)$").FindStringSubmatch(connectionString)
	if matches == nil || len(matches) != 3 {
		log.Fatal(errors.New("invalid db connection string"))
	}

	//drop database if exists this instance and project
	db, err := spannerContext.AdminClient.GetDatabase(ctx, &adminpb.GetDatabaseRequest{Name: connectionString})
	if err != nil {
		log.Fatalf("get database %+v", err)
	}
	if db != nil {
		if err := spannerContext.AdminClient.DropDatabase(ctx, &adminpb.DropDatabaseRequest{Database: connectionString}); err != nil {
			log.Fatalf("drop database %+v", err)
		}
	}

	//create database
	op, err := spannerContext.AdminClient.CreateDatabase(ctx, &adminpb.CreateDatabaseRequest{
		Parent:          matches[1],
		CreateStatement: "CREATE DATABASE `" + matches[2] + "`",
		ExtraStatements: []string{
			`CREATE TABLE todo (
					todo_id 					STRING(1024) NOT NULL,
					name 							STRING(1024) NOT NULL,
					description 			STRING(1024) NOT NULL,
					last_update_time 	TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true)
			) PRIMARY KEY (todo_id)`,
		},
	})
	if err != nil {
		log.Fatalf("execute create database %+v", err)
	}
	if _, err := op.Wait(ctx); err != nil {
		log.Fatalf("wait create database %+v", err)
	}
}

func useErrorCtx(test func()) {
	errorCtx, errorCtxCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer errorCtxCancel()
	errorSpannerContext := infrastructure.NewSpannerContext(errorCtx, connectionString)
	errorSpannerContext.Close()
	repository = repositories.NewRepository(errorSpannerContext)
	test()
	repository = repositories.NewRepository(spannerContext)
}

func teardown() {
	cancel()
	spannerContext.Close()
}

func TestRepository_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tasks := []entities.Task{}
		if err := repository.List(&tasks); err != nil {
			t.Error(err)
		}
		task := entities.Task{
			Id:          "id",
			Name:        "name",
			Description: "desc",
			UpdatedAt:   time.Now(),
		}
		if err := repository.Add(task); err != nil {
			t.Fatalf("execute add data usecase %+v", err)
		}
		if err := repository.List(&tasks); err != nil {
			t.Error(err)
		}
	})
	t.Run("internal server error", func(t *testing.T) {
		useErrorCtx(func() {
			tasks := []entities.Task{}
			if err := repository.List(&tasks); err == nil {
				t.Error("use closed context.")
			}
		})
	})
	t.Run("invalid column defs", func(t *testing.T) {
		// add data
		task := entities.Task{
			Id:          "id",
			Name:        "name",
			Description: "desc",
			UpdatedAt:   time.Now(),
		}
		if err := repository.Add(task); err != nil {
			t.Fatalf("execute add data usecase %+v", err)
		}

		//drop column
		op, err := spannerContext.AdminClient.UpdateDatabaseDdl(ctx, &adminpb.UpdateDatabaseDdlRequest{
			Database:   connectionString,
			Statements: []string{"ALTER TABLE todo ALTER COLUMN description BYTES(MAX)"},
		})
		if err != nil {
			t.Fatalf("execute alter table drop column %+v", err)
		}
		if err := op.Wait(ctx); err != nil {
			t.Fatalf("wait alter table drop column %+v", err)
		}

		//execute test
		tasks := []entities.Task{}
		if err := repository.List(&tasks); err == nil {
			t.Fatal("use not enough column table.")
		} else {
			code := spanner.ErrCode(err)
			if code != codes.InvalidArgument {
				t.Fatalf("unexpected error. %+v", err)
			}
		}
	})
	migration()
}

func TestRepository_Add(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		task := entities.Task{
			Id:          "id",
			Name:        "name",
			Description: "desc",
			UpdatedAt:   time.Now(),
		}
		if err := repository.Add(task); err != nil {
			t.Error(err)
		}
	})
	t.Run("invalid data", func(t *testing.T) {
		task := entities.Task{}
		for i := 0; i < 10000; i++ {
			task.Id += fmt.Sprintf("%d", i)
		}
		if err := repository.Add(task); err == nil {
			t.Error("registerd invalid data")
		}
	})
}
