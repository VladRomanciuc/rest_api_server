package main

import (
	"github.com/VladRomanciuc/Go-classes/api/service"
	"github.com/VladRomanciuc/Go-classes/api/dbapi"
	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/VladRomanciuc/Go-classes/api/router"
	"github.com/VladRomanciuc/Go-classes/api/controller"
	"github.com/VladRomanciuc/Go-classes/api/cache"
)

var (
	//DB switch
	dbops models.DbOps = dbapi.NewFirestoreOps() //NO ERRORS
	//dbops models.DbOps = dbapi.NewSQLiteDb() //NO ERRORS
	//dbops models.DbOps = dbapi.NewDynamoDB() //NO ERRORS

	//Post service require the db
	postService models.PostService = service.NewPostService(dbops) //NO ERRORS
	//Cache layer needs credentials
	postCache models.PostCache = cache.NewRedisCache("localhost:49154", "redispw", 0, 360) //NO ERRORS

	//The controller requieres both post and cache services
	postController models.PostController = controller.NewPostController(postService, postCache) //NO ERRORS
	
	//Router switch
	//api models.Router = router.NewRouterMux() //NO ERRORS
	api models.Router = router.NewRouterChi() //NO ERRORS

	//Example of gathering info from differnt API
	carDetailsService models.CarDetailsService = service.NewCarDetailsService()
	carDetailsController models.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
)

func main() {
    port := ":8080"
	//Basic ops with a post
	api.GET("/posts", postController.GetAll)
	api.POST("/posts", postController.AddPost)
	api.GET("/posts/{id}", postController.GetById)
	api.DELETE("/posts/{id}", postController.DeleteById)
	//Gather info from 2 differnt json
    api.GET("/cardetails", carDetailsController.GetCarDetails)
	api.SERVE(port)
}