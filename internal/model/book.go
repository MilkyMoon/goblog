package model

import (
	"time"
)

type Books []Book

type Book struct {
	Name string `json:"name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type BookTree struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Children [] BookTree `json:"children"`
}

