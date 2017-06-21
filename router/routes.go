package router

import (
	"blog-api/blogs"
)

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
