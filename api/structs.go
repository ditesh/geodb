package api

import "geodb/config"

type wrapper struct{}

type Server struct {
	Config config.APIServerConfig
}
