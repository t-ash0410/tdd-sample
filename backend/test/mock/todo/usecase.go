package todo

import (
	"errors"
	"time"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
)

type SuccessListUsecase struct{}

func (u SuccessListUsecase) Handle(result *[]entities.Task) error {
	*result = append(*result, entities.Task{
		Id:          "",
		Name:        "",
		Description: "",
		UpdatedAt:   time.Now(),
	})
	return nil
}

type EmptyListUsecase struct{}

func (u EmptyListUsecase) Handle(result *[]entities.Task) error {
	return nil
}

type FailListUsecase struct{}

func (u FailListUsecase) Handle(result *[]entities.Task) error {
	return errors.New("some error")
}

type SuccessAddUsecase struct{}

func (u SuccessAddUsecase) Handle(name string, description string) error {
	return nil
}

type FailAddUsecase struct{}

func (u FailAddUsecase) Handle(name string, description string) error {
	return errors.New("some error")
}
