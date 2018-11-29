package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ProcessMessage Send The Message
func ProcessMessage(event Messaging) {
	client := &http.Client{}
	response := Response{
		Recipient: User{
			ID: event.Sender.ID,
		},
		Message: Message{
			Attachment: &Attachment{
				Type: "image",
				Payload: Payload{
					URL: IMAGE,
				},
			},
		},
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)

	// url := fmt.Sprintf(FACEBOOK_API, os.Getenv("PAGE_ACCESS_TOKEN"))
	url := fmt.Sprintf(FACEBOOK_API, "EAAEQ9Pux1zsBANVo69TVsLCXE9A5mBsiGDhhfjUObhr8CT5aK8wRbyNeC3m4HJ6o8aq63TVUGojG6hcZBV35IHZBSDdyFlKH9wEe3yf5aASI7uKbT9G39rVDiIFUtHwyloLMoZBW0foMNnKT4pRndEdwqsb86kZAIZBZCank7CDQZDZD")
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
