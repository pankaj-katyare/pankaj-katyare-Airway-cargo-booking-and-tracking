package model

type BookingTask struct {
	Id          string `json:"id"`
	BookingId   string `json:"booking_id"`
	TaskID      string `json"task_id"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at"`
}
