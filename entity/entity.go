package entity

import "time"

type NewsFullDetailed struct {
	ID      int       `json:"ID"`
	Title   string    `json:"Title"`
	Content string    `json:"Content"`
	Link    string    `json:"Link"`
	PubTime int64     `json:"PubTime"`
	Comment []Comment `json:"Comment"`
}

type Comment struct {
	ID       string    `json:"id,omitempty"`
	ParentId string    `json:"parent_id,omitempty"`
	NewsId   string    `json:"news_id"`
	Content  string    `json:"content"`
	Created  time.Time `json:"created,omitempty"`
	Reply    []Comment `json:"reply,omitempty"`
}

type NewsShortDetailed struct {
	ID      int    `json:"ID"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Link    string `json:"Link"`
	PubTime int64  `json:"PubTime"`
}
