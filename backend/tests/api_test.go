package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"unified-atlassian-platform/api/jira"
	"unified-atlassian-platform/api/confluence"
	"unified-atlassian-platform/api/bitbucket"

	"github.com/stretchr/testify/assert"
)

func TestJiraHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/jira", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(jira.HandleRequests)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestConfluenceHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/confluence", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(confluence.HandleRequests)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestBitbucketHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/bitbucket", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(bitbucket.HandleRequests)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
