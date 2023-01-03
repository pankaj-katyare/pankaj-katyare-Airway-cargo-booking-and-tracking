package model

type Booking struct {
	Id               string `json:"id"`
	BookingRequestId string `json:"booking_request_id"`
	Status           string `json:"booking_status"`
	CustomerId       string `json:"customer_id"`
	Source           string `json:"source"`
	Destination      string `json:"destination"`
	City             string `json:"city"`
}
