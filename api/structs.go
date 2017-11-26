package api

import "geodb/config"

type wrapper struct{}

// Server is a holder struct for the API server configuration
type Server struct {
	Config config.APIServerConfig
}
