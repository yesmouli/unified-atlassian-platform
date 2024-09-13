package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGetBitbucketRepos(t *testing.T) {
    req, err := http.NewRequest("GET", "/api/bitbucket/repos", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetBitbucketRepos)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    // Add more assertions to check response content if needed
}
