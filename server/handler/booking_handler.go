package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/constant"
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
		Source:           sql.NullString{String: booking.Source, Valid: true},
		Destination:      sql.NullString{String: booking.Destination, Valid: true},
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
	var tokenData interface{}
	tokenData, isExists := context.Get("claims")
	if !isExists {
		response.ErrorResponse(context, http.StatusUnauthorized, "Claims not found in request, request unauthorised")
		return
	}

	claims := tokenData.(model.TokenData)
	id := context.Request.URL.Query().Get("id")

	if id == "" {
		response.ErrorResponse(context, http.StatusNotFound, "ID not specified")
		return
	}

	var booking repository.GetBookingRow
	var err error
	if claims.Role == constant.CUSTOMER_ROLE {
		booking, err = handler.queries.GetBooking(context, repository.GetBookingParams{
			ID:         id,
			CustomerID: sql.NullString{Valid: true, String: claims.CustomerID},
		})

	} else {
		booking, err = handler.queries.AdminGetBooking(context, id)
	}

	if err != nil {
		fmt.Println(err.Error())
		response.ErrorResponse(context, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "quote Data",
		"data":    booking,
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
		Source:           booking.Source,
		Destination:      booking.Destination,
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

	var tokenData interface{}

	tokenData, isExists := context.Get("claims")
	if !isExists {
		response.ErrorResponse(context, http.StatusUnauthorized, "Claims not found in request, request unauthorised")
		return
	}

	claims := tokenData.(model.TokenData)

	booking := []repository.ListBookingsRow{}
	bookings := []repository.AdminListBookingsRow{}
	var err error
	if claims.Role == constant.CUSTOMER_ROLE {

		booking, err = handler.queries.ListBookings(context, claims.CustomerID)
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
			"data":    booking,
		})
		return

	}

	bookings, err = handler.queries.AdminListBookings(context)

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
		"data":    bookings,
	})
}
