package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id      primitive.ObjectID `json:"id,omiempty"`
	Title   string             `json:"title,omitempty"`
	Content string             `json:"content,omiempty"`
}
