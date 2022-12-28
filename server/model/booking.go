package model

type Booking struct {
	Id               string `json:"id"`
	BookingRequestId string `json:"booking_request_id"`
	Status           string `json:"status"`
	CustomerId       string `json:"customer_id"`
	TaskId           string `json:"task_id"`
	QuoteId          string `json:"quote_id"`
	MilestoneId      string `json:"milestone_id"`
	LinerId          string `json:"liner_id"`
	Source           string `json:"source"`
	Destination      string `json:"destination"`
	City             string `json:"city"`
}
