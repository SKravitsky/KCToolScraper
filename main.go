package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

type toolLink struct {
	desc string
	link string
}

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func scrape() toolLink {
	log.Println("Starting Scraping")
	var tools toolLink

	c := colly.NewCollector(
		colly.AllowedDomains("kctool.com", "www.kctool.com"),
		colly.MaxDepth(1),
	)

	c.OnHTML(".card-figure a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "card-figure__link" {
			tools.desc = e.Attr("aria-label")
			tools.link = e.Attr("href")
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL.String())
	})

	c.Visit("https://www.kctool.com/tool-of-the-day")
	return tools
}

func mail(t *toolLink) {
	from := os.Getenv("MAILFROM")
	password := os.Getenv("GPW")
	to := []string{os.Getenv("MAILTO")}

	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	tp, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("From: Tool API \nSubject: This is the new Tool of the Day \n%s\n\n", mimeHeaders)))

	tp.Execute(&body, struct {
		Desc string
		Link string
	}{
		Desc: t.desc,
		Link: t.link,
	})

	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, body.Bytes())

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Email Sent!")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Make sure one exists!")
	}

	tools := scrape()
	mail(&tools)
}
