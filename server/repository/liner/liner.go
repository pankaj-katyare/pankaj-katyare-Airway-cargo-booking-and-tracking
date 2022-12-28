package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/db"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
)

type (
	Liner interface {
		Insert(ctx context.Context, liners model.Liner) error
		Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) (model.Liner, error)
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

func (s *SQL) Insert(ctx context.Context, Liner model.Liner) error {
	query := `insert into liners (id, name) 
	values (?,?);`

	_, err := s.db.ExecContext(ctx, query, Liner.Id, Liner.Name)
	return err
}

func (s *SQL) Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) ([]*model.Liner, error) {
	var liners []*model.Liner
	mainQuery := db.SelectBuilder("liners", where, orderBy, groupBy)
	err := s.db.GetContext(ctx, &liners, mainQuery)
	if err != nil {
		return nil, err
	}
	return liners, nil
}

func (s *SQL) Update(ctx context.Context, set map[string]interface{}, condition []string) (int64, error) {
	if len(set) <= 0 {
		return 0, errors.New("no data to update in update call")
	}
	query := db.UpdateBuilder("liners", set, condition)
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
	query := db.DeleteBuilder("liners", condition)
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
