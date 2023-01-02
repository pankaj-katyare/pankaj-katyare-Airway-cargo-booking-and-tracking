package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/repository"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/response"
)

type BookingHandler struct {
	DB      *sqlx.DB
	queries *repository.Queries
}

func NewBookingHandler(DB *sqlx.DB) *BookingHandler {
	return &BookingHandler{
		DB:      DB,
		queries: repository.New(DB),
	}
}

func (handler BookingHandler) CreateBooking(context *gin.Context) {

	var booking model.Booking

	if err := context.ShouldBind(&booking); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}

	state, err := handler.queries.CreateBooking(context, repository.CreateBookingParams{
		ID:               uuid.New().String(),
		BookingRequestID: uuid.New().String(),
		BookingStatus:    sql.NullString{String: booking.Status, Valid: true},
		CustomerID:       sql.NullString{String: booking.CustomerId, Valid: true},
		TaskID:           sql.NullString{String: booking.TaskId, Valid: true},
		QuoteID:          sql.NullString{String: booking.QuoteId, Valid: true},
		MilestoneID:      sql.NullString{String: booking.MilestoneId, Valid: true},
		LinerID:          sql.NullString{String: booking.LinerId, Valid: true},
		Source:           sql.NullString{String: booking.Source, Valid: true},
		Destination:      sql.NullString{String: booking.Destination, Valid: true},
		City:             sql.NullString{String: booking.City, Valid: true},
	})

	if err != nil {
		fmt.Println("error", err)
		response.ErrorResponse(context, http.StatusBadRequest, "Error inserting Booking")
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Request Booking Created successfuly",
		"data":    state,
	})

}

func (handler BookingHandler) GetBookingByID(context *gin.Context) {

	id := context.Request.URL.Query().Get("id")

	if id == "" {
		response.ErrorResponse(context, http.StatusNotFound, " id is invalid ")
		return
	}

	bookings, err := handler.queries.GetBooking(context, repository.GetBookingParams{
		ID: id,
	})

	if err == nil {
		response.ErrorResponse(context, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Quote Data",
		"data":    bookings,
	})
}

func (handler BookingHandler) UpdateBooking(context *gin.Context) {

	var booking repository.UpdateBookingParams

	if err := context.ShouldBind(&booking); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}

	err := handler.queries.UpdateBooking(context, repository.UpdateBookingParams{
		ID:               booking.ID,
		BookingRequestID: booking.BookingRequestID,
		BookingStatus:    booking.BookingStatus,
		CustomerID:       booking.CustomerID,
		TaskID:           booking.TaskID,
		QuoteID:          booking.QuoteID,
		MilestoneID:      booking.MilestoneID,
		LinerID:          booking.LinerID,
		Source:           booking.Source,
		Destination:      booking.Destination,
		City:             booking.City,
	})

	if err != nil {
		// TODO: return, error updating quote in database
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "success",
			"message": "Error updating quote in database",
		})
	}
	// TODO return, quote updated successfuly
	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Booking updated successfuly",
	})

}

func (handler BookingHandler) ListBookings(context *gin.Context) {

	id := context.Request.URL.Query().Get("id")

	quotes, err := handler.queries.ListQuotes(context, id)

	if err != nil {
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "success",
			"message": "Error int get all booking",
			"error":   err,
		})
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Fetched all booking list",
		"data":    quotes,
	})
}
