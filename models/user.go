package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty"`
	FirstName         string               `bson:"first_name,omitempty"`
	LastName          string               `bson:"last_name,omitempty"`
	Username          string               `bson:"username,omitempty"`
	Email             string               `bson:"email,omitempty"`
	Password          string               `bson:"password,omitempty"`
	Followers         []primitive.ObjectID `bson:"followers"`
	Following         []primitive.ObjectID `bson:"following"`
	CoursesOwned      []primitive.ObjectID `bson:"courses_owned"`
	CoursesEnrolledIn []primitive.ObjectID `bson:"courses_enrolled_in"`
}
