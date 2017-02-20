package models

import "time"

type Blog struct { //首字母必须大写
	UUID      string
	Author    string
	PostTime time.Time
	Title     string
	Tag	  string
	Content   string
	Comment   []BlogComment
}

type BlogComment struct {
	UUID 	string
	Author  string
	Mail 	string
	Tel 	string
	Content string
	PostTime time.Time
}
