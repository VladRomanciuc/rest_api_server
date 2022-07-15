package dbapi

import (
	"context"
	"log"
	"time"
	"github.com/VladRomanciuc/Go-classes/api/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/spf13/viper"
)

//Declare the name of the collection
var mongoDB = "api-go"
var mongoColl = "posts"
//define the struct of the collection
type mongoDbColl struct{
	mongoColl string
}
//Constructor for the operations on named collection
func NewMongoOps() models.DbOps{
	return &mongoDbColl{
		mongoColl: mongoColl,
	}
}

//Viper function to read the env file and get the keys
func getMongoEnv(key string) string {
	//Set the file
	viper.SetConfigFile(".env")
	//read the file and handle the errors
	err := viper.ReadInConfig()
	if err != nil {
	  log.Fatalf("Error while reading config file %s", err)
	}
	//Get the required value call and error handling
	value, ok := viper.Get(key).(string)
	if !ok {
	  log.Fatalf("Invalid type assertion")
	}
	//return the value of the key
	return value
}

//Create the mongo db client
func mongoDBClient() *mongo.Client {
	//set the version
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	//load the configuration and access keys
	clientOptions := options.Client().ApplyURI(getMongoEnv("MONGODB_URI")).SetServerAPIOptions(serverAPIOptions)
	//Declare context and the timeout for connection
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Prepare the client and handle error
	client, err := mongo.Connect(c, clientOptions)
	if err != nil {
    	log.Fatal(err)
	}
	//return the client
	return client
}

//Add post function
func (mongo *mongoDbColl) AddPost(post *models.Post) (*models.Post, error) {
	//Get the context
	c := context.Background()
	// Get a new DynamoDB client
	mongoDbClient := mongoDBClient()
	//Access the named collection
	coll := mongoDbClient.Database(mongoDB).Collection(mongoColl)
	//Mapping the post using internal mongo json decoder
	doc := bson.D{{"id", post.Id},{"title", post.Title},{"text", post.Text}}
	//add the post to the collection
	_, err := coll.InsertOne(c, doc)
	//Handle the error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//return the post and error
	return post, err
}
//Get all function
func (mongo *mongoDbColl) GetAll() ([]models.Post, error) {
	//Get the context
	c := context.Background()
	// Get a new DynamoDB client
	mongoDbClient := mongoDBClient()
	//Access the named collection
	coll := mongoDbClient.Database(mongoDB).Collection(mongoColl)
	//Get result by passing an empty json
	result, err := coll.Find(c, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Declare a slice of post type
	var posts []models.Post = []models.Post{}
	//Looping the result, decoding to the post struct
	for result.Next(c){
		post := models.Post{}
		d :=result.Decode(&post)
		if d != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	//Stop when there are no results
	defer result.Close(c)
	//return slice and nil error
	return posts, nil
}
//Function to get the post by id
func (mongo *mongoDbColl) GetById(id string) (*models.Post, error) {
	//Get the context
	c := context.Background()
	// Get a new DynamoDB client
	mongoDbClient := mongoDBClient()
	//Access the named collection
	coll := mongoDbClient.Database(mongoDB).Collection(mongoColl)
	//Mapping the post id using internal mongo json decoder
	doc := bson.D{
		{"id", id},
	}
	var post models.Post
	//get the post by id and map it to post
	err := coll.FindOne(c, doc).Decode(&post)
	//Handle the error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//return the post and error
	return &post, nil
}
//Function to delete the post
func (mongo *mongoDbColl) DeleteById(id string) (error) {
	//Get the context
	c := context.Background()
	// Get a new DynamoDB client
	mongoDbClient := mongoDBClient()
	//Access the named collection
	coll := mongoDbClient.Database(mongoDB).Collection(mongoColl)
	//Mapping the post id using internal mongo json decoder
	docId := bson.D{
		{"id", id},
	}
	//delete the post from the collection
	_, err := coll.DeleteOne(c, docId)
	//Handle the error
	if err != nil {
		log.Fatal(err)
		return err
	}
	//return the post and error
	return nil
}
