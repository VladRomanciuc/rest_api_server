package service

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/VladRomanciuc/Go-classes/api/models"
)

//Assign db operations
var db models.DbOps

type service struct{}

//constructor for db operations
func NewPostService(dbops models.DbOps) models.PostService{
	db = dbops
	return &service{}
}

//Validator function for empty post and title
func (*service) Validate(post *models.Post) error {
	if post == nil {
		err := errors.New("post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("title is empty")
		return err
	}
	return nil
}
//add post function
func (*service) AddPost(post *models.Post) (*models.Post, error) {
	//generate an id with random number
	i := rand.Int63()
	//Convert to string
	post.Id = strconv.FormatInt(i, 10)
	return db.AddPost(post)
}

func (*service) GetAll() ([]models.Post, error) {
	return db.GetAll()
}

func (*service) GetById(id string) (*models.Post, error) {
	return db.GetById(id)
}

func (*service) DeleteById(id string) (*models.Post, error) {
	return db.DeleteById(id)
}