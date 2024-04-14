package entity

import "time"

type Status = string

const (
	New       Status = "new"
	Processed Status = "processed"
)

type Action = string

const (
	Add    Action = "add"
	Delete Action = "delete"
	Modify Action = "modify"

	AddKeyword    Action = "add_keyword"
	DeleteKeyword Action = "delete_keyword"
)

type DocAction struct {
	Id        string     `bson:"_id,omitempty"`
	Keyword   string     `bson:"keyword,omitempty"`
	ArticleId string     `bson:"article_id,omitempty"`
	Status    Status     `bson:"status"`
	Action    Action     `bson:"action"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"`
	Document  *Document  `bson:"document,omitempty"`
}
