package handler

import (
	"database/sql"
	"fmt"
	"net/http"

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

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Request Quote Created successfuly",
		"data":    state,
	})
}

func (handler QuoteHandler) GetQuoteByID(context *gin.Context) {

	id := context.Request.URL.Query().Get("id")
	customer_id := context.Request.URL.Query().Get("customer_id")

	if id == "" {
		response.ErrorResponse(context, http.StatusNotFound, "ID not specified")
		return
	}
	if customer_id == "" {
		response.ErrorResponse(context, http.StatusNotFound, "customer_id not specified")
		return
	}

	state, err := handler.queries.GetQuote(context, repository.GetQuoteParams{
		ID:         id,
		CustomerID: customer_id,
	})

	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "quote Data",
		"data":    state,
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

	if claims.Role == constant.CUSTOMER_ROLE {
		data, err = handler.queries.GetQuote(context, repository.GetQuoteParams{
			ID:         updateQuoteRequest.ID,
			CustomerID: claims.CustomerID,
		})
	} else {
		data, err = handler.queries.AdminGetQuote(context, updateQuoteRequest.ID)
	}

	if err != nil {
		// TODO: return quote not found
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
		data.LinerID.String = updateQuoteRequest.LinerId
		data.LinerID.Valid = true
	}
	if updateQuoteRequest.PartnerId != "" {
		data.PartnerID.String = updateQuoteRequest.PartnerId
		data.PartnerID.Valid = true
	}
	if updateQuoteRequest.Validity != "" {
		data.Validity.String = updateQuoteRequest.Validity
		data.Validity.Valid = true
	}
	if updateQuoteRequest.TransmitDays != "" {
		data.TransmitDays.String = updateQuoteRequest.TransmitDays
		data.TransmitDays.Valid = true
	}
	if updateQuoteRequest.FreeDays != "" {
		data.FreeDays.String = updateQuoteRequest.FreeDays
		data.FreeDays.Valid = true
	}
	if updateQuoteRequest.Currency != "" {
		data.Currency.String = updateQuoteRequest.Currency
		data.Currency.Valid = true
	}
	if updateQuoteRequest.PartnerTax != "" {
		data.PartnerTax.String = updateQuoteRequest.PartnerTax
		data.PartnerTax.Valid = true
	}

	err = handler.queries.UpdateQuote(context, repository.UpdateQuoteParams{
		LinerID:      data.LinerID,
		PartnerID:    data.PartnerID,
		Validity:     data.Validity,
		TransmitDays: data.TransmitDays,
		FreeDays:     data.FreeDays,
		Currency:     data.Currency,
		Buy:          data.Buy,
		Sell:         data.Sell,
		PartnerTax:   data.PartnerTax,
	})

	if err != nil {
		// TODO: return, error updating quote in database
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "success",
			"message": "Error updating quote in database",
		})
		return
	}
	// TODO return, quote updated successfuly
	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Quote updated successfuly",
		"data":    data,
	})
}

func (handler QuoteHandler) GetAllQuote(context *gin.Context) {

	customer_id := context.Request.URL.Query().Get("customer_id")

	quotes, err := handler.queries.ListQuotes(context, customer_id)

	if err != nil {
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "success",
			"message": "Error int get all quote",
			"error":   err,
		})
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Fetched all quote list",
		"data":    quotes,
	})
}
