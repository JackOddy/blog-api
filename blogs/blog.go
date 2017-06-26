package blogs

import (
	"blog-api/redis"
	"encoding/json"
	"io"
	"time"
)

type BlogMeta struct {
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	Key       string    `json:"-"`
	TimeStamp time.Time `json:"timeStamp"`
	Slug      string    `json:"slug"`
}

type Blog struct {
	Meta    BlogMeta `json:"meta"`
	Content string   `json:"content"`
}

type Blogs []Blog

func GetBlog(slug string) Blog {
	meta := getMetaData(slug)
	content := getContent("blogs/" + slug)
	return Blog{meta, content}
}

func getMetaData(key string) (meta BlogMeta) {
	data, err := redis.Client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(data), &meta)
	return
}

func getContent(key string) (content string) {
	content, err := redis.Client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return
}

func NewBlog(body io.Reader) (blog Blog) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&blog)
	if err != nil {
		panic(err)
	}
	blog.Meta.TimeStamp = time.Now()
	blog.Meta.Key = "blogs/" + blog.Meta.Slug
	return
}

func setMeta(blogMeta *BlogMeta) {
	meta, err := json.Marshal(blogMeta)
	if err != nil {
		panic(err)
	}
	redis.Client.SetNX(blogMeta.Slug, meta, 0)
}

func setContent(blog *Blog) {
	redis.Client.SetNX(blog.Meta.Key, blog.Content, 0)
}

func SetBlog(blog *Blog) {
	setMeta(&blog.Meta)
	setContent(blog)
}
