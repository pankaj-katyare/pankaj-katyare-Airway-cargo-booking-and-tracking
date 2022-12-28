package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/db"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
)

type (
	Task interface {
		Insert(ctx context.Context, tasks model.Task) error
		Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) (model.Task, error)
		Update(ctx context.Context, set map[string]interface{}, condition []string) (int64, error)
		Delete(ctx context.Context, condition []string) (int64, error)
	}
)

type SQL struct {
	db *sqlx.DB
}
type Config struct {
	DB *sqlx.DB
}

func NewConnection(config *Config) *SQL {
	return &SQL{
		db: config.DB,
	}
}

func (s *SQL) Insert(ctx context.Context, Task model.Task) error {
	query := `insert into tasks (id, name) 
	values (?,?);`

	_, err := s.db.ExecContext(ctx, query, Task.Id, Task.Name)
	return err
}

func (s *SQL) Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) ([]*model.Task, error) {
	var tasks []*model.Task
	mainQuery := db.SelectBuilder("tasks", where, orderBy, groupBy)
	err := s.db.GetContext(ctx, &tasks, mainQuery)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *SQL) Update(ctx context.Context, set map[string]interface{}, condition []string) (int64, error) {
	if len(set) <= 0 {
		return 0, errors.New("no data to update in update call")
	}
	query := db.UpdateBuilder("tasks", set, condition)
	result, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *SQL) Delete(ctx context.Context, condition []string) (int64, error) {
	query := db.DeleteBuilder("tasks", condition)
	result, err := s.db.ExecContext(ctx, query)

	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
