package models

type HealthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

type VersionResponse struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Repository  string `json:"repository"`
	Environment string `json:"environment"`
}

type StatsResponse struct {
	TotalRequests uint64 `json:"total_requests"`
	Uptime        string `json:"uptime"`
	StartedAt     string `json:"started_at"`
}

type CheckResponse struct {
	URL        string `json:"url"`
	Up         bool   `json:"up"`
	StatusCode int    `json:"status_code"`
	LatencyMs  int64  `json:"latency_ms"`
	CheckedAt  string `json:"checked_at"`
	Error      string `json:"error,omitempty"`
}
