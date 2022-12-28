package model

import (
	"database/sql"
)

type Quote struct {
	Id              string       `json:"id"`
	Type            string       `json:"type"`
	CustomerId      string       `json:"customer_id"`
	Source          string       `json:"source"`
	Destination     string       `json:"destination"`
	DoorPickup      bool         `json:"door_pickup"`
	DoorAddress     string       `json:"door_address"`
	DoorDelivery    bool         `json:"door_delivery"`
	DeliveryAddress string       `json:"delivery_address"`
	LinerId         string       `json:"liner_id"`
	PartnerId       string       `json:"partner_id"`
	Validity        sql.NullTime `json:"validity"`
	TransmitDays    int          `json:"transmit_days"`
	FreeDays        int          `json:"free_days"`
	Currency        string       `json:"currency"`
	Buy             int          `json:"buy"`
	Sell            int          `json:"sell"`
	PartnerTax      int          `json:"partner_tax"`
}
