package main

import (
	"blog-api/blogs"
	"io"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"content-type",
		"application/json",
	)
	io.WriteString(res, "index")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/blogs/", blogs.HandleReq)
	defer http.ListenAndServe(":9000", nil)
	println("-->")
	println("server started on localhost:9000")
}
