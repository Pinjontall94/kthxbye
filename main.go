package main

import(
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/xuri/excelize/v2"
)

func main() {
	base_url := "https://edd.ca.gov"
	warn_report_url := base_url + "/en/jobs_and_training/Layoff_Services_WARN"
	warn_xlsx_url := base_url + "/siteassets/files/jobs_and_training/warn"
	pattern := regexp.QuoteMeta(warn_xlsx_url) + "/warn_report1\\.xlsx"
	strings_lst := [...]string{base_url, warn_report_url, warn_xlsx_url, pattern}
	for _, str := range strings_lst {
		fmt.Printf("%s\n", str)
	}

	c := colly.NewCollector(
		colly.AllowedDomains("edd.ca.gov"),
		colly.URLFilters(regexp.MustCompile(pattern)),
	)

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

	c.OnResponse(func(r *colly.Response) {
		warn_report := excelize.NewFile()
		defer func() {
			if err := warn_report.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		//err := warn_report.Read(r.Body)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.Visit(warn_report_url)
	fmt.Println("henlo! :3")
}
