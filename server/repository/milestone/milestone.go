package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/db"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
)

type (
	Milestone interface {
		Insert(ctx context.Context, milestones model.Milestone) error
		Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) (model.Milestone, error)
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

func (s *SQL) Insert(ctx context.Context, Milestone model.Milestone) error {
	query := `insert into milestones (id, name) 
	values (?,?);`

	_, err := s.db.ExecContext(ctx, query, Milestone.Id, Milestone.Name)
	return err
}

func (s *SQL) Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) ([]*model.Milestone, error) {
	var partners []*model.Milestone
	mainQuery := db.SelectBuilder("milestones", where, orderBy, groupBy)
	err := s.db.GetContext(ctx, &partners, mainQuery)
	if err != nil {
		return nil, err
	}
	return partners, nil
}

func (s *SQL) Update(ctx context.Context, set map[string]interface{}, condition []string) (int64, error) {
	if len(set) <= 0 {
		return 0, errors.New("no data to update in update call")
	}
	query := db.UpdateBuilder("milestones", set, condition)
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
	query := db.DeleteBuilder("milestones", condition)
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
