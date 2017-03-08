package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"io/ioutil"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("slacker")
	viper.AddConfigPath("$HOME/.slacker")
	viper.AddConfigPath(".slacker")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("SLACKER")
	viper.BindEnv("API_KEY")
	viper.BindEnv("CHANNEL_ID")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config: %v\n.", err)
		return
	}

	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		fmt.Println("Missing $SLACKER_API_KEY.")
		return
	}

	channelID := viper.GetString("CHANNEL_ID")
	if channelID == "" {
		fmt.Println("Missing $SLACKER_CHANNEL_ID.")
		return
	}

	message := strings.Join(os.Args[1:], " ")
	if len(message) == 0 {
		reader := bufio.NewReader(os.Stdin)
		read, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Printf("Error reading from stdin: %v\n.", err)
			return
		}
		message = string(read)
	}

	api := slack.New(apiKey)
	_, _, err = api.PostMessage(channelID, message, slack.PostMessageParameters{})
	if err != nil {
		fmt.Printf("Error posting message to Slack: %v\n", err)
		return
	}

	fmt.Println(message)
}
