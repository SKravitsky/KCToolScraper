package main

import (
	"log"

	// "net/http"

	"github.com/gocolly/colly/v2"
)

// func ping(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Ping")
// 	w.Write([]byte("ping"))
// }

// func getData(w http.ResponseWriter, r *http.Request) {
// 	URL := "https://www.kctool.com/tool-of-the-day/"
// 	c := colly.NewCollector(colly.MaxDepth(1))

// }

func scrape() {
	log.Println("Starting Scraping")
	c := colly.NewCollector(
		colly.AllowedDomains("kctool.com", "www.kctool.com"),
		colly.MaxDepth(1),
	)

	c.OnHTML(".card-figure a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "card-figure__link" {
			log.Println("Got the description: ", e.Attr("aria-label"))
			log.Println("Link to tool: ", e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL.String())
	})

	c.Visit("https://www.kctool.com/tool-of-the-day")

}

// func crawl() {
// 	log.Println("Testing Crawler")
// 	c := colly.NewCollector(
// 		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
// 	)

// 	c.OnHTML(".factsList li", func(e *colly.HTMLElement) {
// 		factId, err := strconv.Atoi(e.Attr("id"))
// 		if err != nil {
// 			log.Println("Could not get ID")
// 		}
// 		log.Println("Got the FactID of ", factId)

// 		factDesc := e.Text
// 		log.Println("Got the description: ", factDesc)
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		log.Println("Visiting ", r.URL.String())
// 	})

// 	c.Visit("https://www.factretriever.com/rhino-facts")

// }

func main() {
	scrape()
	// crawl()
	// port := ":8888"
	// http.HandleFunc("/ping", ping)

	// log.Println("Listening on port", port)
	// log.Fatal(http.ListenAndServe(port, nil))
}
