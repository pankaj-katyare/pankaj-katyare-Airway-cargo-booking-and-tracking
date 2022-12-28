package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
)

type QuoteService struct {
	DB *sqlx.DB
}

func (service QuoteService) RequestQuote(req model.Quote) model.Quote {
	return model.Quote{
		Id:              req.Id,
		Type:            req.Type,
		CustomerId:      req.CustomerId,
		Source:          req.Source,
		Destination:     req.Destination,
		DoorPickup:      req.DoorPickup,
		DoorAddress:     req.DoorAddress,
		DoorDelivery:    req.DoorDelivery,
		DeliveryAddress: req.DeliveryAddress,
		LinerId:         req.LinerId,
		PartnerId:       req.PartnerId,
		Validity:        req.Validity,
		TransmitDays:    req.TransmitDays,
		FreeDays:        req.FreeDays,
		Currency:        req.Currency,
		Buy:             req.Buy,
		Sell:            req.Sell,
		PartnerTax:      req.PartnerTax,
	}
}
