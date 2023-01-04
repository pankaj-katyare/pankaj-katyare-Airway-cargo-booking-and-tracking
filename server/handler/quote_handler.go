package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/constant"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/repository"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/response"

	"github.com/gin-gonic/gin"
)

type QuoteHandler struct {
	DB      *sqlx.DB
	queries *repository.Queries
}

func NewQuoteHandler(DB *sqlx.DB) *QuoteHandler {
	return &QuoteHandler{
		DB:      DB,
		queries: repository.New(DB),
	}
}

func (handler QuoteHandler) RequestQuote(context *gin.Context) {

	var quote model.Quote

	if err := context.ShouldBind(&quote); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		fmt.Println("error", err)
		return
	}

	quote.Id = uuid.New().String()

	state, err := handler.queries.RequestQuote(context, repository.RequestQuoteParams{
		ID:              uuid.New().String(),
		QuoteType:       quote.Type,
		CustomerID:      quote.CustomerId,
		Source:          quote.Source,
		Destination:     quote.Destination,
		DoorPickup:      sql.NullString{String: quote.DoorPickup, Valid: true},
		DoorAddress:     sql.NullString{String: quote.DoorAddress, Valid: true},
		DoorDelivery:    sql.NullString{String: quote.DoorDelivery, Valid: true},
		DeliveryAddress: sql.NullString{String: quote.DeliveryAddress, Valid: true},
		LinerID:         sql.NullString{String: quote.LinerId, Valid: false},
		PartnerID:       sql.NullString{String: quote.PartnerId, Valid: false},
		Validity:        sql.NullString{String: quote.Validity, Valid: false},
		TransmitDays:    sql.NullString{String: quote.TransmitDays, Valid: false},
		FreeDays:        sql.NullString{String: quote.FreeDays, Valid: false},
		Currency:        sql.NullString{String: quote.Currency, Valid: false},
		Buy:             sql.NullString{String: quote.Buy, Valid: false},
		Sell:            sql.NullString{String: quote.Sell, Valid: false},
		PartnerTax:      sql.NullString{String: quote.PartnerTax, Valid: false},
	})
	if err != nil {
		fmt.Println("error", err)
		response.ErrorResponse(context, http.StatusBadRequest, "Error inserting quote")
		return
	}
	item := ListAllQuotesResult{
		ID:              state.ID,
		QuoteType:       state.QuoteType,
		CustomerID:      state.CustomerID,
		Source:          state.Source,
		Destination:     state.Destination,
		DoorPickup:      state.DoorPickup.String,
		DoorAddress:     state.DoorAddress.String,
		DoorDelivery:    state.DoorDelivery.String,
		DeliveryAddress: state.DeliveryAddress.String,
		LinerID:         state.LinerID.String,
		PartnerID:       state.PartnerID.String,
		Validity:        state.Validity.String,
		TransmitDays:    state.TransmitDays.String,
		FreeDays:        state.FreeDays.String,
		Currency:        state.Currency.String,
		Buy:             state.Buy.String,
		Sell:            state.Sell.String,
		PartnerTax:      state.PartnerTax.String,
		QuoteStatus:     state.QuoteStatus.String,
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Request Quote Created successfuly",
		"data":    item,
	})
}

func (handler QuoteHandler) GetQuoteByID(context *gin.Context) {

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

	var quote repository.Quote
	var err error
	if claims.Role == constant.CUSTOMER_ROLE {
		quote, err = handler.queries.GetQuote(context, repository.GetQuoteParams{
			ID:         id,
			CustomerID: claims.CustomerID,
		})

	} else {
		quote, err = handler.queries.AdminGetQuote(context, id)
	}

	if err != nil {
		fmt.Println(err.Error())
		response.ErrorResponse(context, http.StatusNotFound, err.Error())
		return
	}

	item := ListAllQuotesResult{
		ID:              quote.ID,
		QuoteType:       quote.QuoteType,
		CustomerID:      quote.CustomerID,
		Source:          quote.Source,
		Destination:     quote.Destination,
		DoorPickup:      quote.DoorPickup.String,
		DoorAddress:     quote.DoorAddress.String,
		DoorDelivery:    quote.DoorDelivery.String,
		DeliveryAddress: quote.DeliveryAddress.String,
		LinerID:         quote.LinerID.String,
		PartnerID:       quote.PartnerID.String,
		Validity:        quote.Validity.String,
		TransmitDays:    quote.TransmitDays.String,
		FreeDays:        quote.FreeDays.String,
		Currency:        quote.Currency.String,
		Buy:             quote.Buy.String,
		Sell:            quote.Sell.String,
		PartnerTax:      quote.PartnerTax.String,
		QuoteStatus:     quote.QuoteStatus.String,
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Quote found",
		"data":    item,
	})
}

