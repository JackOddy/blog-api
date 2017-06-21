package main

import (
	"blog-api/router"
	"net/http"
)

func main() {
	app := router.NewRouter()

	defer http.ListenAndServe(":9000", app)
	println("-->")
	println("server started on localhost:9000")
}
