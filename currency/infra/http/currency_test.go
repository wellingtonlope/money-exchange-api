package http

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/money-exchange-api/currency/app/repository"
	"github.com/wellingtonlope/money-exchange-api/currency/app/usecase"
	"github.com/wellingtonlope/money-exchange-api/currency/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCurrencyRequest_ToCreateInput(t *testing.T) {
	request := CurrencyRequest{"BRL", "Real", "Brazil"}
	expected := usecase.CreateInput{ID: request.ID, Description: request.Description, Country: request.Country}

	got := request.ToCreateInput()

	assert.Equal(t, expected, got)
}

func TestNewCurrencyResponse(t *testing.T) {
	t.Run("should create a currency response", func(t *testing.T) {
		output := usecase.Output{ID: "BRL", Description: "Real", Country: "Brazil"}
		expected := CurrencyResponse{output.ID, output.Description, output.Country}

		got := NewCurrencyResponse(output)

		assert.Equal(t, expected, got)
	})
}

func TestCurrency_Create(t *testing.T) {
	testCases := []struct {
		name     string
		currency domain.Currency
		error    error
		status   int
		response string
	}{
		{
			name:     "should create a currency response",
			currency: domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"},
			status:   http.StatusCreated,
			response: `{"id":"BRL","description":"Real","country":"Brazil"}`,
		},
		{
			name:     "should fail with bad request when ID is invalid",
			currency: domain.Currency{ID: "BR", Description: "Real", Country: "Brazil"},
			error:    domain.ErrCurrencyIDInvalid,
			status:   http.StatusBadRequest,
			response: fmt.Sprintf(`{"message":"%s"}`, domain.ErrCurrencyIDInvalid.Error()),
		},
		{
			name:     "should fail with conflict when currency already exists",
			currency: domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"},
			error:    repository.ErrRepositoryCurrencyDuplicated,
			status:   http.StatusConflict,
			response: fmt.Sprintf(`{"message":"%s"}`, repository.ErrRepositoryCurrencyDuplicated.Error()),
		},
		{
			name:     "should fail with internal error when use case fail",
			currency: domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"},
			error:    errors.New("i'm an error"),
			status:   http.StatusInternalServerError,
			response: `{"message":"i'm an error"}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createUC := new(usecase.CreateMock)
			getByIDUC := new(usecase.GetByIDMock)
			updateUC := new(usecase.UpdateMock)
			expectedJSON := fmt.Sprintf(`{"id":"%s","description":"%s","country":"%s"}`, tc.currency.ID, tc.currency.Description, tc.currency.Country)

			ctr := NewCurrency(createUC, getByIDUC, updateUC)
			output := usecase.Output{}
			if tc.error == nil {
				output = usecase.NewOutput(tc.currency)
			}

			createUC.
				On("Handle", usecase.CreateInput{ID: tc.currency.ID, Description: tc.currency.Description, Country: tc.currency.Country}).
				Return(output, tc.error)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/currency", strings.NewReader(expectedJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, ctr.Create(c)) {
				assert.Equal(t, tc.status, rec.Code)
				assert.Equal(t, tc.response, strings.ReplaceAll(rec.Body.String(), "\n", ""))
			}
		})
	}
}
