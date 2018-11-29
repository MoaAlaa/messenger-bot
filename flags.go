package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	verifyToken = flag.String("verify-token", "alaa-testing", "The token used to verify facebook (required)")
	verify      = flag.Bool("should-verify", false, "Whether or not the app should verify itself")
	pageToken   = flag.String("page-token", "not alaa", "The token that is used to verify the page on facebook")
	appSecret   = flag.String("app-secret", "", "The app secret from the facebook developer portal (required)")
	host        = flag.String("host", "localhost", "The host used to serve the messenger bot")
	port        = flag.Int("port", 9999, "The port used to serve the messenger bot")
)

func init() {
	flag.Parse()

	if *verifyToken == "" || *appSecret == "" || *pageToken == "" {
		fmt.Println("missing arguments")
		fmt.Println()
		flag.Usage()

		os.Exit(-1)
	}
}
