package config

// Config is mapping struct for the config file
type Config struct {
	Datapath  string          `json:"datapath"`
	APIServer APIServerConfig `json:"server"`
	Logger    LoggerConfig    `json:"logger"`
}

// APIServerConfig is an API server configuration struct
type APIServerConfig struct {
	Port int `json:"port"`
}

// LoggerConfig is the logger configuration struct
type LoggerConfig struct {
	Type  string `json:"type"`
	Path  string `json:"path"`
	Level string `json:"level"`
}
