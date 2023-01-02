package repository

import (
	"database/sql"
)

type AccountDetail struct {
	ID          string
	Name        sql.NullString
	CompanyName sql.NullString
	Email       sql.NullString
	Mobile      sql.NullString
	Roles       sql.NullString
	City        sql.NullString
	Password    sql.NullString
}

type Booking struct {
	ID               string
	BookingRequestID string
	BookingStatus    sql.NullString
	CustomerID       sql.NullString
	TaskID           sql.NullString
	QuoteID          sql.NullString
	MilestoneID      sql.NullString
	LinerID          sql.NullString
	Source           sql.NullString
	Destination      sql.NullString
	City             sql.NullString
}

type BookingMilestone struct {
	ID              string
	BookingID       sql.NullString
	MilestoneStatus sql.NullString
	CreatedAt       sql.NullString
	CompletedAt     sql.NullString
}

type BookingTask struct {
	ID          string
	BookingID   sql.NullString
	TaskStatus  sql.NullString
	CreatedAt   sql.NullString
	CompletedAt sql.NullString
}

type Liner struct {
	ID   string
	Name sql.NullString
}

type Milestone struct {
	ID   string
	Name sql.NullString
}

type Partner struct {
	ID   string
	Name sql.NullString
}

type Quote struct {
	ID              string
	QuoteType       string
	CustomerID      string
	Source          string
	Destination     string
	DoorPickup      sql.NullString
	DoorAddress     sql.NullString
	DoorDelivery    sql.NullString
	DeliveryAddress sql.NullString
	LinerID         sql.NullString
	PartnerID       sql.NullString
	Validity        sql.NullString
	TransmitDays    sql.NullString
	FreeDays        sql.NullString
	Currency        sql.NullString
	Buy             sql.NullString
	Sell            sql.NullString
	PartnerTax      sql.NullString
	ProcurementID   sql.NullString
	SalesID         sql.NullString
}

type Task struct {
	ID   string
	Name sql.NullString
}
