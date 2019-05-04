package config

import "github.com/spf13/viper"

const (
	AppPort = "APP_PORT"
)

type Config struct {
	cfg  *viper.Viper
	Port int
}

func New(cfg *viper.Viper) (*Config, error) {
	config := &Config{cfg: cfg}
	return config, nil
}

func (config *Config) GetAppPort() int {
	return config.cfg.GetInt(AppPort)
}

