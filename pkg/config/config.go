package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
}

var (
	config     *Config
	configOnce sync.Once
	configErr  error
)

func LoadConfig(path string) (*Config, error) {
	configOnce.Do(func() {
		viper.AddConfigPath(path)
		viper.SetConfigType("env")
		viper.SetConfigName("app")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			configErr = fmt.Errorf("failed to read config file: %w", err)
			return
		}

		config = &Config{}
		if err := viper.Unmarshal(config); err != nil {
			configErr = fmt.Errorf("failed to unmarshal config: %w", err)
			config = nil
			return
		}
	})

	if configErr != nil {
		return nil, configErr
	}
	if config == nil {
		return nil, fmt.Errorf("config not initialized")
	}
	return config, nil
}

func GetConfig() (*Config, error) {
	if config == nil {
		return nil, fmt.Errorf("config not loaded yet")
	}
	return config, configErr
}
