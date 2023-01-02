package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/money-exchange-api/wallet/app/usecase"
	"github.com/wellingtonlope/money-exchange-api/wallet/infra/http"
	"github.com/wellingtonlope/money-exchange-api/wallet/infra/memory"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil && os.Getenv("APP_ENV") != "production" {
		log.Fatalf("Error loading .env file: %v", err)
	}

	walletRepository := memory.NewWallet()
	createUC := usecase.NewCreate(walletRepository)
	getByIDUC := usecase.NewGetByID(walletRepository)
	walletHandle := http.NewWallet(createUC, getByIDUC)

	e := echo.New()

	e.POST("/wallet", walletHandle.Create)
	e.GET("/wallet/:id", walletHandle.GetByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
