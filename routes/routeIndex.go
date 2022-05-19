package routes

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"
	midware "org.Magassians/middleware"
)

func RouterIndex() *mux.Router {

	routes := mux.NewRouter().StrictSlash(true)

	routes.Use(middleware.RequestID)
	routes.Use(middleware.RealIP)
	routes.Use(middleware.Logger)
	routes.Use(midware.Recovery)

	crawlerRouter(routes.PathPrefix("/api").Subrouter())

	return routes
}
