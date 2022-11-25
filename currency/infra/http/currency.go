package http

import (
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/money-exchange-api/currency/app/repository"
	"github.com/wellingtonlope/money-exchange-api/currency/app/usecase"
	"github.com/wellingtonlope/money-exchange-api/currency/domain"
	"net/http"
)

type CurrencyRequest struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Country     string `json:"country"`
}

func (r *CurrencyRequest) ToCreateInput() usecase.CreateInput {
	return usecase.CreateInput{ID: r.ID, Description: r.Description, Country: r.Country}
}

type CurrencyResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Country     string `json:"country"`
}

func NewCurrencyResponse(currency usecase.Output) CurrencyResponse {
	return CurrencyResponse{currency.ID, currency.Description, currency.Country}
}

type Currency interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	GetByID(c echo.Context) error
}

type currency struct {
	createUC  usecase.Create
	getByIDUC usecase.GetByID
	updateUC  usecase.Update
}

func NewCurrency(createUC usecase.Create, getByIDUC usecase.GetByID, updateUC usecase.Update) Currency {
	return &currency{createUC, getByIDUC, updateUC}
}

func (ctr *currency) Create(c echo.Context) error {
	var request CurrencyRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewResponseError(err))
	}

	output, err := ctr.createUC.Handle(request.ToCreateInput())
	if err != nil {
		status := http.StatusInternalServerError
		switch err {
		case domain.ErrCurrencyIDInvalid:
			status = http.StatusBadRequest
		case repository.ErrRepositoryCurrencyDuplicated:
			status = http.StatusConflict
		}
		return c.JSON(status, NewResponseError(err))
	}

	return c.JSON(http.StatusCreated, NewCurrencyResponse(output))
}

func (ctr *currency) Update(c echo.Context) error {
	return nil
}

func (ctr *currency) GetByID(c echo.Context) error {
	return nil
}
