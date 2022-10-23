package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Video struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ParentCourse primitive.ObjectID `bson:"parent_course,omitempty"`
	Title        string             `bson:"video_title,omitempty"`
	Description  string             `bson:"video_description,omitempty"`
	VideoLength  int64              `bson:"video_length,omitempty"`
	VideoURL     string             `bson:"video_url,omitempty"`
}
