package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/handler"
)

func ConfigureRoutes(server *Server) {

	quotesHandler := handler.NewQuoteHandler(server.db)
	accountHandler := handler.NewAccountHandler(server.db)
	bookingHandler := handler.NewBookingHandler(server.db)
	bookingMilestoneHandler := handler.NewBookingMilestoneHandler(server.db)
	bookingTaskHandler := handler.NewBookingTaskHandler(server.db)

	api := server.engine.RouterGroup.Group("/api")

	// quotes APIs
	api.POST("/quote", quotesHandler.RequestQuote)
	api.PUT("/update_quote", quotesHandler.UpdateQuote)
	api.GET("/quote", quotesHandler.GetQuoteByID)
	api.GET("/quotes", quotesHandler.GetAllQuote)

	// account APIs
	api.POST("/login", accountHandler.Login)
	api.POST("/Register", accountHandler.CreateAccount)
	api.PUT("/update_account", accountHandler.UpdateAccountDetails)
	api.GET("/account", accountHandler.GetAccountByID)
	api.GET("/accounts", accountHandler.GetAllAccount)

	// booking APIs
	api.POST("/booking", bookingHandler.CreateBooking)
	api.PUT("/update_booking", bookingHandler.UpdateBooking)
	api.GET("/booking", bookingHandler.GetBookingByID)
	api.GET("/bookings", bookingHandler.ListBookings)

	// booking milestone APIs

	api.POST("/booking_milestone", bookingMilestoneHandler.CreateBookingMilestone)
	api.PUT("/update_booking_milestone", bookingMilestoneHandler.UpdateBookingMilestone)
	api.GET("/booking_milestone", bookingMilestoneHandler.GetBookingMilestone)
	api.GET("/booking_milestones", bookingMilestoneHandler.ListBookingMilestone)

	// booking task APIS

	api.POST("/booking_task", bookingTaskHandler.CreateBookingTask)
	api.PUT("/update_booking_task", bookingTaskHandler.UpdateBookingTask)
	api.GET("/booking_task", bookingTaskHandler.GetBookingTask)
	api.GET("/booking_tasks", bookingTaskHandler.ListBookingTask)

	api.POST("/customers", func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Printf("ERROR: %s", err.Error())
			c.JSON(http.StatusInternalServerError, map[string]string{
				"ERROR": err.Error(),
			})
			return
		}
		fmt.Printf("DATA: %s", string(data))
		c.JSON(http.StatusInternalServerError, map[string]string{
			"DATA": string(data),
		})
	})

}
