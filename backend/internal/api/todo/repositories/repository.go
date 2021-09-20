package repositories

import (
	"time"

	"cloud.google.com/go/spanner"
	"github.com/pkg/errors"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/interfaces"
	"github.com/t-ash0410/tdd-sample/backend/pkg/infrastructure"
	"google.golang.org/api/iterator"
)

type Repository struct {
	ctx        *infrastructure.SpannerContext
	tableName  string
	columnDefs []string
}

func NewRepository(ctx *infrastructure.SpannerContext) interfaces.IRepository {
	return &Repository{
		ctx:       ctx,
		tableName: "todo",
		columnDefs: []string{
			"todo_id",
			"name",
			"description",
			"last_update_time",
		},
	}
}

func (repo *Repository) List(result *[]entities.Task) error {
	iter := repo.ctx.DataClient.Single().Read(*repo.ctx.BaseContext, repo.tableName, spanner.AllKeys(), repo.columnDefs)
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		var id, name, description string
		var lastUpdateTime time.Time
		if err := row.Columns(&id, &name, &description, &lastUpdateTime); err != nil {
			return err
		}

		*result = append(*result, entities.Task{
			Id:          id,
			Name:        name,
			Description: description,
			UpdatedAt:   lastUpdateTime,
		})
	}
	return nil
}

func (repo *Repository) Add(task entities.Task) error {
	m := []*spanner.Mutation{
		spanner.InsertOrUpdate(repo.tableName, repo.columnDefs, []interface{}{task.Id, task.Name, task.Description, task.UpdatedAt}),
	}
	if _, err := repo.ctx.DataClient.Apply(*repo.ctx.BaseContext, m); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