func (handler QuoteHandler) UpdateQuote(context *gin.Context) {

	var tokenData interface{}
	tokenData, isExists := context.Get("claims")
	if !isExists {
		response.ErrorResponse(context, http.StatusUnauthorized, "Claims not found in request, request unauthorised")
		return
	}
	claims := tokenData.(model.TokenData)
	if claims.Role == constant.CUSTOMER_ROLE {
		response.ErrorResponse(context, http.StatusUnauthorized, "Permission denied")
		return
	}
	var updateQuoteRequest model.UpdateQuoteRequest

	if err := context.ShouldBind(&updateQuoteRequest); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		fmt.Println("Error:  ", err)
		return
	}

	var data repository.Quote
	var err error

	data, err = handler.queries.AdminGetQuote(context, updateQuoteRequest.ID)

	if err != nil {
		fmt.Println("Error:  ", err)
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "success",
			"message": "Quote not found",
		})
		return
	}

	if updateQuoteRequest.Buy != data.Buy.String {
		if claims.Role != constant.PROCUREMENT_ROLE && claims.Role != constant.ADMIN_ROLE {
			response.ErrorResponse(context, http.StatusUnauthorized, "Only admin or procurement executive can set Buy rate")
			return
		} else {
			data.Buy.String = updateQuoteRequest.Buy
			data.Buy.Valid = true
		}
	}

	if updateQuoteRequest.Sell != data.Sell.String {
		if claims.Role != constant.SALE_ROLE && claims.Role != constant.ADMIN_ROLE {
			response.ErrorResponse(context, http.StatusUnauthorized, "Only admin or sales executive can set sale rate")
			return
		} else {
			data.Sell.String = updateQuoteRequest.Sell
			data.Sell.Valid = true
		}
	}

	if updateQuoteRequest.LinerId != "" {
		fmt.Println("Updated LinerId: ", updateQuoteRequest.LinerId)
		data.LinerID.String = updateQuoteRequest.LinerId
		data.LinerID.Valid = true
	}
	if updateQuoteRequest.PartnerId != "" {
		fmt.Println("Updated LinerId: ", updateQuoteRequest.PartnerId)
		data.PartnerID.String = updateQuoteRequest.PartnerId
		data.PartnerID.Valid = true
	}
	if updateQuoteRequest.Validity != "" {
		fmt.Println("Updated LinerId: ", updateQuoteRequest.Validity)
		data.Validity.String = updateQuoteRequest.Validity
		data.Validity.Valid = true
	}
	if updateQuoteRequest.TransmitDays != "" {
		fmt.Println("Updated LinerId: ", updateQuoteRequest.TransmitDays)
		data.TransmitDays.String = updateQuoteRequest.TransmitDays
		data.TransmitDays.Valid = true
	}
	if updateQuoteRequest.FreeDays != "" {
		fmt.Println("Updated LinerId: ", updateQuoteRequest.FreeDays)
		data.FreeDays.String = updateQuoteRequest.FreeDays
		data.FreeDays.Valid = true
	}
	if updateQuoteRequest.Currency != "" {
		fmt.Println("Updated LinerId: ", updateQuoteRequest.Currency)
		data.Currency.String = updateQuoteRequest.Currency
		data.Currency.Valid = true
	}

	if updateQuoteRequest.CustomerId != "" {
		fmt.Println(" Customer id: ", updateQuoteRequest.Currency)
		data.CustomerID = updateQuoteRequest.CustomerId
		data.Currency.Valid = true
	}

	finalQuote := repository.AdminUpdateQuoteParams{
		ID:           data.ID,
		LinerID:      data.LinerID,
		PartnerID:    data.PartnerID,
		Validity:     data.Validity,
		TransmitDays: data.TransmitDays,
		FreeDays:     data.FreeDays,
		Currency:     data.Currency,
		Buy:          data.Buy,
		Sell:         data.Sell,
		PartnerTax:   data.PartnerTax,
	}

	err = handler.queries.AdminUpdateQuote(context, finalQuote)

	if err != nil {
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "failed",
			"message": "Error updating quote in database",
		})
		return
	}

	item := ListAllQuotesResult{
		ID:              data.ID,
		QuoteType:       data.QuoteType,
		CustomerID:      data.CustomerID,
		Source:          data.Source,
		Destination:     data.Destination,
		DoorPickup:      data.DoorPickup.String,
		DoorAddress:     data.DoorAddress.String,
		DoorDelivery:    data.DoorDelivery.String,
		DeliveryAddress: data.DeliveryAddress.String,
		LinerID:         data.LinerID.String,
		PartnerID:       data.PartnerID.String,
		Validity:        data.Validity.String,
		TransmitDays:    data.TransmitDays.String,
		FreeDays:        data.FreeDays.String,
		Currency:        data.Currency.String,
		Buy:             data.Buy.String,
		Sell:            data.Sell.String,
		PartnerTax:      data.PartnerTax.String,
		QuoteStatus:     data.QuoteStatus.String,
	}

	// TODO return, quote updated successfuly
	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Quote updated successfuly",
		"data":    item,
	})
}

