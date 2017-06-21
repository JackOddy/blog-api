package blogs

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Blog Post Index")
}

func Create(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Blog Post Index")
}

func Read(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(res, "Blog Post Read: "+vars["id"])
}

func Update(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(res, "Blog Post Update: "+vars["id"])
}

func Delete(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(res, "Blog Post Delete: "+vars["id"])
}
