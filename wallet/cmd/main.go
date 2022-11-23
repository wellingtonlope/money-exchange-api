package main

import (
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/money-exchange-api/wallet/app/usecase"
	"github.com/wellingtonlope/money-exchange-api/wallet/infra/http"
	"github.com/wellingtonlope/money-exchange-api/wallet/infra/memory"
)

func main() {
	walletRepository := memory.NewWallet()
	createUC := usecase.NewCreate(walletRepository)
	getByIDUC := usecase.NewGetByID(walletRepository)
	walletHandle := http.NewWallet(createUC, getByIDUC)

	e := echo.New()

	e.POST("/wallet", walletHandle.Create)
	e.GET("/wallet/:id", walletHandle.GetByID)

	e.Logger.Fatal(e.Start(":1300"))
}