func (handler QuoteHandler) GetAllQuote(context *gin.Context) {
	var tokenData interface{}
	tokenData, isExists := context.Get("claims")
	if !isExists {
		response.ErrorResponse(context, http.StatusUnauthorized, "Claims not found in request, request unauthorised")
		return
	}

	claims := tokenData.(model.TokenData)

	var quotes []repository.Quote
	var err error
	if claims.Role == constant.CUSTOMER_ROLE {
		quotes, err = handler.queries.ListQuotes(context, claims.CustomerID)
	} else {
		quotes, err = handler.queries.AdminListQuotes(context)
	}

	if err != nil {
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "success",
			"message": "Error int get all quote",
			"error":   err,
		})
		return
	}

	var dataMap []ListAllQuotesResult
	for _, quote := range quotes {
		item := ListAllQuotesResult{
			ID:              quote.ID,
			QuoteType:       quote.QuoteType,
			CustomerID:      quote.CustomerID,
			Source:          quote.Source,
			Destination:     quote.Destination,
			DoorPickup:      quote.DoorPickup.String,
			DoorAddress:     quote.DoorAddress.String,
			DoorDelivery:    quote.DoorDelivery.String,
			DeliveryAddress: quote.DeliveryAddress.String,
			LinerID:         quote.LinerID.String,
			PartnerID:       quote.PartnerID.String,
			Validity:        quote.Validity.String,
			TransmitDays:    quote.TransmitDays.String,
			FreeDays:        quote.FreeDays.String,
			Currency:        quote.Currency.String,
			Buy:             quote.Buy.String,
			Sell:            quote.Sell.String,
			PartnerTax:      quote.PartnerTax.String,
			QuoteStatus:     quote.QuoteStatus.String,
		}
		dataMap = append(dataMap, item)
	}
	result := map[string]interface{}{
		"code":    "success",
		"message": "Fetched all quote list",
		"data":    dataMap,
	}

	response.SuccessResponse(context, result)
}

