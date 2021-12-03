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
	//Write header
	writer.Write([]string{"Name", "Price", "Toko"})

	c := colly.NewCollector(
		colly.UserAgent("xy"),
		colly.AllowURLRevisit(),
		colly.AllowedDomains("www.tokopedia.com", "tokopedia.com"),
	)

	var productName, productPrice, namaToko, imageLink string

	// productDesc, productRating

	c.Visit("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=23&sc=24&limit=100")
	c.OnHTML(".css-16vw0vn", func(h *colly.HTMLElement) {
		productName = h.ChildText("span.css-1bjwylw")
		productPrice = h.ChildText("span.css-o5uqvq")
		namaToko = h.ChildText("span.css-1kr22w3:last-child")
		imageLink = h.ChildAttr("img", "src")
		log.Println(imageLink)
	})

	c.OnHTML(".css-89jnbj", func(h *colly.HTMLElement) {
		a := h.Request.Visit(h.Attr("href"))

		log.Println(a)
	})

	// c.OnHTML("img[src]", func(h *colly.HTMLElement) {
	// 	imageLink = h.Request.AbsoluteURL(h.Attr("src"))
	// 	log.Println(imageLink)
	// })

	writer.Write([]string{
		productName,
		productPrice,
		namaToko,
	})
	for i := 0; i < 3; i++ {

		fmt.Printf("Scraping page: %d\n", i)

		log.Println("ini tokped page", c.Visit("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=23&sc=24&limit=100&page="+strconv.Itoa(i)))
	}

	log.Printf("Scraping Complete\n")
	log.Println(c)
}
