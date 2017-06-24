package blogs

import (
	"blog-api/redis"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
)

func Index(res http.ResponseWriter, req *http.Request) {
}

func Create(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	blog := Blog{Headers: BlogInfo{TimeStamp: time.Now()}}
	err := decoder.Decode(&blog)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	headers := blog.Headers
	key := "blogs/" + headers.Slug
	headers.Key = key
	blogHeaders, err := json.Marshal(headers)
	redis.Client.Set(headers.Slug, blogHeaders, 0)
	redis.Client.Set(key, blog.Content, 0)
	json.NewEncoder(res).Encode(blog)
}

func Read(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blog, err := redis.Client.Get(vars["id"]).Result()
	if err != nil {
		panic(err)
	}
	var blogInfo BlogInfo
	json.Unmarshal([]byte(blog), &blogInfo)
	content, err := redis.Client.Get(blogInfo.Key).Result()

	constructedBlog := Blog{blogInfo, content}
	fmt.Println(blog)
	json.NewEncoder(res).Encode(constructedBlog)
}

func Update(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(res, "Blog Post Update: "+vars["id"])
}

func Delete(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(res, "Blog Post Delete: "+vars["id"])
}
