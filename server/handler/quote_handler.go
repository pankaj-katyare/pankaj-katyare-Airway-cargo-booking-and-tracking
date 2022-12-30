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

// func (handler QuoteHandler) GetQuoteByID(context *gin.Context) {

// 	id := context.Request.URL.Query().Get("id")
// 	if id == "" {
// 		response.ErrorResponse(context, http.StatusNotFound, "ID not specified")
// 		return
// 	}

// 	state, err := handler.queries.GetQuote(context,database.GetQuoteParams{ID: "", CustomerID: ""})

// 	quotes, err := handler.quotesDB.Select(context, map[string]interface{}{
// 		"id": id,
// 	}, nil, nil)

// 	if err != nil {
// 		response.ErrorResponse(context, http.StatusNotFound, err.Error())
// 		return
// 	}

// 	response.SuccessResponse(context, map[string]interface{}{
// 		"code":    "success",
// 		"message": "Quote Data",
// 		"data":    quotes[0],
// 	})
// }

// type UpdateQuoteRequest struct {
// 	QuoteId      string `json:"id" form:"id" binding:"required"`
// 	SellRate     string `json:"sell" form:"sell" binding:"required"`
// 	BuyRate      string `json:"buy" form:"buy" binding:"required"`
// 	LinerId      string `json:"liner_id" form:"liner_id" binding:"required"`
// 	PartnerId    string `json:"partner_id" form:"partner_id" binding:"required"`
// 	Validity     string `json:"validity" form:"validity" binding:"required"`
// 	TransmitDays string `json:"transmit_days" form:"transmit_days" binding:"required"`
// 	FreeDays     string `json:"free_days" form:"free_days" binding:"required"`
// 	Currency     string `json:"currency" form:"currency" binding:"required"`
// 	PartnerTax   string `json:"partner_tax" form:"partner_tax" binding:"required"`
// }

// func (handler QuoteHandler) UpdateQuote(context *gin.Context) {

// 	var updateQuoteRequest UpdateQuoteRequest

// 	if err := context.ShouldBind(&updateQuoteRequest); err != nil {
// 		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
// 		return
// 	}

// 	data, err := handler.quotesDB.Select(context, map[string]interface{}{
// 		"id": updateQuoteRequest.QuoteId,
// 	}, nil, nil)

// 	if err != nil {
// 		// TODO: return quote not found
// 		response.SuccessResponse(context, map[string]interface{}{
// 			"code":    "success",
// 			"message": "Quote not found",
// 		})
// 	}
// 	if len(data) == 0 {
// 		// TODO: return quote not found
// 		response.SuccessResponse(context, map[string]interface{}{
// 			"code":    "success",
// 			"message": "Quote not found",
// 		})
// 	} else if len(data) > 1 {
// 		// TODO: internal server error, more than one quotes with same id
// 		response.SuccessResponse(context, map[string]interface{}{
// 			"code":    "success",
// 			"message": "internal server error, more than one quotes with same id",
// 		})
// 	}
// 	var isUpdated bool
// 	updateBody := make(map[string]interface{})

// 	if updateQuoteRequest.BuyRate != "" {
// 		updateBody["buy"] = updateQuoteRequest.BuyRate
// 		updateBody["liner_id"] = updateQuoteRequest.LinerId
// 		updateBody["partner_id"] = updateQuoteRequest.PartnerId
// 		updateBody["validity"] = updateQuoteRequest.Validity
// 		updateBody["transmit_days"] = updateQuoteRequest.TransmitDays
// 		updateBody["free_days"] = updateQuoteRequest.FreeDays
// 		updateBody["currency"] = updateQuoteRequest.Currency
// 		updateBody["partner_tax"] = updateQuoteRequest.PartnerId
// 	}

// 	if updateQuoteRequest.SellRate != "" {
// 		if data[0].Buy != "" {
// 			//TODO: return Buydate should be set first
// 		} else {
// 			updateBody["sell"] = updateQuoteRequest.SellRate
// 		}
// 	}

// 	if isUpdated {
// 		_, err := handler.quotesDB.Update(context, updateBody, []string{
// 			fmt.Sprintf("id = '%s'", updateQuoteRequest.QuoteId),
// 		})
// 		if err != nil {
// 			// TODO: return, error updating quote in database
// 			response.SuccessResponse(context, map[string]interface{}{
// 				"code":    "success",
// 				"message": "Error updating quote in database",
// 			})
// 		}
// 		// TODO return, quote updated successfuly
// 		response.SuccessResponse(context, map[string]interface{}{
// 			"code":    "success",
// 			"message": "Quote updated successfuly",
// 			"data":    data[0],
// 		})
// 	}
// 	// TODO return, nothing to update
// 	response.SuccessResponse(context, map[string]interface{}{
// 		"code":    "success",
// 		"message": "Nothing to process",
// 	})
// }

// func (handler QuoteHandler) GetAllQuote(context *gin.Context) {

// 	quotes, err := handler.quotesDB.Select(context, map[string]interface{}{}, nil, nil)

// 	if err != nil {
// 		response.SuccessResponse(context, map[string]interface{}{
// 			"code":    "success",
// 			"message": "Error int get all quote",
// 			"error":   err,
// 		})
// 		return
// 	}

// 	response.SuccessResponse(context, map[string]interface{}{
// 		"code":    "success",
// 		"message": "Fetched all quote list",
// 		"data":    quotes,
// 	})
// }
