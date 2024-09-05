package config

type LogConfig struct {
	Level        string `kiper_value:"name:level;help:log level;default:info"`
	Format       string `kiper_value:"name:format;help:log format;default:json"`
	File         string `kiper_value:"name:file;help:log file path"`
	LLMChainFile string `kiper_value:"name:llm_chain_file;help:llm chain log file path"`
}

func newLogConfig() *LogConfig {
	return &LogConfig{}
}
