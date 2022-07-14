package dbapi

import (

	"log"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/spf13/viper"
)

//Declare the name of the table
var tableName = "posts"
//Struct for the table
type dynamoDBTable struct {
	tableName string
}
//Viper function to read the env file and get the keys
func getEnv(key string) string {
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
//Constructor for the operations on named table
func NewDynamoDB() models.DbOps {
	return &dynamoDBTable{
		tableName: tableName,
	}
}
//DynamoDB Client
func createDynamoDBClient() *dynamodb.Client {
	//Get the context
	c := context.Background()
	//Load the client with default config and handle the error
	cfg, err := config.LoadDefaultConfig(c,
		config.WithRegion(getEnv("Region")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	//Start the client with access key from env file
	awsClient := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.Credentials = credentials.NewStaticCredentialsProvider(getEnv("KeyID"), getEnv("Key"), "")
	})
	//return the client
	return awsClient
}
//Add post function
func (table *dynamoDBTable) AddPost(post *models.Post) (*models.Post, error) {
	//Get the context
	c := context.Background()
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()
	//Mapping the post to map[string] using types.AttributeValue and assign the values
	_, err := dynamoDBClient.PutItem(c, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: post.Id},
			"title": &types.AttributeValueMemberS{Value: post.Title},
			"text": &types.AttributeValueMemberS{Value: post.Text},
			},	
	})
	//Handle the error
	if err != nil {
		return nil, err
	}
	//return pos and the error
	return post, err
}
//Get all function
func (table *dynamoDBTable) GetAll() ([]models.Post, error) {
	//Get the context
	c := context.Background()
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()
	//Query the named table as params
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	//Scan the table and handle the errors
	result, err := dynamoDBClient.Scan(c, params)
	if err != nil {
		return nil, err
	}
	//Declare a slice of post type
	var posts []models.Post = []models.Post{}
	//Looping the result, unmarshaling to the post struct
	for _, i := range result.Items {
		post := models.Post{}
		err = attributevalue.UnmarshalMap(i, &post)
		//if error when unmarsheling
		if err != nil {
			panic(err)
		}
		//add each post to the slice
		posts = append(posts, post)
	}
	//return the slice, and nil error
	return posts, nil
}
//Function to get the post by id
func (table *dynamoDBTable) GetById(id string) (*models.Post, error) {
	//Get the context
	c := context.Background()
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()
	//Mapping the post to map[string] using types.AttributeValue and assign the values for id (Get operation)
	result, err := dynamoDBClient.GetItem(c, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: *aws.String(id)},
		},
	})
	//Error handler
	if err != nil {
		return nil, err
	}
	//Declare post as Post struct
	post := models.Post{}
	//Unmarshal the result to post struct and handle the error
	err = attributevalue.UnmarshalMap(result.Item, &post)
	if err != nil {
		panic(err)
	}
	//return the post and nil error
	return &post, nil
}
//Function to delete the post
func (table *dynamoDBTable) DeleteById(id string) (*models.Post, error) {
	//Get the context
	c := context.Background()
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()
	//Mapping the post to map[string] using types.AttributeValue and assign the values for id (Delete operation)
	_, err := dynamoDBClient.DeleteItem(c, &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: *aws.String(id)},
		},
	})
	//Error handler
	if err != nil {
		return nil, err
	}
	//Return nil for both values
	return nil, nil
}