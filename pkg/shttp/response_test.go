package shttp

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseWithDetails(t *testing.T) {
	testCases := []struct {
		name          string
		handler       gin.HandlerFunc
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			handler: func(c *gin.Context) {
				OkResponse(c)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.Contains(t, recorder.Header().Get("Content-Type"), "application/json")
				require.JSONEq(t, `{ "ok": true }`, recorder.Body.String())
			},
		},
		{
			name: "Ok with result",
			handler: func(c *gin.Context) {
				OkResponseWithResult(c, 1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.Contains(t, recorder.Header().Get("Content-Type"), "application/json")
				require.JSONEq(t, `{ "ok": true, "result": 1 }`, recorder.Body.String())
			},
		},
		{
			name: "Ok with result error marshal",
			handler: func(c *gin.Context) {
				OkResponseWithResult(c, json.RawMessage("syntax: error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.Contains(t, recorder.Header().Get("Content-Type"), "application/json")
				require.JSONEq(t, `{ "ok": false, "error_code": 500, "description": "internal server error" }`, recorder.Body.String())
			},
		},
		{
			name: "Error",
			handler: func(c *gin.Context) {
				ErrorResponse(c, http.StatusForbidden, "access denied")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.Contains(t, recorder.Header().Get("Content-Type"), "application/json")
				require.JSONEq(t, `{ "ok": false, "error_code": 403, "description": "access denied" }`, recorder.Body.String())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// settings
			router := gin.New()
			router.GET("/", tc.handler)
			recorder := httptest.NewRecorder()

			// test
			request, err := http.NewRequest(http.MethodGet, "/", nil)
			require.NoError(t, err)

			router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
