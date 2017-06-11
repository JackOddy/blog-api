package blogs

import (
	"encoding/json"
	"io"
	"net/http"
)

type Blog struct {
	Title, Body string
}

func HandleReq(res http.ResponseWriter, req *http.Request) {
	var (
		err      error
		response string
	)

	switch req.Method {
	case "GET":
		err, response = getBlog(req.URL.Path)
	case "POST":
		err, response = postBlogs(req)
	default:
		err, response = error(nil), "default"
	}

	if err != nil {
		panic(err)
	}

	io.WriteString(
		res,
		response,
	)
}

func postBlogs(req *http.Request) (error, string) {
	decoder := json.NewDecoder(req.Body)

	var blog Blog
	err := decoder.Decode(&blog)
	println(req.URL.Path)
	return err, "post posted"
}

func getBlog(title string) (error, string) {

	// get the blog post from file or db
	if title == "" {
		title = "hello"
	}

	return error(nil), title
}
