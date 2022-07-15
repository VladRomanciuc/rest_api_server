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
	post := models.Post{Id: "1", Title: "", Text: "some"}
	
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
	var identifier string
	post := models.Post{Id: identifier, Title: "some", Text: "some"}
	mockDbOps.On("GetAll").Return([]models.Post{post}, nil)

	testService := NewPostService(mockDbOps)
	response, _ := testService.GetAll()

	mockDbOps.AssertExpectations(t)
	assert.Equal(t, identifier, response[0].Id)
	assert.Equal(t, "some", response[0].Title)
	assert.Equal(t, "some", response[0].Text)
}

func (mock *MockDbOps)	GetById(id string) (*models.Post, error) {
	call := mock.Called()
	response := call.Get(0)
	return response.(*models.Post), call.Error(1)
}

func TestGetById(t *testing.T){
	mockDbOps := new(MockDbOps)
	var identifier string
	post := models.Post{Id: identifier, Title: "some", Text: "some"}
	mockDbOps.On("GetById").Return(&post, nil)

	testService := NewPostService(mockDbOps)
	response, _ := testService.GetById(identifier)

	mockDbOps.AssertExpectations(t)
	assert.Equal(t, identifier, response.Id)
	assert.Equal(t, "some", response.Title)
	assert.Equal(t, "some", response.Text)
}

func (mock *MockDbOps)	DeleteById(id string) (error) {
	call := mock.Called()
	call.Get(0)
	return nil
}
func TestDeleteById(t *testing.T){
	mockDbOps := new(MockDbOps)
	var identifier string
	mockDbOps.On("DeleteById").Return(nil)

	testService := NewPostService(mockDbOps)
	response := testService.DeleteById(identifier)

	mockDbOps.AssertExpectations(t)
	assert.Nil(t, response)
}