package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
	"unicode"

	"github.com/cmumford/go-starter.git/api"
	"github.com/stretchr/testify/assert"
)

func removeTrailingWhitespaceBytes(b []byte) []byte {
	if len(b) == 0 {
		return b
	}
	i := len(b) - 1
	for i >= 0 && unicode.IsSpace(rune(b[i])) {
		i--
	}
	return b[:i+1]
}

func isMinimizedJSON(jsonStr []byte) bool {
	input := removeTrailingWhitespaceBytes(jsonStr)
	var temp interface{}
	if err := json.Unmarshal([]byte(input), &temp); err != nil {
		return false
	}
	unmarshalledStr, err := json.Marshal(temp)
	if err != nil {
		return false
	}
	return len(unmarshalledStr) == len(input)
}

func TestHealthHandler_ReturnsOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err, "Failed to create request")
	rr := httptest.NewRecorder()
	api.HealthHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.Equal(t, "OK", rr.Body.String())
}

func TestRootHandler_ReturnsStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err, "Failed to create request")
	rr := httptest.NewRecorder()
	api.RootHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
}

func TestRootHandler_ContentType_IsJSON(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	api.RootHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content-Type should be application/json")
}

func TestRootHandler_TestMessage_ExpectedPrefix(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	api.RootHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var resp api.Response
	err = json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err, "Failed to decode JSON")
	assert.True(t, strings.HasPrefix(resp.Message, "My name is"))
}

func TestRootHandler_TestTimestamp_ExpectRecentTime(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	api.RootHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var resp api.Response
	err = json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err, "Failed to decode JSON")
	now := time.Now().UnixMilli()
	assert.LessOrEqual(t, resp.Timestamp, now)
	assert.GreaterOrEqual(t, resp.Timestamp, now-5000)
}

func TestRootHandler_TestCommitID_UsesEnvVar(t *testing.T) {
	os.Setenv("GIT_COMMIT_ID", "1234567890abcdef")

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	api.RootHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var resp api.Response
	err = json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err, "Failed to decode JSON")
	assert.Equal(t, resp.CommitID, "1234567890abcdef")
}

func TestRootHandler_AnalyzeJSON_IsMinimized(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	api.RootHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.True(t, isMinimizedJSON(rr.Body.Bytes()), "Message should be minimized")
}
