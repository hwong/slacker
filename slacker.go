package slacker

import (
	"github.com/nlopes/slack"
)

// Post .
func Post(config *Config, message string) error {
	api := slack.New(config.APIKey)
	_, _, err := api.PostMessage(config.ChannelID, message, slack.PostMessageParameters{})
	return err
}
