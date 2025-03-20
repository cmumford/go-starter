package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	CommitID  string `json:"git-commit-id"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message:   "My name is timestamper",
		Timestamp: time.Now().UnixMilli(),
		CommitID:  os.Getenv("GIT_COMMIT_ID"),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
