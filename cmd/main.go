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

	var ProductName, ProductPrice, NamaToko, ImageLink string

	// productDesc, productRating, imageLink

	c.Visit("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=23&sc=24&limit=100")
	c.OnHTML(".css-89jnbj", func(h *colly.HTMLElement) {
		// log.Println(h.Request.AbsoluteURL(h.Attr("href")))
		ProductName = h.ChildText("span.css-1bjwylw")
		ProductPrice = h.ChildText("span.css-o5uqvq")
		NamaToko = h.ChildText("span.css-1kr22w3:last-child")
		ImageLink = h.ChildAttr("css-t8frx0>img.success.fade", "src")
		log.Println(ImageLink)

		writer.Write([]string{
			ProductName,
			ProductPrice,
			NamaToko,
		})
	})

	// c.OnHTML(".css-89jnbj", func(h *colly.HTMLElement) {
	// 	a := h.Request.Visit(h.Attr("href"))

	// 	log.Println(a)
	// })

	// c.OnHTML("img[src]", func(h *colly.HTMLElement) {
	// 	imageLink = h.Request.AbsoluteURL(h.Attr("src"))
	// 	log.Println(imageLink)
	// })

	for i := 0; i < 15; i++ {

		fmt.Printf("Scraping page: %d\n", i)

		log.Println("ini tokped page", c.Visit("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=23&sc=24&limit=100&page="+strconv.Itoa(i)))
	}

	log.Printf("Scraping Complete\n")
	log.Println(c)
}
