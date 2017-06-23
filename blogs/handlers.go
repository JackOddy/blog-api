package blogs

import (
	"blog-api/redis"
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
)

func Index(res http.ResponseWriter, req *http.Request) {
}

func Create(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	blog := BlogHeader{TimeStamp: time.Now()}
	err := decoder.Decode(&blog)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	mappedBlog := structs.Map(blog)
	redis.Client.HMSet("blogs."+blog.Slug, mappedBlog)
	json.NewEncoder(res).Encode(blog)
}

func Read(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blog, err := redis.Client.HGetAll("blogs." + vars["id"]).Result()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(res).Encode(blog)
}

func Update(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(res, "Blog Post Update: "+vars["id"])
}

func Delete(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(res, "Blog Post Delete: "+vars["id"])
}
