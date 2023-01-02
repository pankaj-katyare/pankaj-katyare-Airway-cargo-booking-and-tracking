package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/repository"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/response"

	"github.com/gin-gonic/gin"
)

type BookingTaskHandler struct {
	DB      *sqlx.DB
	queries *repository.Queries
}

func NewBookingTaskHandler(DB *sqlx.DB) *BookingTaskHandler {
	return &BookingTaskHandler{
		DB:      DB,
		queries: repository.New(DB),
	}
}

func (handler BookingTaskHandler) CreateBookingTask(context *gin.Context) {

	var bookingTask repository.BookingTask

	if err := context.ShouldBind(&bookingTask); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		fmt.Println("error", err)
		return
	}

	state, err := handler.queries.CreateBookingTask(context, repository.CreateBookingTaskParams{
		ID:          uuid.New().String(),
		BookingID:   sql.NullString{String: bookingTask.BookingID.String, Valid: true},
		TaskStatus:  sql.NullString{String: bookingTask.TaskStatus.String, Valid: true},
		CreatedAt:   sql.NullString{String: bookingTask.CreatedAt.String, Valid: true},
		CompletedAt: sql.NullString{String: bookingTask.CompletedAt.String, Valid: true},
	})
	if err != nil {
		fmt.Println("error", err)
		response.ErrorResponse(context, http.StatusBadRequest, "Error inserting Booking milestone")
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Request Booking milestone Created successfuly",
		"data":    state,
	})
}

func (handler BookingTaskHandler) GetBookingTask(context *gin.Context) {

	id := context.Request.URL.Query().Get("id")

	if id == "" {
		response.ErrorResponse(context, http.StatusNotFound, "ID not specified")
		return
	}

	state, err := handler.queries.GetBookingTask(context, id)

	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "account Data",
		"data":    state,
	})
}

func (handler BookingTaskHandler) UpdateBookingTask(context *gin.Context) {

	var bookingTask repository.BookingTask

	if err := context.ShouldBind(&bookingTask); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}

	state := handler.queries.UpdateBookingTask(context, repository.UpdateBookingTaskParams{
		ID:          uuid.New().String(),
		BookingID:   bookingTask.BookingID,
		TaskStatus:  bookingTask.TaskStatus,
		CreatedAt:   bookingTask.CreatedAt,
		CompletedAt: bookingTask.CompletedAt,
	})

	// TODO return, nothing to update
	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Updated suceessfully",
		"data":    state,
	})
}

func (handler BookingTaskHandler) ListBookingTask(context *gin.Context) {

	bookingMilestones, err := handler.queries.ListBookingTask(context)

	if err != nil {
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "success",
			"message": "Error int get all account",
			"error":   err,
		})
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Fetched all account list",
		"data":    bookingMilestones,
	})
}