type ListAllQuotesResult struct {
	ID              string `json:"id"`
	QuoteType       string `json:"quote_type"`
	CustomerID      string `json:"customer_id"`
	Source          string `json:"source"`
	Destination     string `json:"destination"`
	DoorPickup      string `json:"door_pickup"`
	DoorAddress     string `json:"door_address"`
	DoorDelivery    string `json:"door_delivery"`
	DeliveryAddress string `json:"delivery_address"`
	LinerID         string `json:"liner_id"`
	PartnerID       string `json:"partner_id"`
	Validity        string `json:"validity"`
	TransmitDays    string `json:"transmit_days"`
	FreeDays        string `json:"free_days"`
	Currency        string `json:"currency"`
	Buy             string `json:"buy"`
	Sell            string `json:"sell"`
	PartnerTax      string `json:"partner_tax"`
	QuoteStatus     string `json:"quote_status"`
}

func (handler QuoteHandler) ConfirmQuote(context *gin.Context) {

	var tokenData interface{}
	tokenData, isExists := context.Get("claims")
	if !isExists {
		response.ErrorResponse(context, http.StatusUnauthorized, "Claims not found in request, request unauthorised")
		return
	}
	claims := tokenData.(model.TokenData)
	if claims.Role == constant.CUSTOMER_ROLE {
		response.ErrorResponse(context, http.StatusUnauthorized, "Permission denied")
		return
	}
	var confirmQuote model.ConfirmQuoteRequest

	if err := context.ShouldBind(&confirmQuote); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		fmt.Println("Error:  ", err)
		return
	}

	var data repository.Quote
	var err error

	if claims.Role == constant.CUSTOMER_ROLE {
		data, err = handler.queries.GetQuote(context, repository.GetQuoteParams{
			ID:         confirmQuote.QuoteID,
			CustomerID: claims.CustomerID,
		})
	} else {
		data, err = handler.queries.AdminGetQuote(context, confirmQuote.QuoteID)
	}

	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Quote not found")
		return
	}

	if data.QuoteStatus.String == constant.QUOTE_CONFIRMED {
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "success",
			"message": "Quote already confirmed",
		})
		return
	}
	if data.Buy.String == "" || data.Sell.String == "" {
		response.ErrorResponse(context, http.StatusNotFound, "Procurement or seller values not attached with quote")
		return
	}

	booking, err := handler.queries.CreateBooking(context, repository.CreateBookingParams{
		ID:               uuid.New().String(),
		BookingRequestID: uuid.New().String(),
		BookingStatus:    sql.NullString{String: constant.BOOKING_CREATED, Valid: true},
		CustomerID:       sql.NullString{String: claims.CustomerID, Valid: true},
		Source:           sql.NullString{String: data.Source, Valid: true},
		Destination:      sql.NullString{String: data.Destination, Valid: true},
	})
	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Unable to create booking")
		return
	}

	_, err = handler.queries.CreateBookingMilestone(context, repository.CreateBookingMilestoneParams{
		ID:              uuid.New().String(),
		BookingID:       sql.NullString{String: booking.ID, Valid: true},
		MilestoneName:   sql.NullString{String: "Booking Created", Valid: true},
		MilestoneStatus: sql.NullString{String: constant.MILESTONE_COMPLETED, Valid: true},
		CreatedAt:       sql.NullString{String: time.Now().UTC().String(), Valid: true},
		CompletedAt:     sql.NullString{Valid: false},
	})
	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Unable to create Milestone 1")
		return
	}
	_, err = handler.queries.CreateBookingMilestone(context, repository.CreateBookingMilestoneParams{
		ID:              uuid.New().String(),
		BookingID:       sql.NullString{String: booking.ID, Valid: true},
		MilestoneName:   sql.NullString{String: "Booking Confirmed", Valid: true},
		MilestoneStatus: sql.NullString{String: constant.MILESTONE_COMPLETED, Valid: true},
		CreatedAt:       sql.NullString{String: time.Now().UTC().String(), Valid: true},
		CompletedAt:     sql.NullString{Valid: false},
	})
	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Unable to create Milestone 2")
		return
	}
	_, err = handler.queries.CreateBookingMilestone(context, repository.CreateBookingMilestoneParams{
		ID:              uuid.New().String(),
		BookingID:       sql.NullString{String: booking.ID, Valid: true},
		MilestoneName:   sql.NullString{String: "Vessel Departured", Valid: true},
		MilestoneStatus: sql.NullString{String: constant.MILESTONE_IN_PROGRESS, Valid: true},
		CreatedAt:       sql.NullString{String: time.Now().UTC().String(), Valid: true},
		CompletedAt:     sql.NullString{Valid: false},
	})
	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Unable to create Milestone 3")
		return
	}

	_, err = handler.queries.CreateBookingMilestone(context, repository.CreateBookingMilestoneParams{
		ID:              uuid.New().String(),
		BookingID:       sql.NullString{String: booking.ID, Valid: true},
		MilestoneName:   sql.NullString{String: "Vessel Arrived", Valid: true},
		MilestoneStatus: sql.NullString{String: constant.MILESTONE_IN_PROGRESS, Valid: true},
		CreatedAt:       sql.NullString{String: time.Now().UTC().String(), Valid: true},
		CompletedAt:     sql.NullString{Valid: false},
	})
	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Unable to create Milestone 4")
		return
	}

	_, err = handler.queries.CreateBookingMilestone(context, repository.CreateBookingMilestoneParams{
		ID:              uuid.New().String(),
		BookingID:       sql.NullString{String: booking.ID, Valid: true},
		MilestoneName:   sql.NullString{String: "Booking Complete", Valid: true},
		MilestoneStatus: sql.NullString{String: constant.MILESTONE_IN_PROGRESS, Valid: true},
		CreatedAt:       sql.NullString{String: time.Now().UTC().String(), Valid: true},
		CompletedAt:     sql.NullString{Valid: false},
	})
	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Unable to create Milestone 5")
		return
	}

	_, err = handler.queries.CreateBookingTask(context, repository.CreateBookingTaskParams{
		ID:          uuid.New().String(),
		BookingID:   sql.NullString{String: booking.ID, Valid: true},
		TaskName:    sql.NullString{String: "Generate Invoice", Valid: true},
		TaskStatus:  sql.NullString{String: constant.TASK_IN_PROGRESS, Valid: true},
		CreatedAt:   sql.NullString{String: time.Now().UTC().String(), Valid: true},
		CompletedAt: sql.NullString{Valid: false},
	})
	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Unable to create Task 1")
		return
	}

	_, err = handler.queries.CreateBookingTask(context, repository.CreateBookingTaskParams{
		ID:          uuid.New().String(),
		BookingID:   sql.NullString{String: booking.ID, Valid: true},
		TaskName:    sql.NullString{String: "Create HBL", Valid: true},
		TaskStatus:  sql.NullString{String: constant.TASK_IN_PROGRESS, Valid: true},
		CreatedAt:   sql.NullString{String: time.Now().UTC().String(), Valid: true},
		CompletedAt: sql.NullString{Valid: false},
	})
	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Unable to create Task 2")
		return
	}

	_, err = handler.queries.CreateBookingTask(context, repository.CreateBookingTaskParams{
		ID:          uuid.New().String(),
		BookingID:   sql.NullString{String: booking.ID, Valid: true},
		TaskName:    sql.NullString{String: "Upload Documents", Valid: true},
		TaskStatus:  sql.NullString{String: constant.TASK_IN_PROGRESS, Valid: true},
		CreatedAt:   sql.NullString{String: time.Now().UTC().String(), Valid: true},
		CompletedAt: sql.NullString{Valid: false},
	})
	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Unable to create Task 3")
		return
	}

	err = handler.queries.AdminConfirmQuote(context, repository.AdminConfirmQuoteParams{
		QuoteStatus: sql.NullString{String: constant.QUOTE_CONFIRMED, Valid: true},
		ID:          confirmQuote.QuoteID,
	})

	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Unable to Quote status to ")
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":       "success",
		"message":    "Quote Confirm successfully",
		"booking_id": booking.ID,
	})

}
