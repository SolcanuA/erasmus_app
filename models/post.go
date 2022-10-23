package models

type Course struct {
	Id          int64
	Author      User
	Title       string
	Description string
}
