package config

type MoonshotConfig struct {
	APIKey string `kiper_value:"name:api_key;help:api key;default:"`
}

func newMoonshotConfig() *MoonshotConfig {
	return &MoonshotConfig{}
}
