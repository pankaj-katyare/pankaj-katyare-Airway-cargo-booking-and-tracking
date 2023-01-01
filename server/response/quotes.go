package response

type GetQuoteByIDResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type RequestQuoteResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type UpdateQuoteResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type GenericResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	ID              string `json:"id"`
	QuoteType       string `json:"quote_type"`
	CustomerID      string `json:"customer_id"`
	Source          string `json:"source"`
	Destination     string `json:"destination"`
	DoorPickup      string `json:"doorPickup"`
	DoorAddress     string `json:"doorAddress"`
	DoorDelivery    string `json:"door_delivery"`
	DeliveryAddress string `json:"delivery_address"`
	LinerID         string `json:"liner_id"`
	PartnerID       string `json:"partner_id"`
	Validity        string `json:"validity"`
	TransmitDays    string `json:"transmit_days"`
	FreeDays        string `json:"free_days"`
	Currency        string `json:"currency"`
	Buy             string `json:"buy"`
	Sell            string `json:"sell"`
	PartnerTax      string `json:"partner_tax"`
}
