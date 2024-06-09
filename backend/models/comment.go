package models

import "time"

type Comment struct {
	PostID     uint64    `db:"post_id" json:"post_id"`
	ParentID   uint64    `db:"parent_id" json:"parent_id"`
	CommentID  uint64    `db:"comment_id" json:"comment_id"`
	AuthorID   uint64    `db:"author_id" json:"author_id"`
	Content    string    `db:"content" json:"content"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
	HeadImg    string    `db:"headImg" json:"headImg"`
	Name       string    `db:"name" json:"name"`
	Analyse    uint8     `db:"analyse" json:"analyse"`
	CommentIDstr string  `db:"-" json:"commentIDstr"`
}
