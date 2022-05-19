package routes

import (
	"github.com/gorilla/mux"
	"org.Magassians/services/crawler"
)

func crawlerRouter(routes *mux.Router) {

	routes.HandleFunc("/crawl", crawler.Crawl()).Methods("POST")

}
