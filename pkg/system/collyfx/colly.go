package collyfx

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// ProvideColly to fx
func ProvideColly() *colly.Collector {
	c := colly.NewCollector()
	extensions.Referer(c)
	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	return c
}
