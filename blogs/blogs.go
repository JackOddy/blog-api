package blogs

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
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
		err, response = getBlogs(req)
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

func getBlogs(req *http.Request) (error, string) {
	x := Blog{"hello", "this is a great blog"}

	blog, err := json.Marshal(x)

	// extract path variables from the URL - returns []string
	params := strings.Split(req.URL.Path, "/")
	for _, x := range params {
		println(x)
	}
	return err, string(blog)
}
