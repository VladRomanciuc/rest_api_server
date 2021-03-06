package controller

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"io"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/VladRomanciuc/Go-classes/api/service"
	"github.com/VladRomanciuc/Go-classes/api/dbapi"
	"github.com/VladRomanciuc/Go-classes/api/cache"
)

var (
	Id string = "123"
	Title string = "Title 1"
	Text string = "Text 1"
)

var (
	dbo models.DbOps = dbapi.NewSQLiteDb()
	postServ models.PostService = service.NewPostService(dbo)
	postCach models.PostCache = cache.NewRedisCache("localhost:49154", "redispw", 0, 360)
	postCont models.PostController = NewPostController(postServ, postCach)
)

func TestAddPost(t *testing.T){
	//http post request
	jsonB := []byte(`{"title":"` + Title + `", "text":"` + Text +`"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonB))

	//handler
	handler :=http.HandlerFunc(postCont.AddPost)

	//assign the response
	response := httptest.NewRecorder()

	//make the call
	handler.ServeHTTP(response, req)

	//add assertion
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("wrong status code: %v", status)
	}

	//decode json
	var post models.Post

	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	//assert json
	assert.NotNil(t, post.Id)
	assert.Equal(t, Title, post.Title)
	assert.Equal(t, Text, post.Text)

	//delete
	deletePost(post.Id)
}

func deletePost(id string) {
	var post models.Post = models.Post{
		Id: id,
	}
	dbo.DeleteById(post.Id)
}

func setup() {
	var post models.Post = models.Post{
		Id:    Id,
		Title: Title,
		Text:  Text,
	}
	dbo.AddPost(&post)
}


func TestGetAll(t *testing.T){

	// Insert new post
	setup()

	//http get request
	req, _ := http.NewRequest("GET", "/posts", nil)

	//handler
	handler := http.HandlerFunc(postCont.GetAll)

	//assign the response
	response := httptest.NewRecorder()

	//make the call
	handler.ServeHTTP(response, req)

	//add assertion
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("wrong status code: %v", status)
	}

	//decode json
	var posts []models.Post

	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	//assert json
	assert.Equal(t, Id, posts[0].Id)
	assert.Equal(t, Title, posts[0].Title)
	assert.Equal(t, Text, posts[0].Text)

	deletePost(Id)
}