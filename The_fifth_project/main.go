package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

// import(
// 	"fmt"
// 	"os"
// 	"github.com/slack-go/slack"
// )

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-318246501329-2246888119766-C2TQU9hIq6qMtSvloSXVSKQ7")
	os.Setenv("CHANNEL_ID", "c9478e238")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"ZIPL.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{

			// slack.FileUploadParameters is the function that slack gives us and it has two parameters :
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
