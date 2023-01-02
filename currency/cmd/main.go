package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/money-exchange-api/currency/app/usecase"
	"github.com/wellingtonlope/money-exchange-api/currency/infra/http"
	"github.com/wellingtonlope/money-exchange-api/currency/infra/memory"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil && os.Getenv("APP_ENV") != "production" {
		log.Fatalf("Error loading .env file: %v", err)
	}

	currencyRepository := memory.NewCurrency()
	createUC := usecase.NewCreate(currencyRepository)
	getByIDUC := usecase.NewGetByID(currencyRepository)
	updateUC := usecase.NewUpdate(currencyRepository)
	currencyHandle := http.NewCurrency(createUC, getByIDUC, updateUC)

	e := echo.New()

	e.POST("/currency", currencyHandle.Create)
	e.GET("/currency/:id", currencyHandle.GetByID)
	e.PUT("/currency/:id", currencyHandle.Update)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
