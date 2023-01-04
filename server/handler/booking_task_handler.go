package handler

import (
	"database/sql"
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

type UpdateBookingTaskRequest struct {
	ID          string `json:"id"`
	TaskStatus  string `json:"task_status"`
	CompletedAt string `json:"completed_at"`
}

func (handler BookingTaskHandler) UpdateBookingTask(context *gin.Context) {

	var bookingTask UpdateBookingTaskRequest

	if err := context.ShouldBind(&bookingTask); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}

	state := handler.queries.UpdateBookingTask(context, repository.UpdateBookingTaskParams{
		ID:          uuid.New().String(),
		TaskStatus:  sql.NullString{String: bookingTask.TaskStatus, Valid: true},
		CompletedAt: sql.NullString{String: bookingTask.CompletedAt, Valid: true},
	})

	// TODO return, nothing to update
	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Updated suceessfully",
		"data":    state,
	})
}
