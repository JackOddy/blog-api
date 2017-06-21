package router

import (
	"blog-api/blogs"
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

var routes = Routes{
	Route{
		"BlogPostIndex",
		"GET",
		"/blog-posts",
		blogs.Index,
	},
	Route{
		"BlogPostRead",
		"GET",
		"/blog-posts/{id}",
		blogs.Read,
	},
	Route{
		"BlogPostCreate",
		"POST",
		"/blog-posts",
		blogs.Create,
	},
	Route{
		"BlogPostUpdate",
		"PUT",
		"/blog-posts/{id}",
		blogs.Update,
	},
	Route{
		"BlogPostDelete",
		"DELETE",
		"/blog-posts/{id}",
		blogs.Delete,
	},
}
