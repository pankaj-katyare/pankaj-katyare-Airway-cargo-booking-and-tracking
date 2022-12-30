package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/repository"
)

type BookingHandler struct {
	DB      *sqlx.DB
	queries *repository.Queries
}

func NewHandler(DB *sqlx.DB) *BookingHandler {
	return &BookingHandler{
		DB:      DB,
		queries: repository.New(DB),
	}
}

// func (handler BookingHandler) CreateBooking(context *gin.Context) {

// 	var booking model.Booking

// 	if err := context.ShouldBind(&booking); err != nil {
// 		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
// 		return
// 	}

// 	handler.bookingsDB.Insert(context, booking)
// 	response.SuccessResponse(context, booking)

// }

// func (handler BookingHandler) GetBookingByID(context *gin.Context) {

// 	id := context.Param("id")
// 	if id == "" {
// 		response.ErrorResponse(context, http.StatusNotFound, " id is invalid ")
// 		return
// 	}

// 	bookings, err := handler.bookingsDB.Select(context, map[string]interface{}{
// 		"id": id,
// 	}, nil, nil)

// 	if err == nil {
// 		response.ErrorResponse(context, http.StatusNotFound, err.Error())
// 		return
// 	}

// 	response.SuccessResponse(context, map[string]interface{}{
// 		"code":    "success",
// 		"message": "Quote Data",
// 		"data":    bookings[0],
// 	})
// }

// func (handler BookingHandler) UpdateBooking(context *gin.Context) {

// 	var quote model.Quote

// 	if err := context.ShouldBind(&quote); err != nil {
// 		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
// 		return
// 	}
// 	claims := jwt.ExtractClaims(context)
// 	id := claims["id"]

// 	if id == nil {
// 		id = uuid.New()
// 	}

// 	quoteService := service.QuoteService{DB: handler.DB}
// 	quotes := quoteService.UpdateQuote(quote)
// 	// quoteRepository := bookingsDB.NewConnection(&bookingsDB.Config{DB: handler.DB})
// 	// quoteRepository.Insert(context, quotes)
// 	response.SuccessResponse(context, quotes)

// }
