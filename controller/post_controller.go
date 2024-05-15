package controller

import (
	"net/http"
	"strconv"

	"Blogging_App/service"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{postService: postService}
}

func (pc *PostController) GetAllPosts(c *gin.Context) {
	posts, err := pc.postService.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (pc *PostController) GetPostsPaginated(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.Query("pageNo"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	posts, err := pc.postService.GetPostsPaginated(pageNo, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch paginated posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (pc *PostController) GetPostsPaginatedByTitle(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.Query("pageNo"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	posts, err := pc.postService.GetPostsPaginatedByTitle(pageNo, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch paginated posts by title"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (pc *PostController) GetPostsPaginatedByAuthorID(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.Query("pageNo"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	posts, err := pc.postService.GetPostsPaginatedByAuthorID(pageNo, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch paginated posts by author ID"})
		return
	}
	c.JSON(http.StatusOK, posts)
}
