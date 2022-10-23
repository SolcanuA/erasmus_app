package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                primitive.ObjectID   `json:"_id,omitempty"        bson:"_id,omitempty"`
	FirstName         string               `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName          string               `json:"last_name,omitempty"  bson:"last_name,omitempty"`
	Username          string               `json:"username,omitempty"   bson:"username,omitempty"`
	Email             string               `json:"email,omitempty"      bson:"email,omitempty"`
	Password          string               `json:"password,omitempty"   bson:"password,omitempty"`
	Age               int8                 `json:"age,omitempty"        bson:"age,omitempty"`
	Followers         []primitive.ObjectID `json:"followers"  		    bson:"followers"`
	Following         []primitive.ObjectID `json:"following"            bson:"following"`
	CoursesOwned      []primitive.ObjectID `json:"courses_owned"        bson:"courses_owned"`
	CoursesEnrolledIn []primitive.ObjectID `json:"courses_enrolled_in"  bson:"courses_enrolled_in"`
}
