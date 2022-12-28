package server

import (
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/handler"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/provider"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	quotesHandler := handler.QuoteHandler{DB: server.db}

	jwtAuth := provider.NewJwtAuth(server.db)

	server.engine.POST("/login", jwtAuth.Middleware().LoginHandler)
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	needsAuth := server.engine.Group("/").Use(jwtAuth.Middleware().MiddlewareFunc())
	needsAuth.GET("/refresh", jwtAuth.Middleware().RefreshHandler)

	needsAuth.POST("/quote", quotesHandler.RequestQuote)
}
