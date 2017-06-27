package blogs

import (
	"blog-api/src/redis"
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
	return Blog{
		getMetaData(slug),
		getContent("blogs/" + slug),
	}
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

func marshalMeta(blogMeta BlogMeta) (meta []byte) {
	meta, err := json.Marshal(blogMeta)

	if err != nil {
		panic(err)
	}

	return
}

func SetBlog(blog *Blog) {
	redis.Client.MSetNX(
		blog.Meta.Slug, marshalMeta(blog.Meta),
		blog.Meta.Key, blog.Content,
	)
}
func exists(key string) bool {

	t, _ := redis.Client.Exists(key).Result()

	if t == 1 {
		return true
	}

	return false
}

func UpdateBlog(key string, blog Blog) bool {
	if exists(key) {
		redis.Client.MSet(
			key, marshalMeta(blog.Meta),
			"blogs/"+key, blog.Content,
		)
		return true
	}
	return false
}

func DeleteBlog(slug string) bool {
	t, _ := redis.Client.Del(slug, "blogs/"+slug).Result()

	if t != 0 {
		return true
	}

	return false
}
