package handler

import (
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/constant"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
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

func (handler BookingMilestoneHandler) GetBookingMilestone(context *gin.Context) {

	var tokenData interface{}
	tokenData, isExists := context.Get("claims")
	if !isExists {
		response.ErrorResponse(context, http.StatusUnauthorized, "Claims not found in request, request unauthorised")
		return
	}

	claims := tokenData.(model.TokenData)
	if claims.Role == constant.CUSTOMER_ROLE {
		response.ErrorResponse(context, http.StatusBadRequest, "Customer can't get milestone")
		return
	}

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

type UpdateBookingMilestoneRequest struct {
	ID              string `json:"id"`
	MilestoneStatus string `json:"milestone_status"`
	CompletedAt     string `json:"completed_at"`
}

func (handler BookingMilestoneHandler) UpdateBookingMilestone(context *gin.Context) {

	var tokenData interface{}
	tokenData, isExists := context.Get("claims")
	if !isExists {
		response.ErrorResponse(context, http.StatusUnauthorized, "Claims not found in request, request unauthorised")
		return
	}

	claims := tokenData.(model.TokenData)

	if claims.Role == constant.CUSTOMER_ROLE {
		response.ErrorResponse(context, http.StatusBadRequest, "Customer can't update milestone")
		return
	}

	var bookingMilestone UpdateBookingMilestoneRequest

	if err := context.ShouldBind(&bookingMilestone); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}

	err := handler.queries.UpdateBookingMilestone(context, repository.UpdateBookingMilestoneParams{
		ID:              bookingMilestone.ID,
		MilestoneStatus: sql.NullString{String: bookingMilestone.MilestoneStatus, Valid: true},
		CompletedAt:     sql.NullString{String: bookingMilestone.CompletedAt, Valid: true},
	})

	if err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Unable to update record")
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Updated suceessfully",
	})
}
