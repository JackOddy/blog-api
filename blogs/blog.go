package blogs

import (
	"time"
)

type BlogHeader struct {
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	TimeStamp time.Time `json:"timeStamp"`
	Slug      string    `json:"slug"`
}

type Blogs []BlogHeader
