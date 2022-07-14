package cache

import (
	"encoding/json"
	"time"
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/VladRomanciuc/Go-classes/api/models"
)

//Redis cache structure for client
type RedisCache struct {
	host    string
	db      int
	pswd    string
	expires time.Duration
}
//Mapping resis cache with function
func NewRedisCache(host string, pswd string, db int, exp time.Duration) models.PostCache {
	return &RedisCache{
		host:    host,
		db:      db,
		pswd:	pswd,
		expires: exp,
	}
}

//Start a new redis client with options
func (cache *RedisCache) getClient() *redis.Client {
	
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: cache.pswd,
		DB:       cache.db,
	})
}

//Function to set the entry in redis cache
func (cache *RedisCache) Set(key string, post *models.Post) {
	//assign the client and context
	client := cache.getClient()
	c := context.Background()
	//serialize Post object to JSON and handle the error
	json, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}
	//Write in cache data
	client.Set(c, key, json, cache.expires*time.Second)
}

//Function to get the entry from cache
func (cache *RedisCache) Get(key string) *models.Post {
	//assign the client and context
	client := cache.getClient()
	c := context.Background()
	//Get the entry by id and handle the error
	val, err := client.Get(c, key).Result()
	if err != nil {
		return nil
	}
	//Variable post of Post struct to store the data
	post := models.Post{}
	//Unmarshal result and handle the error
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}
	//returning data
	return &post
}
//Function to delete a entry from cashe
func (cache *RedisCache) Del(key string) error {
	//assign the client and context
	client := cache.getClient()
	c := context.Background()
	//call the delete method and hanle the error
	_, err := client.Del(c, key).Result()
	if err != nil {
		panic(err)
		return nil
	}
	//return no error if succesuful
	return nil
}
