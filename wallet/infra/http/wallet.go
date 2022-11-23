package http

import (
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/money-exchange-api/wallet/app/repository"
	"github.com/wellingtonlope/money-exchange-api/wallet/app/usecase"
	"net/http"
)

type Wallet interface {
	Create(c echo.Context) error
	GetByID(c echo.Context) error
}

type wallet struct {
	createUC usecase.Create
	getByID  usecase.GetByID
}

func NewWallet(create usecase.Create, getByID usecase.GetByID) Wallet {
	return &wallet{create, getByID}
}

type WalletResponse struct {
	ID string `json:"id"`
}

func NewWalletResponseFromCreate(output usecase.CreateOutput) WalletResponse {
	return WalletResponse{ID: output.ID}
}

func NewWalletResponseFromGetByID(output usecase.GetByIDOutput) WalletResponse {
	return WalletResponse{ID: output.ID}
}

func (h *wallet) Create(c echo.Context) error {
	result, err := h.createUC.Handle(usecase.CreateInput{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewResponseError(err))
	}

	return c.JSON(http.StatusCreated, NewWalletResponseFromCreate(result))
}

func (h *wallet) GetByID(c echo.Context) error {
	ID := c.Param("id")

	result, err := h.getByID.Handle(ID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == repository.ErrRepositoryWalletNotFound {
			statusCode = http.StatusNotFound
		}
		return c.JSON(statusCode, NewResponseError(err))
	}
	return c.JSON(http.StatusOK, NewWalletResponseFromGetByID(result))
}
