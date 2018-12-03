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

	// HandleMessage Should Be The Last Function

	client.GreetingSetting("Hello Fucker")

	// Setup a handler to be triggered when a message is delivered
	client.HandleDelivery(func(d messenger.Delivery, r *messenger.Response) {
		fmt.Println("Delivered at:", d.Watermark().Format(time.UnixDate))
	})

	// Setup a handler to be triggered when a message is read
	client.HandleRead(func(m messenger.Read, r *messenger.Response) {
		fmt.Println("Read at:", m.Watermark().Format(time.UnixDate))
	})

	// Setup a handler to be triggered when a message is received
	client.HandleMessage(func(m messenger.Message, r *messenger.Response) {
		fmt.Printf("%v (Sent, %v)\n", m.Text, m.Time.Format(time.UnixDate))

		p, err := client.ProfileByID(m.Sender.ID)
		if err != nil {
			fmt.Println("Something went wrong!", err)
		}

		mes := strings.ToLower(m.Text)

		if mes == "help" {
			fmt.Println("Help")
			help(p, r, "We Take The Name Of The Country And Then Search For It's Information Just Type The Country Name Or Part Of It.")
		} else {
			var con []Country

			getCountries(mes, &con)

			fmt.Println(len(con))

			if len(con) == 0 {

				fmt.Println("Not Data Found")
				help(p, r, "Error Not Data Found. Here is what I understand Fucker, To Can I Search By.")

			} else {
				var countryData = []messenger.StructuredMessageElement{}

				getCountryMessageData(con, &countryData)

				countryResponseTemplate(r, countryData)
			}

		}
	})

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("Serving messenger bot on", addr)

	log.Fatal(http.ListenAndServe(addr, client.Handler()))
}

// getCountries Get The Countries From The Api By The Specified Name Or Part Of The Name
func getCountries(mes string, con *[]Country) {

	res, err := http.Get(fmt.Sprintf("https://restcountries.eu/rest/v2/name/%s?fields=name;capital;flag;region;nativeName", mes))

	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	body, errr := ioutil.ReadAll(res.Body)

	if errr != nil {
		fmt.Println(errr.Error())
	}

	json.Unmarshal(body, &con)
}

// getCountryMessageData Send The Countries Names And Data With Buttons
func getCountryMessageData(con []Country, cd *[]messenger.StructuredMessageElement) {
	for _, c := range con {

		u := fmt.Sprintf("http://google.com/%s/", strings.ToLower(strings.Replace(c.Name, " ", "-", -1)))
		*cd = append(*cd, messenger.StructuredMessageElement{
			Title:    c.NativeName,
			ImageURL: c.Flag,
			ItemURL:  u,
			Subtitle: c.Name,
		})
	}
}

// countryButtons will present to the user a button that can be used to
func countryButton(r *messenger.Response, countryData []messenger.StructuredMessageButton) error {
	fmt.Println("Button Sent")
	buttons := &countryData
	return r.ButtonTemplate("Go And See :)", buttons, messenger.ResponseType)
}

// countryResponseTemplate will present to the user a Template / button that can be used to
func countryResponseTemplate(r *messenger.Response, countryData []messenger.StructuredMessageElement) error {
	fmt.Println("Template Sent")

	fmt.Printf("%+v", countryData)

	t := &countryData

	return r.GenericTemplate(t, messenger.ResponseType)
}

// help displays possibles actions to the user.
func help(p messenger.Profile, r *messenger.Response, m string) error {
	text := fmt.Sprintf(
		"%s, %s",
		p.FirstName,
		m,
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
