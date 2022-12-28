package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/docs"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/db"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	connection := db.InitDB()

	defer func() {
		if err := connection.DB.Close(); err != nil {
			log.Print(err)
		}
	}()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("EXPOSE_PORT"))

	app := server.NewServer(connection)
	server.ConfigureRoutes(app)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}
