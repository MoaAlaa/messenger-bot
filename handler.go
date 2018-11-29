package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
)

// VerificationEndpoint Use To Verify The Facebook Token
func VerificationEndpoint(c echo.Context) error {
	challenge := c.QueryParam("hub.challenge")
	mode := c.QueryParam("hub.mode")
	token := c.QueryParam("hub.verify_token")
	var err error

	// if mode != "" && token == os.Getenv("VERIFY_TOKEN") {
	if mode != "" && token == "Alaa Chat Bot Testing" {
		err = c.String(http.StatusOK, challenge)
	} else {
		err = c.String(http.StatusNotFound, "Error, wrong validation token")
	}

	return err
}

// MessagesEndpoint Verify If The Message And The Event Is Ok To Be Send
func MessagesEndpoint(c echo.Context) error {
	fmt.Printf("%+v", c.Request().Body)
	var callback Callback
	var err error

	// json.Unmarshal([]byte(c.Request().Body), &callback)

	json.NewDecoder(c.Request().Body).Decode(&callback)

	fmt.Printf("%+v", callback)
	if callback.Object == "page" {
		for _, entry := range callback.Entry {
			for _, event := range entry.Messaging {
				if !reflect.DeepEqual(event.Message, Message{}) && event.Message.Text != "" {
					ProcessMessage(event)
				}
			}
		}
		println(" Page")
		err = c.String(http.StatusOK, "Got Your Message")
	} else {
		println("Not Page")
		err = c.String(http.StatusNotFound, "Message Not Supported")
	}

	return err
}
