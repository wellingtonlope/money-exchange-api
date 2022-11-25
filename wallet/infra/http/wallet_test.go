package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/money-exchange-api/wallet/app/repository"
	"github.com/wellingtonlope/money-exchange-api/wallet/app/usecase"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewWalletResponseFromCreate(t *testing.T) {
	t.Run("should create a create response", func(t *testing.T) {
		output := usecase.CreateOutput{ID: "123"}
		expectedWalletResponse := WalletResponse{ID: "123"}

		got := NewWalletResponseFromCreate(output)

		assert.Equal(t, expectedWalletResponse, got)
	})
}

func TestNewWalletResponseFromGetByID(t *testing.T) {
	t.Run("should create a create response", func(t *testing.T) {
		output := usecase.GetByIDOutput{ID: "123"}
		expectedWalletResponse := WalletResponse{ID: "123"}

		got := NewWalletResponseFromGetByID(output)

		assert.Equal(t, expectedWalletResponse, got)
	})
}

func TestWallet_Create(t *testing.T) {
	t.Run("should create an wallet", func(t *testing.T) {
		createUC := new(usecase.CreateMock)
		getByIDUC := new(usecase.GetByIDMock)
		ctr := NewWallet(createUC, getByIDUC)

		createUC.
			On("Handle", usecase.CreateInput{}).
			Return(usecase.CreateOutput{ID: "1"}, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/wallet", strings.NewReader("{}"))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, ctr.Create(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, `{"id":"1"}`, strings.ReplaceAll(rec.Body.String(), "\n", ""))
		}
	})

	t.Run("should fail when use case fail to create an wallet", func(t *testing.T) {
		createUC := new(usecase.CreateMock)
		getByIDUC := new(usecase.GetByIDMock)
		ctr := NewWallet(createUC, getByIDUC)

		createUC.
			On("Handle", usecase.CreateInput{}).
			Return(usecase.CreateOutput{}, errors.New("i'm an error"))

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, ctr.Create(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "{\"message\":\"i'm an error\"}\n", rec.Body.String())
		}
	})
}

func TestWallet_GetByID(t *testing.T) {
	t.Run("should get wallet by ID", func(t *testing.T) {
		createUC := new(usecase.CreateMock)
		getByIDUC := new(usecase.GetByIDMock)
		ctr := NewWallet(createUC, getByIDUC)

		getByIDUC.
			On("Handle", "1").
			Return(usecase.GetByIDOutput{ID: "1"}, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("{}"))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/wallet/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, ctr.GetByID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "{\"id\":\"1\"}\n", rec.Body.String())
		}
	})

	t.Run("should return not found when wallet not exists", func(t *testing.T) {
		createUC := new(usecase.CreateMock)
		getByIDUC := new(usecase.GetByIDMock)
		ctr := NewWallet(createUC, getByIDUC)

		getByIDUC.
			On("Handle", "1").
			Return(usecase.GetByIDOutput{}, repository.ErrRepositoryWalletNotFound)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("{}"))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/wallet/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, ctr.GetByID(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "{\"message\":\"wallet not found\"}\n", rec.Body.String())
		}
	})

	t.Run("should fail when use case fail", func(t *testing.T) {
		createUC := new(usecase.CreateMock)
		getByIDUC := new(usecase.GetByIDMock)
		ctr := NewWallet(createUC, getByIDUC)

		getByIDUC.
			On("Handle", "1").
			Return(usecase.GetByIDOutput{}, errors.New("i'm an error"))

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("{}"))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/wallet/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, ctr.GetByID(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "{\"message\":\"i'm an error\"}\n", rec.Body.String())
		}
	})
}
