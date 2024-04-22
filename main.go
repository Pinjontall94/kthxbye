package main

import(
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("edd.ca.gov"),
	)
	ca_warn_url := "https://edd.ca.gov/en/jobs_and_training/Layoff_Services_WARN"

	// place a callback on every a tag with an href
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print each link found
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on the page
		// Only visit links in the allowed domains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.Visit(ca_warn_url)
	fmt.Println("henlo! :3")
}
