package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Response struct {
	Count int `json:"count"`
}

func TestMainHandler(t *testing.T) {
	router := Router()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	response := Response{}
	json.Unmarshal([]byte(rec.Body.String()), &response)
	assert.Equal(t, response.Count, 42)
}
