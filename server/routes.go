package server

import (
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/handler"
)

func ConfigureRoutes(server *Server) {

	quotesHandler := handler.NewQuoteHandler(server.db)
	api := server.engine.RouterGroup.Group("/api")
	api.POST("/quote", quotesHandler.RequestQuote)

	// api.PUT("/update_quote", quotesHandler.UpdateQuote)
	// api.GET("/quote", quotesHandler.GetQuoteByID)
	// api.GET("/quotes", quotesHandler.GetAllQuote)

}
