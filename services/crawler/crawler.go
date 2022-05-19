package crawler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
	"org.Magassians/model"
	"org.Magassians/util"
)

func Crawl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var (
			res  model.Response
			urls model.UrlInfo
		)

		decodeError := json.NewDecoder(r.Body).Decode(&urls)
		if decodeError != nil {

			res.Error = " JSON format problem"
			responseBody, err := json.Marshal(res)
			util.CheckError(err)
			util.Log.Info().Msgf(" JSON format problem")
			w.WriteHeader(400)
			fmt.Fprintf(w, "%s", responseBody)

			return
		}

		urls.UrlList = util.FilterDuplicateItem(urls.UrlList)

		for _, url := range urls.UrlList {
			CrawledData := crawlingOperation(url)
			if len(CrawledData) < 1 {
				CrawledData = "Wrong URL provide"
			}
			res.Result = append(res.Result, model.CrawlDatas{Url: url, Data: CrawledData})
		}

		res.Error = "null"
		responseBody, err := json.Marshal(res)
		util.CheckError(err)

		util.Log.Info().Msgf("%s", responseBody)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", responseBody)
	}
}

func crawlingOperation(urlLink string) string {
	c := colly.NewCollector(colly.Async(true))
	var CrawledData string

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		subUrlLink := e.Attr("href")
		CrawledData += subUrlLink + "\n"
	})

	c.OnError(func(r *colly.Response, err error) {
		util.Log.Error().Msgf("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(urlLink)
	c.Wait()
	return CrawledData
}
