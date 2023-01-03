package model

type BookingMilestone struct {
	Id          string `json:"id"`
	BookingId   string `json:"booking_id"`
	MilestoneID string `json"milestone_id"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at"`
}
