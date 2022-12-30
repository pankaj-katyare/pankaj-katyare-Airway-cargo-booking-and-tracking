package repository

import (
	"context"
	"database/sql"
)

const createAccountDetails = `-- name: CreateAccountDetails :one
INSERT INTO account_details (
id, name, email, mobile, roles, city) 
VALUES (
   $1,$2,$3,$4,$5, $6)
RETURNING id, name, email, mobile, roles, city
`

type CreateAccountDetailsParams struct {
	ID     string
	Name   sql.NullString
	Email  sql.NullString
	Mobile sql.NullString
	Roles  sql.NullString
	City   sql.NullString
}

func (q *Queries) CreateAccountDetails(ctx context.Context, arg CreateAccountDetailsParams) (AccountDetail, error) {
	row := q.db.QueryRowContext(ctx, createAccountDetails,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Mobile,
		arg.Roles,
		arg.City,
	)
	var i AccountDetail
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Mobile,
		&i.Roles,
		&i.City,
	)
	return i, err
}

const createBooking = `-- name: CreateBooking :one
INSERT INTO booking (
id,booking_request_id,booking_status,customer_id,task_id,quote_id,milestone_id,liner_id,source,destination,city) 
VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
RETURNING id, booking_request_id, booking_status, customer_id, task_id, quote_id, milestone_id, liner_id, source, destination, city
`

type CreateBookingParams struct {
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

func (q *Queries) CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error) {
	row := q.db.QueryRowContext(ctx, createBooking,
		arg.ID,
		arg.BookingRequestID,
		arg.BookingStatus,
		arg.CustomerID,
		arg.TaskID,
		arg.QuoteID,
		arg.MilestoneID,
		arg.LinerID,
		arg.Source,
		arg.Destination,
		arg.City,
	)
	var i Booking
	err := row.Scan(
		&i.ID,
		&i.BookingRequestID,
		&i.BookingStatus,
		&i.CustomerID,
		&i.TaskID,
		&i.QuoteID,
		&i.MilestoneID,
		&i.LinerID,
		&i.Source,
		&i.Destination,
		&i.City,
	)
	return i, err
}

const createBookingMilestone = `-- name: CreateBookingMilestone :one
INSERT INTO booking_milestone (
id, booking_id, milestone_status, created_at, completed_at) 
VALUES (
   $1,$2,$3,$4,$5)
RETURNING id, booking_id, milestone_status, created_at, completed_at
`

type CreateBookingMilestoneParams struct {
	ID              string
	BookingID       sql.NullString
	MilestoneStatus sql.NullString
	CreatedAt       sql.NullString
	CompletedAt     sql.NullString
}

func (q *Queries) CreateBookingMilestone(ctx context.Context, arg CreateBookingMilestoneParams) (BookingMilestone, error) {
	row := q.db.QueryRowContext(ctx, createBookingMilestone,
		arg.ID,
		arg.BookingID,
		arg.MilestoneStatus,
		arg.CreatedAt,
		arg.CompletedAt,
	)
	var i BookingMilestone
	err := row.Scan(
		&i.ID,
		&i.BookingID,
		&i.MilestoneStatus,
		&i.CreatedAt,
		&i.CompletedAt,
	)
	return i, err
}

const createBookingTask = `-- name: CreateBookingTask :one
INSERT INTO booking_task (
id, booking_id, task_status, created_at, completed_at) 
VALUES (
   $1,$2,$3,$4,$5)
RETURNING id, booking_id, task_status, created_at, completed_at
`

type CreateBookingTaskParams struct {
	ID          string
	BookingID   sql.NullString
	TaskStatus  sql.NullString
	CreatedAt   sql.NullString
	CompletedAt sql.NullString
}

