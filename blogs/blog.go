package blogs

import "time"

type Blog struct {
	Title, Body, Type string
	TimeStamp         time.Time
}

type Blogs []Blog
