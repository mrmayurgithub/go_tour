package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3107445446615-3119193364741-SVPlTpFaHQ9af4uY6MtjupNk")
	os.Setenv("CHANNEL_ID", "C033L1P5H27")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"The Psychology of Money Timeless Lessons on Wealth, Greed, and Happiness by Morgan Housel.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			panic(err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
