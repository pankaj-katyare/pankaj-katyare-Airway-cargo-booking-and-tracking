package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/db"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
)

type (
	BookingTask interface {
		Insert(ctx context.Context, booking_task model.BookingTask) error
		Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) (model.BookingTask, error)
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

func (s *SQL) Insert(ctx context.Context, BookingTask model.BookingTask) error {
	query := `insert into booking_task (id, name) 
	values (?,?,?,?,?);`

	_, err := s.db.ExecContext(ctx, query, BookingTask.Id, BookingTask.BookingId, BookingTask.Status, BookingTask.CreatedAt, BookingTask.CompletedAt)
	return err
}

func (s *SQL) Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) ([]*model.BookingTask, error) {
	var booking_task []*model.BookingTask
	mainQuery := db.SelectBuilder("booking_task", where, orderBy, groupBy)
	err := s.db.GetContext(ctx, &booking_task, mainQuery)
	if err != nil {
		return nil, err
	}
	return booking_task, nil
}

func (s *SQL) Update(ctx context.Context, set map[string]interface{}, condition []string) (int64, error) {
	if len(set) <= 0 {
		return 0, errors.New("no data to update in update call")
	}
	query := db.UpdateBuilder("booking_task", set, condition)
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
	query := db.DeleteBuilder("booking_task", condition)
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
