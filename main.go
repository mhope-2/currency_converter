package main

import (
	"log"
	"os"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mhope-2/currency_converter/database"
	"github.com/mhope-2/currency_converter/database/models"
	"github.com/mhope-2/currency_converter/database/postgres"
	"github.com/mhope-2/currency_converter/handler"
	"github.com/mhope-2/currency_converter/server"
)

func main() {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	ENV := os.Getenv("USER")
	if ENV == "mhope" {
		err := godotenv.Load()

		if err != nil {
			logger.Fatal("Error loading .env file", err)
		} 
		log.Println("Loaded .env file")			
	}


	db, err := postgres.New(&postgres.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBurl:    os.Getenv("DATABASE_URL"),
	})

	if err != nil {
		logger.Fatal("Failed to Connect to Postgresql database", err)
	}

	err = postgres.SetupDatabase(db,
		&models.Currency{},
		&models.ExchangeRates{},
	)

	if err != nil {
		logger.Fatal("Failed to Setup Tables", err)
	}

	database.RunSeeds(db, []database.SeedFunc{
		database.SeedCurrencies,
	})


	if err != nil {
		logger.Fatal(err)
	}

	if err != nil {
		logger.Fatal(err)
	}

	s := server.New()
	h := handler.New(db)
	
	routes := s.Group("/v1")
	h.Register(routes)

	server.Start(&s, &server.Config{
		Port: fmt.Sprintf(":%s", os.Getenv("PORT")),
	})


}
