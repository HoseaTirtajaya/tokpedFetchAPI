package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// run task list
	var res []string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.tokopedia.com/p/handphone-tablet/handphone`),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Click(`.css-1dq1dix e1nlzfl1`, chromedp.NodeVisible)
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}