package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Author         primitive.ObjectID `bson:"review_author,omitempty"`
	ReveiwedCourse primitive.ObjectID `bson:"reviewed_course,omitempty"`
	Rating         uint8              `bson:"rating,omitempty"`
	Content        string             `bson:"review_content,omitempty"`
}
