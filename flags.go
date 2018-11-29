package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	verifyToken = flag.String("verify-token", "alaa-testing", "The token used to verify facebook (required)")
	verify      = flag.Bool("should-verify", true, "Whether or not the app should verify itself")
	pageToken   = flag.String("page-token", "EAAEQ9Pux1zsBAGUOXueXQ97Dt5x04C7AGD0p38VjAI5BZA4ZCfRLoQDhdgUJm1Ey9yz5GClGPFXWrdajPOGC4yQK6HMERjBqy4fgZAoDjtt9z1k1o20gagusErYus3ToczerpZAUFBSosN9ByLLJoWUGeA2zXikvmUtMC4mhiQZDZD", "The token that is used to verify the page on facebook (required)")
	appSecret   = flag.String("app-secret", "46a9533b73569fb273fd6d8bcd6a7c55", "The app secret from the facebook developer portal (required)")
	host        = flag.String("host", "localhost", "The host used to serve the messenger bot")
	port        = flag.Int("port", 9999, "The port used to serve the messenger bot")
)

func init() {
	flag.Parse()
	println(*verifyToken)
	println(*appSecret)
	println(*pageToken)
	if *verifyToken == "" || *appSecret == "" || *pageToken == "" {
		fmt.Println("missing arguments")
		fmt.Println()
		flag.Usage()

		os.Exit(-1)
	}
}
