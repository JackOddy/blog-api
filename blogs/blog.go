package blogs

import (
	"time"
)

type BlogInfo struct {
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	Key       string    `json:"key"`
	TimeStamp time.Time `json:"timeStamp"`
	Slug      string    `json:"slug"`
}

type Blog struct {
	Headers BlogInfo `json:"headers"`
	Content string   `json:"content"`
}

type Blogs []Blog
