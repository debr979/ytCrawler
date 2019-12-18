package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {
	ChromeDp()
}

func ChromeDp() {
	var videoUrls []map[string]string
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var channelName string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.youtube.com/user/kamiyu666/videos`),
		chromedp.Text(`div[id="inner-header-container"] > div[id="meta"] > ytd-channel-name[id="channel-name"] > div[id="container"] > div[id="text-container"] > yt-formatted-string[id="text"]`, &channelName, chromedp.ByQueryAll),
		chromedp.AttributesAll(`a[id="thumbnail"]`, &videoUrls, chromedp.ByQueryAll),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(channelName)
	for i, url := range videoUrls {
		if url["href"] != "" {
			log.Printf("%d:%s", i+1, "https://www.youtube.com"+url["href"])
		}
	}
}


