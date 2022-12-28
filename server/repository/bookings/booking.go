package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/db"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
)

type (
	Booking interface {
		Insert(ctx context.Context, booking model.Booking) error
		Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) (model.Booking, error)
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

func (s *SQL) Insert(ctx context.Context, Booking model.Booking) error {
	query := `insert into booking (id, boking_request_id, status, customer_id, task_id, quote_id, milestone_id, liner_id, source, destination, city) 
	values (?,?,?,?,?,?,?,?,?,?,?,?);`

	_, err := s.db.ExecContext(ctx, query, Booking.Id, Booking.BookingRequestId, Booking.CustomerId, Booking.TaskId, Booking.QuoteId, Booking.MilestoneId, Booking.LinerId, Booking.Source, Booking.Destination, Booking.City)
	return err
}

func (s *SQL) Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) ([]*model.Booking, error) {
	var bookings []*model.Booking
	mainQuery := db.SelectBuilder("booking", where, orderBy, groupBy)
	err := s.db.GetContext(ctx, &bookings, mainQuery)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *SQL) Update(ctx context.Context, set map[string]interface{}, condition []string) (int64, error) {
	if len(set) <= 0 {
		return 0, errors.New("no data to update in update call")
	}
	query := db.UpdateBuilder("booking", set, condition)
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
	query := db.DeleteBuilder("booking", condition)
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