func (q *Queries) CreateBookingTask(ctx context.Context, arg CreateBookingTaskParams) (BookingTask, error) {
	row := q.db.QueryRowContext(ctx, createBookingTask,
		arg.ID,
		arg.BookingID,
		arg.TaskStatus,
		arg.CreatedAt,
		arg.CompletedAt,
	)
	var i BookingTask
	err := row.Scan(
		&i.ID,
		&i.BookingID,
		&i.TaskStatus,
		&i.CreatedAt,
		&i.CompletedAt,
	)
	return i, err
}

const createLiner = `-- name: CreateLiner :one
INSERT INTO liners (
id, name) 
VALUES (
    $1,$2)
RETURNING id, name
`

type CreateLinerParams struct {
	ID   string
	Name sql.NullString
}

func (q *Queries) CreateLiner(ctx context.Context, arg CreateLinerParams) (Liner, error) {
	row := q.db.QueryRowContext(ctx, createLiner, arg.ID, arg.Name)
	var i Liner
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createMilestone = `-- name: CreateMilestone :one
INSERT INTO milestones (
id, name) 
VALUES (
    $1,$2)
RETURNING id, name
`

type CreateMilestoneParams struct {
	ID   string
	Name sql.NullString
}

func (q *Queries) CreateMilestone(ctx context.Context, arg CreateMilestoneParams) (Milestone, error) {
	row := q.db.QueryRowContext(ctx, createMilestone, arg.ID, arg.Name)
	var i Milestone
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createPartner = `-- name: CreatePartner :one
INSERT INTO partners (
id, name) 
VALUES (
    $1,$2)
RETURNING id, name
`

type CreatePartnerParams struct {
	ID   string
	Name sql.NullString
}

func (q *Queries) CreatePartner(ctx context.Context, arg CreatePartnerParams) (Partner, error) {
	row := q.db.QueryRowContext(ctx, createPartner, arg.ID, arg.Name)
	var i Partner
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
id, name) 
VALUES (
    $1,$2)
RETURNING id, name
`

type CreateTaskParams struct {
	ID   string
	Name sql.NullString
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.ID, arg.Name)
	var i Task
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteAccountDetails = `-- name: DeleteAccountDetails :exec
DELETE FROM account_details
WHERE id = $1
`

func (q *Queries) DeleteAccountDetails(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteAccountDetails, id)
	return err
}

const deleteBooking = `-- name: DeleteBooking :exec
DELETE FROM booking
WHERE id = $1
`

func (q *Queries) DeleteBooking(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteBooking, id)
	return err
}

const deleteBookingMilestone = `-- name: DeleteBookingMilestone :exec
DELETE FROM booking_milestone
WHERE id = $1
`

func (q *Queries) DeleteBookingMilestone(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteBookingMilestone, id)
	return err
}

const deleteBookingTask = `-- name: DeleteBookingTask :exec
DELETE FROM booking_task
WHERE id = $1
`

func (q *Queries) DeleteBookingTask(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteBookingTask, id)
	return err
}

const deleteLiner = `-- name: DeleteLiner :exec
DELETE FROM liners
WHERE id = $1
`

func (q *Queries) DeleteLiner(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteLiner, id)
	return err
}

const deleteMilestone = `-- name: DeleteMilestone :exec
DELETE FROM milestones
WHERE id = $1
`

func (q *Queries) DeleteMilestone(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteMilestone, id)
	return err
}

const deletePartner = `-- name: DeletePartner :exec
DELETE FROM partners
WHERE id = $1
`

func (q *Queries) DeletePartner(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePartner, id)
	return err
}

const deleteQuote = `-- name: DeleteQuote :exec
DELETE FROM quotes
WHERE id = $1
`

func (q *Queries) DeleteQuote(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteQuote, id)
	return err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getAccountDetails = `-- name: GetAccountDetails :one
SELECT id, name, email, mobile, roles, city FROM account_details
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccountDetails(ctx context.Context, id string) (AccountDetail, error) {
	row := q.db.QueryRowContext(ctx, getAccountDetails, id)
	var i AccountDetail
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Mobile,
		&i.Roles,
		&i.City,
	)
	return i, err
}

const getBooking = `-- name: GetBooking :one
SELECT id, booking_request_id, booking_status, customer_id, task_id, quote_id, milestone_id, liner_id, source, destination, city FROM booking
WHERE id = $1 AND booking_request_id=$2 LIMIT 1
`

type GetBookingParams struct {
	ID               string
	BookingRequestID string
}

func (q *Queries) GetBooking(ctx context.Context, arg GetBookingParams) (Booking, error) {
	row := q.db.QueryRowContext(ctx, getBooking, arg.ID, arg.BookingRequestID)
	var i Booking
	err := row.Scan(
		&i.ID,
		&i.BookingRequestID,
		&i.BookingStatus,
		&i.CustomerID,
		&i.TaskID,
		&i.QuoteID,
		&i.MilestoneID,
		&i.LinerID,
		&i.Source,
		&i.Destination,
		&i.City,
	)
	return i, err
}

const getBookingMilestone = `-- name: GetBookingMilestone :one
SELECT id, booking_id, milestone_status, created_at, completed_at FROM booking_milestone
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBookingMilestone(ctx context.Context, id string) (BookingMilestone, error) {
	row := q.db.QueryRowContext(ctx, getBookingMilestone, id)
	var i BookingMilestone
	err := row.Scan(
		&i.ID,
		&i.BookingID,
		&i.MilestoneStatus,
		&i.CreatedAt,
		&i.CompletedAt,
	)
	return i, err
}

const getBookingTask = `-- name: GetBookingTask :one
SELECT id, booking_id, task_status, created_at, completed_at FROM booking_task
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBookingTask(ctx context.Context, id string) (BookingTask, error) {
	row := q.db.QueryRowContext(ctx, getBookingTask, id)
	var i BookingTask
	err := row.Scan(
		&i.ID,
		&i.BookingID,
		&i.TaskStatus,
		&i.CreatedAt,
		&i.CompletedAt,
	)
	return i, err
}

const getLiner = `-- name: GetLiner :one
SELECT id, name FROM liners
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLiner(ctx context.Context, id string) (Liner, error) {
	row := q.db.QueryRowContext(ctx, getLiner, id)
	var i Liner
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getMilestone = `-- name: GetMilestone :one
SELECT id, name FROM milestones
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMilestone(ctx context.Context, id string) (Milestone, error) {
	row := q.db.QueryRowContext(ctx, getMilestone, id)
	var i Milestone
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getPartner = `-- name: GetPartner :one
SELECT id, name FROM partners
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPartner(ctx context.Context, id string) (Partner, error) {
	row := q.db.QueryRowContext(ctx, getPartner, id)
	var i Partner
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getQuote = `-- name: GetQuote :one
SELECT id, quote_type, customer_id, source, destination, door_pickup, door_address, door_delivery, delivery_address, liner_id, partner_id, validity, transmit_days, free_days, currency, buy, sell, partner_tax FROM quotes
WHERE id = $1 AND customer_id=$2 LIMIT 1
`

type GetQuoteParams struct {
	ID         string
	CustomerID string
}

func (q *Queries) GetQuote(ctx context.Context, arg GetQuoteParams) (Quote, error) {
	row := q.db.QueryRowContext(ctx, getQuote, arg.ID, arg.CustomerID)
	var i Quote
	err := row.Scan(
		&i.ID,
		&i.QuoteType,
		&i.CustomerID,
		&i.Source,
		&i.Destination,
		&i.DoorPickup,
		&i.DoorAddress,
		&i.DoorDelivery,
		&i.DeliveryAddress,
		&i.LinerID,
		&i.PartnerID,
		&i.Validity,
		&i.TransmitDays,
		&i.FreeDays,
		&i.Currency,
		&i.Buy,
		&i.Sell,
		&i.PartnerTax,
	)
	return i, err
}

const getTask = `-- name: GetTask :one
SELECT id, name FROM tasks
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTask(ctx context.Context, id string) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listAccountDetails = `-- name: ListAccountDetails :many
SELECT id, name, email, mobile, roles, city FROM account_details
`

func (q *Queries) ListAccountDetails(ctx context.Context) ([]AccountDetail, error) {
	rows, err := q.db.QueryContext(ctx, listAccountDetails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AccountDetail
	for rows.Next() {
		var i AccountDetail
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Mobile,
			&i.Roles,
			&i.City,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBookingMilestone = `-- name: ListBookingMilestone :many
SELECT id, booking_id, milestone_status, created_at, completed_at FROM booking_milestone
`

func (q *Queries) ListBookingMilestone(ctx context.Context) ([]BookingMilestone, error) {
	rows, err := q.db.QueryContext(ctx, listBookingMilestone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BookingMilestone
	for rows.Next() {
		var i BookingMilestone
		if err := rows.Scan(
			&i.ID,
			&i.BookingID,
			&i.MilestoneStatus,
			&i.CreatedAt,
			&i.CompletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBookingTask = `-- name: ListBookingTask :many
SELECT id, booking_id, task_status, created_at, completed_at FROM booking_task
`

func (q *Queries) ListBookingTask(ctx context.Context) ([]BookingTask, error) {
	rows, err := q.db.QueryContext(ctx, listBookingTask)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BookingTask
	for rows.Next() {
		var i BookingTask
		if err := rows.Scan(
			&i.ID,
			&i.BookingID,
			&i.TaskStatus,
			&i.CreatedAt,
			&i.CompletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBookings = `-- name: ListBookings :many
SELECT id, booking_request_id, booking_status, customer_id, task_id, quote_id, milestone_id, liner_id, source, destination, city FROM booking
WHERE customer_id=$1
`

func (q *Queries) ListBookings(ctx context.Context, customerID sql.NullString) ([]Booking, error) {
	rows, err := q.db.QueryContext(ctx, listBookings, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Booking
	for rows.Next() {
		var i Booking
		if err := rows.Scan(
			&i.ID,
			&i.BookingRequestID,
			&i.BookingStatus,
			&i.CustomerID,
			&i.TaskID,
			&i.QuoteID,
			&i.MilestoneID,
			&i.LinerID,
			&i.Source,
			&i.Destination,
			&i.City,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLiners = `-- name: ListLiners :many
SELECT id, name FROM liners
`

func (q *Queries) ListLiners(ctx context.Context) ([]Liner, error) {
	rows, err := q.db.QueryContext(ctx, listLiners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Liner
	for rows.Next() {
		var i Liner
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMilestones = `-- name: ListMilestones :many
SELECT id, name FROM milestones
`

func (q *Queries) ListMilestones(ctx context.Context) ([]Milestone, error) {
	rows, err := q.db.QueryContext(ctx, listMilestones)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Milestone
	for rows.Next() {
		var i Milestone
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPartners = `-- name: ListPartners :many
SELECT id, name FROM partners
`

func (q *Queries) ListPartners(ctx context.Context) ([]Partner, error) {
	rows, err := q.db.QueryContext(ctx, listPartners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Partner
	for rows.Next() {
		var i Partner
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listQuotes = `-- name: ListQuotes :many
SELECT id, quote_type, customer_id, source, destination, door_pickup, door_address, door_delivery, delivery_address, liner_id, partner_id, validity, transmit_days, free_days, currency, buy, sell, partner_tax FROM quotes
WHERE customer_id=$1
`

func (q *Queries) ListQuotes(ctx context.Context, customerID string) ([]Quote, error) {
	rows, err := q.db.QueryContext(ctx, listQuotes, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quote
	for rows.Next() {
		var i Quote
		if err := rows.Scan(
			&i.ID,
			&i.QuoteType,
			&i.CustomerID,
			&i.Source,
			&i.Destination,
			&i.DoorPickup,
			&i.DoorAddress,
			&i.DoorDelivery,
			&i.DeliveryAddress,
			&i.LinerID,
			&i.PartnerID,
			&i.Validity,
			&i.TransmitDays,
			&i.FreeDays,
			&i.Currency,
			&i.Buy,
			&i.Sell,
			&i.PartnerTax,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTasks = `-- name: ListTasks :many
SELECT id, name FROM tasks
`

func (q *Queries) ListTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const requestQuote = `-- name: RequestQuote :one
INSERT INTO quotes (
    id,quote_type,customer_id,source,destination,door_pickup,door_address,
    door_delivery,delivery_address,liner_id,partner_id,validity,transmit_days,
    free_days,currency,buy,sell,partner_tax) 
VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)
RETURNING id, quote_type, customer_id, source, destination, door_pickup, door_address, door_delivery, delivery_address, liner_id, partner_id, validity, transmit_days, free_days, currency, buy, sell, partner_tax
`

type RequestQuoteParams struct {
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
}

func (q *Queries) RequestQuote(ctx context.Context, arg RequestQuoteParams) (Quote, error) {
	row := q.db.QueryRowContext(ctx, requestQuote,
		arg.ID,
		arg.QuoteType,
		arg.CustomerID,
		arg.Source,
		arg.Destination,
		arg.DoorPickup,
		arg.DoorAddress,
		arg.DoorDelivery,
		arg.DeliveryAddress,
		arg.LinerID,
		arg.PartnerID,
		arg.Validity,
		arg.TransmitDays,
		arg.FreeDays,
		arg.Currency,
		arg.Buy,
		arg.Sell,
		arg.PartnerTax,
	)
	var i Quote
	err := row.Scan(
		&i.ID,
		&i.QuoteType,
		&i.CustomerID,
		&i.Source,
		&i.Destination,
		&i.DoorPickup,
		&i.DoorAddress,
		&i.DoorDelivery,
		&i.DeliveryAddress,
		&i.LinerID,
		&i.PartnerID,
		&i.Validity,
		&i.TransmitDays,
		&i.FreeDays,
		&i.Currency,
		&i.Buy,
		&i.Sell,
		&i.PartnerTax,
	)
	return i, err
}

const updateAccountDetails = `-- name: UpdateAccountDetails :exec
UPDATE account_details set 
    id = $1,
    name = $2,
    email = $3,
    mobile = $4,
    roles = $5,
    city = $6
WHERE id = $1
`

type UpdateAccountDetailsParams struct {
	ID     string
	Name   sql.NullString
	Email  sql.NullString
	Mobile sql.NullString
	Roles  sql.NullString
	City   sql.NullString
}

func (q *Queries) UpdateAccountDetails(ctx context.Context, arg UpdateAccountDetailsParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountDetails,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Mobile,
		arg.Roles,
		arg.City,
	)
	return err
}

const updateBooking = `-- name: UpdateBooking :exec
UPDATE booking set 
    Id = $1,
    booking_request_id = $2,
    booking_status= $3,
    customer_id = $4,
    task_id = $5,
    quote_id = $6,
    milestone_id = $7,
    liner_id = $8,
    source = $9,
    destination = $10,
    city = $11
WHERE id = $1
`

type UpdateBookingParams struct {
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

func (q *Queries) UpdateBooking(ctx context.Context, arg UpdateBookingParams) error {
	_, err := q.db.ExecContext(ctx, updateBooking,
		arg.ID,
		arg.BookingRequestID,
		arg.BookingStatus,
		arg.CustomerID,
		arg.TaskID,
		arg.QuoteID,
		arg.MilestoneID,
		arg.LinerID,
		arg.Source,
		arg.Destination,
		arg.City,
	)
	return err
}

const updateBookingMilestone = `-- name: UpdateBookingMilestone :exec
UPDATE booking_milestone set 
    id = $1,
    booking_id = $2,
    milestone_status = $3,
    created_at = $4,
    completed_at = $5
WHERE id = $1
`

type UpdateBookingMilestoneParams struct {
	ID              string
	BookingID       sql.NullString
	MilestoneStatus sql.NullString
	CreatedAt       sql.NullString
	CompletedAt     sql.NullString
}

func (q *Queries) UpdateBookingMilestone(ctx context.Context, arg UpdateBookingMilestoneParams) error {
	_, err := q.db.ExecContext(ctx, updateBookingMilestone,
		arg.ID,
		arg.BookingID,
		arg.MilestoneStatus,
		arg.CreatedAt,
		arg.CompletedAt,
	)
	return err
}

const updateBookingTask = `-- name: UpdateBookingTask :exec
UPDATE booking_task set 
    id = $1,
    booking_id = $2,
    task_status = $3,
    created_at = $4,
    completed_at = $5
WHERE id = $1
`

type UpdateBookingTaskParams struct {
	ID          string
	BookingID   sql.NullString
	TaskStatus  sql.NullString
	CreatedAt   sql.NullString
	CompletedAt sql.NullString
}

func (q *Queries) UpdateBookingTask(ctx context.Context, arg UpdateBookingTaskParams) error {
	_, err := q.db.ExecContext(ctx, updateBookingTask,
		arg.ID,
		arg.BookingID,
		arg.TaskStatus,
		arg.CreatedAt,
		arg.CompletedAt,
	)
	return err
}

const updateLiner = `-- name: UpdateLiner :exec
UPDATE liners set 
    id = $1,
    name = $2
WHERE id = $1
`

type UpdateLinerParams struct {
	ID   string
	Name sql.NullString
}

func (q *Queries) UpdateLiner(ctx context.Context, arg UpdateLinerParams) error {
	_, err := q.db.ExecContext(ctx, updateLiner, arg.ID, arg.Name)
	return err
}

const updateMilestone = `-- name: UpdateMilestone :exec
UPDATE milestones set 
    id = $1,
    name = $2
WHERE id = $1
`

type UpdateMilestoneParams struct {
	ID   string
	Name sql.NullString
}

func (q *Queries) UpdateMilestone(ctx context.Context, arg UpdateMilestoneParams) error {
	_, err := q.db.ExecContext(ctx, updateMilestone, arg.ID, arg.Name)
	return err
}

const updatePartner = `-- name: UpdatePartner :exec
UPDATE partners set 
    id = $1,
    name = $2
WHERE id = $1
`

type UpdatePartnerParams struct {
	ID   string
	Name sql.NullString
}

func (q *Queries) UpdatePartner(ctx context.Context, arg UpdatePartnerParams) error {
	_, err := q.db.ExecContext(ctx, updatePartner, arg.ID, arg.Name)
	return err
}

const updateQuote = `-- name: UpdateQuote :exec
UPDATE quotes set 
    buy = $2,
    sell = $3,
    liner_id = $4,
    partner_id = $5,
    validity = $6,
    transmit_days = $7,
    free_days = $8,
    currency = $9,
    partner_tax = $10
WHERE id = $1
`

type UpdateQuoteParams struct {
	ID           string
	Buy          sql.NullString
	Sell         sql.NullString
	LinerID      sql.NullString
	PartnerID    sql.NullString
	Validity     sql.NullString
	TransmitDays sql.NullString
	FreeDays     sql.NullString
	Currency     sql.NullString
	PartnerTax   sql.NullString
}

func (q *Queries) UpdateQuote(ctx context.Context, arg UpdateQuoteParams) error {
	_, err := q.db.ExecContext(ctx, updateQuote,
		arg.ID,
		arg.Buy,
		arg.Sell,
		arg.LinerID,
		arg.PartnerID,
		arg.Validity,
		arg.TransmitDays,
		arg.FreeDays,
		arg.Currency,
		arg.PartnerTax,
	)
	return err
}

const updateTask = `-- name: UpdateTask :exec
UPDATE tasks set 
    id = $1,
    name = $2
WHERE id = $1
`

type UpdateTaskParams struct {
	ID   string
	Name sql.NullString
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.db.ExecContext(ctx, updateTask, arg.ID, arg.Name)
	return err
}