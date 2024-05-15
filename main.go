package main

import (
	"Blogging_App/controller"
	"Blogging_App/repository"
	"Blogging_App/service"
	"Blogging_App/utility"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
)

func main() {

	utility.DbConnect()

	// Initialize repository, service, and controller
	postRepo := repository.PostRepoImpl{} // Initialize your actual repository implementation
	postService := service.NewPostServiceImpl(postRepo)
	postController := controller.NewPostController(postService)

	//initialised gin Engine
	r := gin.Default()
	r.Use(corsMiddleware())
	r.Use(helmet.Default())
	// Defined routes
	r.GET("/posts", postController.GetAllPosts)                                        //http://localhost:8080/posts
	r.GET("/posts/paginated", postController.GetPostsPaginated)                        // e.g. http://localhost:8080/posts/paginated?pageNo=2&pageSize=2
	r.GET("/posts/paginated-by-title", postController.GetPostsPaginatedByTitle)        //http://localhost:8080/posts/paginated-by-title?pageNo=1&pageSize=3
	r.GET("/posts/paginated-by-author-id", postController.GetPostsPaginatedByAuthorID) //http://localhost:8080/posts/paginated-by-author-id?pageNo=3&pageSize=3

	// Start the server in local system
	r.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.Header("Access-Control-Allow-Origin", "*")
		//Enabling CORS for only GET http method
		context.Header("Access-Control-Allow-Methods", "GET")
		context.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, X-CSRF-Token")
	}
}
