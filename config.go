package slacker

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	configEnvPrefix    = "SLACKER"
	configEnvAPIKey    = "API_KEY"
	configEnvChannelID = "CHANNEL_ID"
)

// Config .
type Config struct {
	APIKey    string
	ChannelID string
}

// GetConfig .
func GetConfig() (*Config, error) {
	viper.SetConfigName("slacker")
	viper.AddConfigPath("$HOME/.slacker")
	viper.AddConfigPath(".slacker")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix(configEnvPrefix)
	viper.BindEnv(configEnvAPIKey)
	viper.BindEnv(configEnvChannelID)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	apiKey := viper.GetString(configEnvAPIKey)
	if apiKey == "" {
		return nil, fmt.Errorf("missing $%s_%s", configEnvPrefix, configEnvAPIKey)
	}

	channelID := viper.GetString(configEnvChannelID)
	if channelID == "" {
		return nil, fmt.Errorf("missing $%s_%s", configEnvPrefix, configEnvChannelID)
	}
	return &Config{
		APIKey:    apiKey,
		ChannelID: channelID,
	}, nil
}
