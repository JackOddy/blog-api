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
	result := "Could not Update"

	vars := mux.Vars(req)
	blog := NewBlog(req.Body)
	success := UpdateBlog(vars["slug"], blog)

	if success {
		result = "Updated"
	}

	io.WriteString(res, result)
}

func Delete(res http.ResponseWriter, req *http.Request) {
	result := "Could not Delete"

	vars := mux.Vars(req)
	success := DeleteBlog(vars["slug"])

	if success {
		result = "Deleted"
	}

	io.WriteString(res, result)
}
