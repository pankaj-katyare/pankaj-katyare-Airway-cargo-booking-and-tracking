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
	api := server.engine.RouterGroup.Group("/api")
	api.POST("/quote", quotesHandler.RequestQuote)

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

	// api.PUT("/update_quote", quotesHandler.UpdateQuote)
	// api.GET("/quote", quotesHandler.GetQuoteByID)
	// api.GET("/quotes", quotesHandler.GetAllQuote)

}
