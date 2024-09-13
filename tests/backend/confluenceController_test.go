package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGetConfluenceDocs(t *testing.T) {
    req, err := http.NewRequest("GET", "/api/confluence/docs", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetConfluenceDocs)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    // Add more assertions to check response content if needed
}
