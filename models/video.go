package models

type Video struct {
	Id           int64
	Title        string
	Description  string
	ParentCourse Course
}
