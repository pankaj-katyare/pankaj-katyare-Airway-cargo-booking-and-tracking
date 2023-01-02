package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

	var updateQuoteRequest repository.UpdateQuoteParams

	if err := context.ShouldBind(&updateQuoteRequest); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}

	data, err := handler.queries.GetQuote(context, repository.GetQuoteParams{
		ID: updateQuoteRequest.ID,
	})
	if err != nil {
		// TODO: return quote not found
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "success",
			"message": "Quote not found",
		})
	}

	var isUpdated bool
	updateBody := make(map[string]interface{})

	if updateQuoteRequest.Buy.String != "" {
		updateBody["buy"] = updateQuoteRequest.Buy
		updateBody["liner_id"] = updateQuoteRequest.LinerID
		updateBody["partner_id"] = updateQuoteRequest.PartnerID
		updateBody["validity"] = updateQuoteRequest.Validity
		updateBody["transmit_days"] = updateQuoteRequest.TransmitDays
		updateBody["free_days"] = updateQuoteRequest.FreeDays
		updateBody["currency"] = updateQuoteRequest.Currency
		updateBody["partner_tax"] = updateQuoteRequest.PartnerTax
		updateBody["procurement_id"] = updateQuoteRequest.ProcurementID
	}

	if updateQuoteRequest.Sell.String != "" {
		if data.Buy.String != "" {
			//TODO: return Buydate should be set first
		} else {
			updateBody["sell"] = updateQuoteRequest.Sell
			updateBody["sales_id"] = updateQuoteRequest.SalesID
		}
	}

	if isUpdated {
		err := handler.queries.UpdateQuote(context, repository.UpdateQuoteParams{
			LinerID:       updateQuoteRequest.LinerID,
			PartnerID:     updateQuoteRequest.PartnerID,
			Validity:      updateQuoteRequest.Validity,
			TransmitDays:  updateQuoteRequest.TransmitDays,
			FreeDays:      updateQuoteRequest.FreeDays,
			Currency:      updateQuoteRequest.Currency,
			Buy:           updateQuoteRequest.Buy,
			Sell:          updateQuoteRequest.Sell,
			PartnerTax:    updateQuoteRequest.PartnerTax,
			ProcurementID: updateQuoteRequest.ProcurementID,
			SalesID:       updateQuoteRequest.SalesID,
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
			"message": "Quote updated successfuly",
			"data":    data,
		})
	}
	// TODO return, nothing to update
	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Nothing to process",
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
