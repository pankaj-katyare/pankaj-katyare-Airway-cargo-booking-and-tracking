package quotes

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/db"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
)

type (
	Quote interface {
		Insert(ctx context.Context, quote []*model.Quote) error
		Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) (model.Quote, error)
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

func (s *SQL) Insert(ctx context.Context, quote model.Quote) error {
	query := `insert into quotes (id, type, customer_id, source, destination, 
		door_pickup, door_address, door_delivery, delivery_address, liner_id, partner_id, validity, transmit_days, free_days, currency, buy, sell, partner_tax) values
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	_, err := s.db.ExecContext(ctx, query, quote.Id, quote.Type, quote.CustomerId, quote.Source, quote.Destination, quote.DoorPickup, quote.DoorAddress, quote.DoorDelivery, quote.DeliveryAddress, quote.LinerId, quote.PartnerId, quote.Validity, quote.TransmitDays, quote.FreeDays, quote.Currency, quote.Buy, quote.Sell, quote.PartnerTax)
	return err
}

func (s *SQL) Select(ctx context.Context, where map[string]interface{}, orderBy map[string]bool, groupBy []string) ([]*model.Quote, error) {
	var quotes []*model.Quote
	mainQuery := db.SelectBuilder("quotes", where, orderBy, groupBy)
	err := s.db.GetContext(ctx, &quotes, mainQuery)
	if err != nil {
		return nil, err
	}
	return quotes, nil
}

func (s *SQL) Update(ctx context.Context, set map[string]interface{}, condition []string) (int64, error) {
	if len(set) <= 0 {
		return 0, errors.New("no data to update in update call")
	}
	query := db.UpdateBuilder("quotes", set, condition)
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
	query := db.DeleteBuilder("quotes", condition)
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
