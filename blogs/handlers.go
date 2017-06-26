package blogs

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
}

func Create(res http.ResponseWriter, req *http.Request) {
	blog := NewBlog(req.Body)
	SetBlog(&blog)
	json.NewEncoder(res).Encode(blog)
}

func Read(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blog := GetBlog(vars["slug"])

	json.NewEncoder(res).Encode(blog)
}

func Update(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(res, "Blog Post Update: "+vars["slug"])
}

func Delete(res http.ResponseWriter, req *http.Request) {
	result := "Successfully Deleted"
	vars := mux.Vars(req)
	_, err := DeleteBlog(vars["slug"])

	if err != nil {
		result = "Error: Could not Delete"
	}

	io.WriteString(res, result)
}
