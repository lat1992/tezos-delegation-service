package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lat1992/tezos-delegation-service/mocks"
	"github.com/lat1992/tezos-delegation-service/structs"
	"github.com/stretchr/testify/assert"
)

func TestDelegations(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		mockReturn     []structs.Data
		mockError      error
		expectedStatus int
		expectedBody   map[string][]structs.Data
	}{
		{
			name: "success",
			mockReturn: []structs.Data{
				{Amount: "100", Timestamp: "2023-01-01"},
				{Amount: "200", Timestamp: "2023-02-01"},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: map[string][]structs.Data{
				"data": {
					{Amount: "100", Timestamp: "2023-01-01"},
					{Amount: "200", Timestamp: "2023-02-01"},
				},
			},
		},
		{
			name:           "Error - Service failure",
			mockReturn:     nil,
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mocks.MockTezosDelegationService)
			mockService.On("GetDelegations", "").Return(tt.mockReturn, tt.mockError)

			router := gin.New()
			router.GET("/delegations", delegations(mockService))

			req, _ := http.NewRequest("GET", "/delegations", nil)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedBody != nil {
				var response map[string][]structs.Data
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestDelegationsByYear(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		year           string
		mockReturn     []structs.Data
		mockError      error
		expectedStatus int
		expectedBody   map[string][]structs.Data
	}{
		{
			name: "success",
			year: "2023",
			mockReturn: []structs.Data{
				{Amount: "100", Timestamp: "2023-01-01"},
				{Amount: "200", Timestamp: "2023-02-01"},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: map[string][]structs.Data{
				"data": {
					{Amount: "100", Timestamp: "2023-01-01"},
					{Amount: "200", Timestamp: "2023-02-01"},
				},
			},
		},
		{
			name:           "Error - Service failure",
			year:           "2023",
			mockReturn:     nil,
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mocks.MockTezosDelegationService)
			mockService.On("GetDelegations", tt.year).Return(tt.mockReturn, tt.mockError).Once()

			router := gin.New()
			router.GET("/delegations/:year", delegationsByYear(mockService))

			req, _ := http.NewRequest("GET", "/delegations/"+tt.year, nil)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedBody != nil {
				var response map[string][]structs.Data
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			}

			mockService.AssertExpectations(t)
		})
	}
}
