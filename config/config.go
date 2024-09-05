package config

import (
	"path/filepath"

	"github.com/leeif/kiper"
)

type Config struct {
	Version   string
	Moonshot  *MoonshotConfig  `kiper_config:"name:moonshot"`
	APIServer *APIServerConfig `kiper_config:"name:api_server"`
	Log       *LogConfig       `kiper_config:"name:log"`
}

func NewConfig(args []string, version string) (*Config, error) {
	c := &Config{
		APIServer: newAPIServerConfig(),
		Moonshot:  newMoonshotConfig(),
		Log:       newLogConfig(),
	}
	if len(args) == 0 {
		args = append(args, "")
	}
	kiper := kiper.NewKiper(filepath.Base(args[0]), "")
	kiper.Kingpin.Version(version)
	kiper.Kingpin.HelpFlag.Short('h')

	kiper.SetConfigFileFlag("config.file", "config file", "./config.json")

	if err := kiper.Parse(c, args[1:]); err != nil {
		return nil, err
	}
	c.Version = version
	return c, nil
}
