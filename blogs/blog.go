package blogs

import (
	"blog-api/redis"
	"encoding/json"
	"io"
	"sync"
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
	metaChannel := make(chan BlogMeta, 1)
	contentChannel := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go getMetaData(slug, metaChannel, &wg)
	go getContent("blogs/"+slug, contentChannel, &wg)
	wg.Wait()
	meta := <-metaChannel
	content := <-contentChannel
	return Blog{meta, content}
}

func getMetaData(key string, c chan BlogMeta, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := redis.Client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	var meta BlogMeta
	json.Unmarshal([]byte(data), &meta)
	c <- meta
}

func getContent(key string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	content, err := redis.Client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	c <- content
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
