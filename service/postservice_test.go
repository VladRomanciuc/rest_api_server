package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/VladRomanciuc/Go-classes/api/models"
)

func TestValidateEmptyPost(t *testing.T){
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "post is empty", err.Error())
}

func TestValidateEmptyTitle(t *testing.T){
	post := models.Post{Id: 1, Title: "", Text: "some"}
	
	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "title is empty", err.Error())
}

type MockDbOps struct {
	mock.Mock	
}

func (mock *MockDbOps) AddPost(post *models.Post) (*models.Post, error) {
	call := mock.Called()
	response := call.Get(0)
	return response.(*models.Post), call.Error(1)
}

func TestAddPost(t *testing.T){
	mockDbOps := new(MockDbOps)
	
	post := models.Post{Title: "some", Text: "some"}

	mockDbOps.On("AddPost").Return(&post, nil)

	testService := NewPostService(mockDbOps)
	response, err := testService.AddPost(&post)

	mockDbOps.AssertExpectations(t)
	assert.NotNil(t, response.Id)
	assert.Equal(t, "some", response.Title)
	assert.Equal(t, "some", response.Text)
	assert.Nil(t, err)
}

func (mock *MockDbOps)	GetAll() ([]models.Post, error) {
	call := mock.Called()
	response := call.Get(0)
	return response.([]models.Post), call.Error(1)
}

func TestGetAll(t *testing.T){
	mockDbOps := new(MockDbOps)
	var identifier int64
	post := models.Post{Id: identifier, Title: "some", Text: "some"}
	mockDbOps.On("GetAll").Return([]models.Post{post}, nil)

	testService := NewPostService(mockDbOps)
	response, _ := testService.GetAll()

	mockDbOps.AssertExpectations(t)
	assert.Equal(t, identifier, response[0].Id)
	assert.Equal(t, "some", response[0].Title)
	assert.Equal(t, "some", response[0].Text)
}