package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`
	Author          primitive.ObjectID   `bson:"author,omitempty"`
	Title           string               `bson:"title,omitempty"`
	Description     string               `bson:"description,omitempty"`
	CourseLength    int64                `bson:"course_length,omitempty"`
	CourseWallpaper string               `bson:"course_wallpaper,omitempty"`
	CourseURL       string               `bson:"course_url,omitempty"`
	Videos          []primitive.ObjectID `bson:"course_videos"`
	Reviews         []primitive.ObjectID `bson:"reviews"`
	EnrolledUsers   []primitive.ObjectID `bson:"course_enrolled_users"`
}
