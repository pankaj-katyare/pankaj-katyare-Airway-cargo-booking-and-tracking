package model

type Quote struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	CustomerId      string `json:"customer_id"`
	Source          string `json:"source"`
	Destination     string `json:"destination"`
	DoorPickup      string `json:"door_pickup"`
	DoorAddress     string `json:"door_address"`
	DoorDelivery    string `json:"door_delivery"`
	DeliveryAddress string `json:"delivery_address"`
	LinerId         string `json:"liner_id"`
	PartnerId       string `json:"partner_id"`
	Validity        string `json:"validity"`
	TransmitDays    string `json:"transmit_days"`
	FreeDays        string `json:"free_days"`
	Currency        string `json:"currency"`
	Buy             string `json:"buy"`
	Sell            string `json:"sell"`
	PartnerTax      string `json:"partner_tax"`
}
