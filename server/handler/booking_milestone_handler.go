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

type BookingMilestoneHandler struct {
	DB      *sqlx.DB
	queries *repository.Queries
}

func NewBookingMilestoneHandler(DB *sqlx.DB) *BookingMilestoneHandler {
	return &BookingMilestoneHandler{
		DB:      DB,
		queries: repository.New(DB),
	}
}

func (handler BookingMilestoneHandler) CreateBookingMilestone(context *gin.Context) {

	var bookingMilestone repository.BookingMilestone

	if err := context.ShouldBind(&bookingMilestone); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		fmt.Println("error", err)
		return
	}

	state, err := handler.queries.CreateBookingMilestone(context, repository.CreateBookingMilestoneParams{
		ID:              uuid.New().String(),
		BookingID:       sql.NullString{String: bookingMilestone.BookingID.String, Valid: true},
		MilestoneStatus: sql.NullString{String: bookingMilestone.MilestoneStatus.String, Valid: true},
		CreatedAt:       sql.NullString{String: bookingMilestone.CreatedAt.String, Valid: true},
		CompletedAt:     sql.NullString{String: bookingMilestone.CompletedAt.String, Valid: true},
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

func (handler BookingMilestoneHandler) GetBookingMilestone(context *gin.Context) {

	id := context.Request.URL.Query().Get("id")

	if id == "" {
		response.ErrorResponse(context, http.StatusNotFound, "ID not specified")
		return
	}

	state, err := handler.queries.GetBookingMilestone(context, id)

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

func (handler BookingMilestoneHandler) UpdateBookingMilestone(context *gin.Context) {

	var bookingMilestone repository.BookingMilestone

	if err := context.ShouldBind(&bookingMilestone); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}

	state := handler.queries.UpdateBookingMilestone(context, repository.UpdateBookingMilestoneParams{
		ID:              uuid.New().String(),
		BookingID:       bookingMilestone.BookingID,
		MilestoneStatus: bookingMilestone.MilestoneStatus,
		CreatedAt:       bookingMilestone.CreatedAt,
		CompletedAt:     bookingMilestone.CompletedAt,
	})

	// TODO return, nothing to update
	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Updated suceessfully",
		"data":    state,
	})
}

func (handler BookingMilestoneHandler) ListBookingMilestone(context *gin.Context) {

	bookingMilestones, err := handler.queries.ListBookingMilestone(context)

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
