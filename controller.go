package fetch

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/chromedp/cdproto/dom"
	_ "github.com/chromedp/cdproto/emulation"
	_ "github.com/chromedp/chromedp"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	// respdata, err := http.Get("https://ace.tokopedia.com/search/product/v3?device=desktop&ob=23&q=handphone&rows=100&source=shop_product&start=0")
	respdata, err := http.Get("https://www.tokopedia.com/p/handphone-tablet/handphone")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(response{
			Message: "Failed to fetch products",
		})
	}

	responseData, err := ioutil.ReadAll(respdata.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(response{
			Message: "Failed to encode products",
		})
	}

	log.Println(string(responseData))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Message: "Hello World!",
		Data:    responseData,
	})
}
