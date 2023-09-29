package handler_test

import (
	"bytes"
	"encurtador_url/src/internal/handler"
	"encurtador_url/src/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateCode(t *testing.T) {
	jsonBody := []byte(`{"url": "http://teste.com"}`)
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/create", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	repo := repository.NewCodeRepositoryInMemory()
	h := handler.NewHandler(repo)

	handler := http.HandlerFunc(h.CreateCode)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}
	assert.Equal(t, http.StatusOK, 200)

}
