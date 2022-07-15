package dbapi

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/VladRomanciuc/Go-classes/api/models"
)

/*
cqlsh
CREATE KEYSPACE apigo WITH replication = {'class':'SimpleStrategy', 'replication_factor': 1};
USE apigo;
CREATE TABLE IF NOT EXISTS posts (id text PRIMARY KEY, title text, text text);
*/

//Declare the name of the table
var cTable = "posts"
//Struct for the table
type cassandraDbKSpace struct {
	cTable string
}

//Constructor for the operations on named table
func NewCassandraDB() models.DbOps {
	return &cassandraDbKSpace{
		cTable: cTable,
	}
}
//Cassandra Client
func CassandraDBClient() *gocql.Session {
	//Call ne cluster and set the address
	cluster := gocql.NewCluster("127.0.0.1")
	//Declare the working keyspace
	cluster.Keyspace = "apigo"
	//Create a new session
	client, err := cluster.CreateSession()
	//Error handler
	if err != nil {
		log.Fatal(err)
	}
	//return the client
	return client
}
//Add post function
func (kSpace *cassandraDbKSpace) AddPost(post *models.Post) (*models.Post, error) {
	// Get a new Cassandra client
	cassandraClient := CassandraDBClient()
	//Add the post to DB
	err := cassandraClient.Query("INSERT INTO posts (id, title, text) VALUES(?, ?, ?);",
	post.Id, post.Title, post.Text).Exec()
	//Error handler
	if err != nil{
		log.Fatalf("Error while inserting: %v", err)
	}
	//return pos and the error
	return post, err
}
//Get all function
func (kSpace *cassandraDbKSpace) GetAll() ([]models.Post, error) {
	// Get a new Cassandra client
	cassandraClient := CassandraDBClient()
	////Declare a slice of posts to store the data
	var posts []models.Post
	//Declare a map interface
	m := map[string]interface{}{}
	//Get the result with iterator
	iter := cassandraClient.Query("SELECT * FROM posts;").Iter()
	//Mapping the result and store in post slice
	for iter.MapScan(m){
		posts = append(posts, models.Post{
			Id: m["id"].(string),
			Title: m["title"].(string),
			Text: m["text"].(string),
			})
			m = map[string]interface{}{}
		}
	//return the slice, and nil error
	return posts, nil
}
//Function to get the post by id
func (kSpace *cassandraDbKSpace) GetById(id string) (*models.Post, error) {
	// Get a new Cassandra client
	cassandraClient := CassandraDBClient()
	//Declare a slice of posts to store the data	
	var posts []models.Post
	//Declare a map interface
	m := map[string]interface{}{}
	//Get the result with iterator
	result := cassandraClient.Query("SELECT id, title, text FROM posts WHERE id= ?;", id).Iter()
	//Mapping the result and store in post slice
	for result.MapScan(m){
		posts = append(posts, models.Post{
			Id: m["id"].(string),
			Title: m["title"].(string),
			Text: m["text"].(string),
			})
			m = map[string]interface{}{}
		}
	//Declare post as Post struct
	post := models.Post{}
	//If posts is empty return nil for both values else return the first post store in slice
	if posts != nil {
		post=posts[0]
		return &post, nil
	} else {
		return nil, nil
	}
}
//Function to delete the post
func (kSpace *cassandraDbKSpace) DeleteById(id string) (error) {
	// Get a new Cassandra client
	cassandraClient := CassandraDBClient()
	//Execute delete statement by id
	err := cassandraClient.Query("DELETE FROM posts WHERE id = ?;", id).Exec()
	//Error handler
	if err != nil {
		log.Fatal(err)
		return err
	}
	//Return nil for both values
	return nil
}