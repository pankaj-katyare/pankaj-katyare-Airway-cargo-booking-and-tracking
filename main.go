package main

import (
	"log"
	"os"

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

	app := server.NewServer(connection)
	server.ConfigureRoutes(app)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}
