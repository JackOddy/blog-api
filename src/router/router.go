package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name, Method, Pattern string
	Handler               http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
