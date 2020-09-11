package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
	"time"
)

func SendEmail(feed *gofeed.Feed) {
	items := feed.Items
	name := feed.Title
	m := mail.NewV3Mail()

	e := mail.NewEmail("Veckans nyheter", os.Getenv("EMAIL_FROM"))
	m.SetFrom(e)

	m.SetTemplateID("d-08b780f303f94dd18a80366c968b377b")

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("", os.Getenv("EMAIL_TO")),
	}
	p.AddTos(tos...)

	p.SetDynamicTemplateData("name", name)
	p.SetDynamicTemplateData("headImage", feed.Image.URL)
	_, week := time.Now().ISOWeek()
	p.SetDynamicTemplateData("subject", fmt.Sprintf("%s vecka %d", name, week))

	var itemList []map[string]string
	var item map[string]string
	for _, v := range items {
		item = make(map[string]string)
		item["text"] = v.Description
		item["article"] = v.Title

		if v.Image != nil {
			item["image"] = v.Image.URL
		}

		itemList = append(itemList, item)
	}
	p.SetDynamicTemplateData("news", itemList)

	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	_, err := sendgrid.API(request)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
