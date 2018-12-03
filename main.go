package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/paked/messenger"
)

func main() {
	// Create a new messenger client
	client := messenger.New(messenger.Options{
		Verify:      *verify,
		AppSecret:   *appSecret,
		VerifyToken: *verifyToken,
		Token:       *pageToken,
	})

	// err := client.EnableChatExtension(messenger.HomeURL{
	// 	URL:                *serverURL,
	// 	WebviewHeightRatio: "tall",
	// 	WebviewShareButton: "show",
	// 	InTest:             true,
	// })

	// if err != nil {
	// 	fmt.Println("Failed to EnableChatExtension, err=", err)
	// }

	// Setup a handler to be triggered when a message is received
	client.HandleMessage(func(m messenger.Message, r *messenger.Response) {
		fmt.Printf("%v (Sent, %v)\n", m.Text, m.Time.Format(time.UnixDate))

		p, err := client.ProfileByID(m.Sender.ID)
		if err != nil {
			fmt.Println("Something went wrong!", err)
		}

		mes := strings.ToLower(m.Text)

		var con []Country

		// res, err := http.Get(fmt.Sprintf("https://restcountries.eu/rest/v2/name/%s?fullText=true", mes))
		res, err := http.Get(fmt.Sprintf("https://restcountries.eu/rest/v2/name/%s", mes))

		if err != nil {
			fmt.Println(err.Error())
		}

		defer res.Body.Close()

		body, errr := ioutil.ReadAll(res.Body)

		if errr != nil {
			fmt.Println(errr.Error())
		}

		json.Unmarshal(body, &con)

		if len(con) == 0 {
			fmt.Println("Error")
			help(p, r)
		} else {
			var countryData = []messenger.StructuredMessageButton{}

			for _, c := range con {

				u := fmt.Sprintf("http://country.io/%s/", strings.ToLower(strings.Replace(c.NativeName, " ", "-", -1)))
				countryData = append(countryData, messenger.StructuredMessageButton{
					Type:  "web_url",
					URL:   u,
					Title: c.NativeName,
				})
			}
			countryButton(r, countryData)
		}
	})

	// Setup a handler to be triggered when a message is delivered
	client.HandleDelivery(func(d messenger.Delivery, r *messenger.Response) {
		fmt.Println("Delivered at:", d.Watermark().Format(time.UnixDate))
	})

	// Setup a handler to be triggered when a message is read
	client.HandleRead(func(m messenger.Read, r *messenger.Response) {
		fmt.Println("Read at:", m.Watermark().Format(time.UnixDate))
	})

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("Serving messenger bot on", addr)

	log.Fatal(http.ListenAndServe(addr, client.Handler()))
}

// countryButtons will present to the user a button that can be used to
func countryButton(r *messenger.Response, countryData []messenger.StructuredMessageButton) error {
	fmt.Println("Button Send")
	buttons := &countryData
	return r.ButtonTemplate("Go And See :)", buttons, messenger.ResponseType)
}

// help displays possibles actions to the user.
func help(p messenger.Profile, r *messenger.Response) error {
	text := fmt.Sprintf(
		"%s, Error Not Data Found. Here is what I understand Fucker, To Can I Search By.",
		p.FirstName,
	)

	replies := []messenger.QuickReply{
		{
			ContentType: "text",
			Title:       "Egypt",
		},
		{
			ContentType: "text",
			Title:       "United",
		},
	}

	return r.TextWithReplies(text, replies, messenger.ResponseType)
}
