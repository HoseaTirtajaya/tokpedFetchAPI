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
		colly.AllowedDomains("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=23&sc=24&limit=100"),
	)

	log.Println("abis ini on html")
	c.OnHTML(".css-16vw0vn", func(h *colly.HTMLElement) {
		log.Println("udah di function on html")
		log.Println(h.ChildText("span"))
		writer.Write([]string{
			h.ChildText("span class=css-1bjwylw"),
		})
	})
	log.Println(c.Visit("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=23&sc=24&limit=100"))

	for i := 0; i < 3; i++ {

		fmt.Printf("Scraping page: %d\n", i)

		c.Visit("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=23&sc=24&limit=100&page=" + strconv.Itoa(i))
	}

	log.Printf("Scraping Complete\n")
	log.Println(c)
}
