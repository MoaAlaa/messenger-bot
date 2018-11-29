package main

import (
	"github.com/labstack/echo"
)

const (
	FACEBOOK_API = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"
	IMAGE        = "http://37.media.tumblr.com/e705e901302b5925ffb2bcf3cacb5bcd/tumblr_n6vxziSQD11slv6upo3_500.gif"
)

func main() {
	e := echo.New()

	e.POST("/webhook", MessagesEndpoint)
	e.GET("/webhook", VerificationEndpoint)

	e.Logger.Fatal(e.Start(":9999"))
}
