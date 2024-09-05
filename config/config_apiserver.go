package config

type APIServerConfig struct {
	Port *Port `kiper_value:"name:port;help:server port;default:8080"`
}

func newAPIServerConfig() *APIServerConfig {
	return &APIServerConfig{
		Port: &Port{},
	}
}
