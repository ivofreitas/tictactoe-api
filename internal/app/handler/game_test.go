package handler

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"tictactoe-api/internal/app/service"
	"tictactoe-api/internal/domain"
	"tictactoe-api/internal/mock"
)

var (
	internalServerErr = errors.New("internal server error")
	badRequestErr     = errors.New("bad request")
)

func TestGet(t *testing.T) {
	testCases := []struct {
		Name           string
		MarshalErr     error
		ExpectedError  error
		ExpectedStatus int
	}{
		{
			Name:           "Test Case 1",
			ExpectedStatus: http.StatusOK,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/game", nil)
			rec := httptest.NewRecorder()

			handler := NewGame(service.NewGame())
			handler.Get(rec, req)

			var result domain.Game
			err := json.Unmarshal(rec.Body.Bytes(), &result)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tc.ExpectedStatus, rec.Code)
			assert.NotNil(t, result)
		})
	}
}

func TestMove(t *testing.T) {
	testCases := []struct {
		Name           string
		Param          *domain.Move
		Result         *domain.Game
		MarshalErr     error
		ExpectedError  error
		ExpectedStatus int
	}{
		{
			Name:           "Test Case 1",
			Param:          &mock.Move,
			ExpectedStatus: http.StatusOK,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			b, _ := json.Marshal(&tc.Param)
			req, _ := http.NewRequest(http.MethodPost, "/game/move", strings.NewReader(string(b)))
			rec := httptest.NewRecorder()

			handler := NewGame(service.NewGame())
			handler.Move(rec, req)

			var result domain.Game
			err := json.Unmarshal(rec.Body.Bytes(), &result)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tc.ExpectedStatus, rec.Code)
			assert.NotNil(t, result)
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		Name           string
		MarshalErr     error
		ExpectedError  error
		ExpectedStatus int
	}{
		{
			Name:           "Test Case 1",
			ExpectedStatus: http.StatusNoContent,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodDelete, "/game", nil)
			rec := httptest.NewRecorder()

			handler := NewGame(service.NewGame())
			handler.Delete(rec, req)

			assert.Equal(t, tc.ExpectedStatus, rec.Code)
			assert.Nil(t, rec.Body.Bytes())
		})
	}
}
