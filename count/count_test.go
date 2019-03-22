package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountRoute(t *testing.T) {
	//router := setupRouter()
	router := main()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/count", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	//assert.Equal(t, "pong", w.Body.String())
}