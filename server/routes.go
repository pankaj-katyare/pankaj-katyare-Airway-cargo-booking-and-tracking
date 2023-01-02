package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/handler"
)

func ConfigureRoutes(server *Server) {

	quotesHandler := handler.NewQuoteHandler(server.db)
	accountHandler := handler.NewAccountHandler(server.db)
	bookingHandler := handler.NewBookingHandler(server.db)
	bookingMilestoneHandler := handler.NewBookingMilestoneHandler(server.db)
	bookingTaskHandler := handler.NewBookingTaskHandler(server.db)

	api := server.engine.RouterGroup.Group("/api")

	// Request
	api.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.New().String())
	})

	// Auth & register apis
	api.POST("/login", accountHandler.Login)
	api.POST("/Register", accountHandler.CreateAccount)

	auth := api.Group("/cargo")

	// Authentication middleware
	auth.Use(handler.AuthorizationMiddleware)

	// Account APIs
	auth.PUT("/update_account", accountHandler.UpdateAccountDetails)
	auth.GET("/account", accountHandler.GetAccountByID)
	auth.GET("/accounts", accountHandler.GetAllAccount)

	// Quotes APIs
	auth.POST("/quote", quotesHandler.RequestQuote)
	auth.PUT("/update_quote", quotesHandler.UpdateQuote)
	auth.GET("/quote", quotesHandler.GetQuoteByID)
	auth.GET("/quotes", quotesHandler.GetAllQuote)

	// Booking APIs
	auth.POST("/booking", bookingHandler.CreateBooking)
	auth.PUT("/update_booking", bookingHandler.UpdateBooking)
	auth.GET("/booking", bookingHandler.GetBookingByID)
	auth.GET("/bookings", bookingHandler.ListBookings)

	// Booking milestone APIs
	auth.POST("/booking_milestone", bookingMilestoneHandler.CreateBookingMilestone)
	auth.PUT("/update_booking_milestone", bookingMilestoneHandler.UpdateBookingMilestone)
	auth.GET("/booking_milestone", bookingMilestoneHandler.GetBookingMilestone)
	auth.GET("/booking_milestones", bookingMilestoneHandler.ListBookingMilestone)

	// Booking task APIS
	auth.POST("/booking_task", bookingTaskHandler.CreateBookingTask)
	auth.PUT("/update_booking_task", bookingTaskHandler.UpdateBookingTask)
	auth.GET("/booking_task", bookingTaskHandler.GetBookingTask)
	auth.GET("/booking_tasks", bookingTaskHandler.ListBookingTask)

}
