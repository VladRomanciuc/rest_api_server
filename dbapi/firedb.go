package dbapi

import (
	"context"
	"log"
	"github.com/VladRomanciuc/Go-classes/api/models"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
  
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

//Declare the name of the collection
var collName = "posts"
//define the struct of the collection
type collection struct{
	collName string
}
//Constructor for the operations on named collection
func NewFirestoreOps() models.DbOps{
	return &collection{
		collName: collName,
	}
}
//Create a firestore client
func firestoreClient() *firestore.Client {
	//Get the context
	c := context.Background()
	//Load the access keys file
	opt := option.WithCredentialsFile("C:\\Users\\alina\\Desktop\\Go classes\\api\\serviceAccountKey.json")
	//Load the client with default config and handle the error
	app, err := firebase.NewApp(c, nil, opt)
	if err != nil {
  		log.Fatalln(err)
	}
	//Start the client and handle the error
	client, err := app.Firestore(c)
	if err != nil {
  		log.Fatalln(err)
	}
	//Return the client
	return client
}
//Add post function
func (*collection) AddPost(post *models.Post) (*models.Post, error) {
	//Get the context, start the client, close connection after operation
	c := context.Background()
	fireClient := firestoreClient()
	defer fireClient.Close()
	//Call the client with a named collection, set the generated id and add the post
	_, err := fireClient.Collection(collName).Doc(post.Id).Set(c, map[string]interface{}{
		"Id": post.Id,
		"Title": post.Title,
		"Text": post.Text,
	})
	//Error handler
	if err != nil {
		log.Fatal("Failed adding a new post: %v", err)
		return nil, err
	}
	//Return the post and error nil
	return post, nil
}
//Get all posts function
func (*collection) GetAll() ([]models.Post, error) {
	//Get the context, start the client, close connection after operation
	c := context.Background()
	fireClient := firestoreClient()
	defer fireClient.Close()
	//Declare a slice of post to store the result
	var posts []models.Post
	//Call the client with a named collection and all docs
	coll := fireClient.Collection(collName).Documents(c)
	//Stop iterate when finish
	defer coll.Stop()
	//Iterate each doc and handle errors
	for {
		doc, err := coll.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal("Failed to return the list of posts: %v", err)
			return nil, err
		}
		//Map the value to post struct
		post := models.Post {
			Id: doc.Data()["Id"].(string),
			Title: doc.Data()["Title"].(string),
			Text: doc.Data()["Text"].(string),
		}
		//Adding posts to the slice
		posts = append(posts, post)
	}
	//Return the slice and error nil
	return posts, nil
}
//Get posts by id function
func (*collection) GetById(id string) (*models.Post, error) {
	//Get the context, start the client, close connection after operation
	c := context.Background()
	fireClient := firestoreClient()
	defer fireClient.Close()
	//Call the client with a named collection and id and get the post, handle error
	doc, err := fireClient.Collection(collName).Doc(id).Get(c)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	//Map the value to post struct
	post := &models.Post{
		Id:    doc.Data()["Id"].(string),
		Title: doc.Data()["Title"].(string),
		Text:  doc.Data()["Text"].(string),
	}
	//Return the post and error nil
	return post, nil
}
//Delete posts by id function
func(*collection) DeleteById(id string) (*models.Post, error) {
	//Get the context, start the client, close connection after operation
	c := context.Background()
	fireClient := firestoreClient()
	defer fireClient.Close()
	//Call the client with a named collection and id and delete the post, error handle
	_, e := fireClient.Collection(collName).Doc(id).Delete(c)
	if e != nil {
		log.Fatalf(e.Error())
		return nil, e
	}
	//Return nil for both
	return nil, nil
}