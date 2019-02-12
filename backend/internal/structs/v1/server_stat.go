package v1

// ServerStat returns statistics of server
type ServerStat struct {
	StartTime string  `json:"start_time"`
	Events    uint64  `json:"events"`
	Config    *Config `json:"config"`
}
