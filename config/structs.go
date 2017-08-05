package config

type Config struct {
	Datapath  string          `json:"datapath"`
	APIServer APIServerConfig `json:"server"`
	Logger    LoggerConfig    `json:"logger"`
}

type APIServerConfig struct {
	Port int `json:"port"`
}

type LoggerConfig struct {
	Type  string `json:"type"`
	Path  string `json:"path"`
	Level string `json:"level"`
}
