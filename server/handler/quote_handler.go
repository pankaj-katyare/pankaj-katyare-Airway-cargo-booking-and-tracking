package handler

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
	quotesDB "github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/repository/quotes"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/response"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/service"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type QuoteHandler struct {
	DB *sqlx.DB
}

func (handler QuoteHandler) RequestQuote(context *gin.Context) {

	var quote model.Quote

	if err := context.ShouldBind(&quote); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}
	claims := jwt.ExtractClaims(context)
	id := claims["id"]

	if id == nil {
		id = uuid.New()
	}

	quoteService := service.QuoteService{DB: handler.DB}
	quotes := quoteService.RequestQuote(quote)
	quoteRepository := quotesDB.NewConnection(&quotesDB.Config{DB: handler.DB})
	quoteRepository.Insert(context, quotes)
	response.SuccessResponse(context, quotes)

}

func (handler QuoteHandler) GetQuoteByID(context *gin.Context) {
	quoteService := service.QuoteService{DB: handler.DB}
	quote := model.Quote{}
	id, _ := strconv.Atoi(context.Param("id"))
	quoteRepository := quotesDB.NewConnection(&quotesDB.Config{DB: handler.DB})
	quoteRepository.Select(id, &quote)

	if post.ID == 0 {
		response.ErrorResponse(context, http.StatusNotFound, "Post not found")
		return
	}

	response.SuccessResponse(context, response.GetPostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	})
}
