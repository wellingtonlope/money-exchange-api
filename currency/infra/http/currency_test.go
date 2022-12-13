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

func TestCurrencyRequest_ToUpdateInput(t *testing.T) {
	request := CurrencyRequest{"BRL", "Real", "Brazil"}
	expected := usecase.UpdateInput{ID: request.ID, Description: request.Description, Country: request.Country}

	got := request.ToUpdateInput()

	assert.Equal(t, expected, got)
}

func TestNewCurrencyResponse(t *testing.T) {
	output := usecase.Output{ID: "BRL", Description: "Real", Country: "Brazil"}
	expected := CurrencyResponse{output.ID, output.Description, output.Country}

	got := NewCurrencyResponse(output)

	assert.Equal(t, expected, got)
}

func TestCurrency_Create(t *testing.T) {
	testCases := []struct {
		name        string
		requestBody string
		ucInput     usecase.CreateInput
		ucOutput    usecase.Output
		ucError     error
		status      int
		response    string
	}{
		{
			name:        "should create a currency",
			requestBody: `{"id":"BRL","description":"Real","country":"Brazil"}`,
			ucInput:     usecase.CreateInput{ID: "BRL", Description: "Real", Country: "Brazil"},
			ucOutput:    usecase.Output{ID: "BRL", Description: "Real", Country: "Brazil"},
			status:      http.StatusCreated,
			response:    `{"id":"BRL","description":"Real","country":"Brazil"}`,
		},
		{
			name:        "should fail when json parse error",
			requestBody: `{`,
			ucInput:     usecase.CreateInput{},
			ucOutput:    usecase.Output{},
			status:      http.StatusBadRequest,
			response:    fmt.Sprintf(`{"message":"%s"}`, ErrJSONParse),
		},
		{
			name:        "should fail with bad request when ID is invalid",
			requestBody: `{"id":"BR","description":"Real","country":"Brazil"}`,
			ucInput:     usecase.CreateInput{ID: "BR", Description: "Real", Country: "Brazil"},
			ucOutput:    usecase.Output{},
			ucError:     domain.ErrCurrencyIDInvalid,
			status:      http.StatusBadRequest,
			response:    fmt.Sprintf(`{"message":"%s"}`, domain.ErrCurrencyIDInvalid.Error()),
		},
		{
			name:        "should fail with conflict when currency already exists",
			requestBody: `{"id":"BRL","description":"Real","country":"Brazil"}`,
			ucInput:     usecase.CreateInput{ID: "BRL", Description: "Real", Country: "Brazil"},
			ucOutput:    usecase.Output{},
			ucError:     repository.ErrRepositoryCurrencyDuplicated,
			status:      http.StatusConflict,
			response:    fmt.Sprintf(`{"message":"%s"}`, repository.ErrRepositoryCurrencyDuplicated.Error()),
		},
		{
			name:        "should fail with internal error when use case fail",
			requestBody: `{"id":"BRL","description":"Real","country":"Brazil"}`,
			ucInput:     usecase.CreateInput{ID: "BRL", Description: "Real", Country: "Brazil"},
			ucOutput:    usecase.Output{},
			ucError:     errors.New("i'm an error"),
			status:      http.StatusInternalServerError,
			response:    `{"message":"i'm an error"}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createUC := new(usecase.CreateMock)
			getByIDUC := new(usecase.GetByIDMock)
			updateUC := new(usecase.UpdateMock)

			ctr := NewCurrency(createUC, getByIDUC, updateUC)

			createUC.
				On("Handle", tc.ucInput).
				Return(tc.ucOutput, tc.ucError)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/currency", strings.NewReader(tc.requestBody))
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

func TestCurrency_GetByID(t *testing.T) {
	testCases := []struct {
		name             string
		output           usecase.Output
		status           int
		error            error
		expectedResponse string
	}{
		{
			name:             "should get currency by ID",
			output:           usecase.Output{ID: "BRL", Description: "Real", Country: "Brazil"},
			status:           http.StatusOK,
			expectedResponse: `{"id":"BRL","description":"Real","country":"Brazil"}`,
		},
		{
			name:             "should fail with not found when currency not exists",
			output:           usecase.Output{ID: "BRL"},
			status:           http.StatusNotFound,
			error:            repository.ErrRepositoryCurrencyNotFound,
			expectedResponse: `{"message":"currency not found"}`,
		},
		{
			name:             "should fail with internal error when use case fail",
			output:           usecase.Output{ID: "BRL"},
			status:           http.StatusInternalServerError,
			error:            errors.New("i'm an error"),
			expectedResponse: `{"message":"i'm an error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createUC := new(usecase.CreateMock)
			getByIDUC := new(usecase.GetByIDMock)
			updateUC := new(usecase.UpdateMock)

			ctr := NewCurrency(createUC, getByIDUC, updateUC)

			getByIDUC.
				On("Handle", tc.output.ID).
				Return(tc.output, tc.error)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/currency/%s", tc.output.ID), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/currency/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.output.ID)

			if assert.NoError(t, ctr.GetByID(c)) {
				assert.Equal(t, tc.status, rec.Code)
				assert.Equal(t, tc.expectedResponse, strings.ReplaceAll(rec.Body.String(), "\n", ""))
			}
		})
	}
}

func TestCurrency_Update(t *testing.T) {
	testCases := []struct {
		name        string
		requestBody string
		ucOutput    usecase.Output
		ucError     error
		status      int
		response    string
	}{
		{
			name:        "should update a currency",
			requestBody: `{"id":"BRL","description":"Real","country":"Brazil"}`,
			ucOutput:    usecase.Output{ID: "BRL", Description: "Real", Country: "Brazil"},
			status:      http.StatusOK,
			response:    `{"id":"BRL","description":"Real","country":"Brazil"}`,
		},
		{
			name:        "should fail when json parse error",
			requestBody: `{`,
			ucOutput:    usecase.Output{ID: "BRL", Description: "Real", Country: "Brazil"},
			status:      http.StatusBadRequest,
			response:    fmt.Sprintf(`{"message":"%s"}`, ErrJSONParse),
		},
		{
			name:        "should fail with not found when currency not exists",
			requestBody: `{"id":"BRL","description":"Real","country":"Brazil"}`,
			ucOutput:    usecase.Output{ID: "BRL", Description: "Real", Country: "Brazil"},
			ucError:     repository.ErrRepositoryCurrencyNotFound,
			status:      http.StatusNotFound,
			response:    fmt.Sprintf(`{"message":"%s"}`, repository.ErrRepositoryCurrencyNotFound.Error()),
		},
		{
			name:        "should fail with internal error when use case fail",
			requestBody: `{"id":"BRL","description":"Real","country":"Brazil"}`,
			ucOutput:    usecase.Output{ID: "BRL", Description: "Real", Country: "Brazil"},
			ucError:     errors.New("i'm an error"),
			status:      http.StatusInternalServerError,
			response:    `{"message":"i'm an error"}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createUC := new(usecase.CreateMock)
			getByIDUC := new(usecase.GetByIDMock)
			updateUC := new(usecase.UpdateMock)

			ctr := NewCurrency(createUC, getByIDUC, updateUC)

			updateUC.
				On("Handle", usecase.UpdateInput{ID: tc.ucOutput.ID, Description: tc.ucOutput.Description, Country: tc.ucOutput.Country}).
				Return(tc.ucOutput, tc.ucError)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/currency", strings.NewReader(tc.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, ctr.Update(c)) {
				assert.Equal(t, tc.status, rec.Code)
				assert.Equal(t, tc.response, strings.ReplaceAll(rec.Body.String(), "\n", ""))
			}
		})
	}
}
