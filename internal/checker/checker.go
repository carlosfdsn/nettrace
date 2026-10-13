package checker

import (
	"net/http"
	"time"

	"github.com/cf-neto/nettrace/internal/models"
)

var client = &http.Client{Timeout: 5 * time.Second}

func Check(target string) models.CheckResponse {
	start := time.Now()
	resp, err := client.Get(target)
	latency := time.Since(start)

	result := models.CheckResponse{
		URL: target,
		LatencyMs: latency.Milliseconds(),
		CheckedAt: time.Now().Format(time.RFC3339),
	}

	if err != nil {
		result.Up = false
		result.Error = err.Error()
		return result
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	result.Up = resp.StatusCode >= 200 && resp.StatusCode < 400

	return result
}
