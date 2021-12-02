package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	fName := "tokpedScrap.csv"

	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Couldn't create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.UserAgent("xy"),
		colly.AllowURLRevisit(),
		colly.AllowedDomains("www.tokopedia.com"),
	)

	c.Visit("https://www.tokopedia.com/p/handphone-tablet/handphone")
	c.OnHTML(".css-16vw0vn", func(h *colly.HTMLElement) {
		log.Println("udah di function on html")
		log.Println(h.ChildText("span"))
		writer.Write([]string{
			h.ChildText("span"),
		})
	})

	for i := 0; i < 3; i++ {

		fmt.Printf("Scraping page: %d\n", i)

		log.Println("ini tokped page", c.Visit("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=23&sc=24&limit=100&page="+strconv.Itoa(i)))
	}

	log.Printf("Scraping Complete\n")
	log.Println(c)
}
