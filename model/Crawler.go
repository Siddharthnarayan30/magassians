package model

type UrlInfo struct {
	UrlList []string `json:"urls"`
}

type CrawlDatas struct {
	Url  string `json:"url"`
	Data string `json:"data"`
}

type Response struct {
	Result []CrawlDatas `json:"result"`
	Error  string       `json:"error"`
}
