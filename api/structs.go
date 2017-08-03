package api

type wrapper struct{}

type Server struct {
	ConfigFile string
	config     Config
}

type Config struct {
	Port int
}
