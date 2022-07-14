package models

import (
	"net/http"
)
//DB operations
type DbOps interface {
	AddPost(post *Post) (*Post, error)
	GetAll() ([]Post, error)
	GetById(id string) (*Post, error)
	DeleteById(id string) (*Post, error)
}
//PostService actions
type PostService interface{
	Validate(post *Post) error
	AddPost(post *Post) (*Post, error)
	GetAll() ([]Post, error)
	GetById(id string) (*Post, error)
	DeleteById(id string) (*Post, error)
}
//Controller handler
type PostController interface{
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
}

//Router handler
type Router interface {
	GET(url string, f func(w http.ResponseWriter, r *http.Request))
	POST(url string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(url string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

//Redis cache actions
type PostCache interface {
	Set(key string, value *Post)
	Get(key string) *Post
	Del(key string) error
}

//The structure of data to be handled + a json mapper for encoding/decoding
type Post struct{
	Id 		string	`json:"Id"`
	Title	string	`json:"Title"`
	Text 	string	`json:"Text"`
}
//A struct for Error messages
type ServiceError struct{
	Message string `json:"message"`
}

