package model

import "time"

type Post struct {
	ID         int64     `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   int64     `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      int64     `json:"likes,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}
