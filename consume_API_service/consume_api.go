package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogPost struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func fetchPostsFromBloggingApp(url string) ([]BlogPost, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var posts []BlogPost
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func main() {
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		url := "http://localhost:8080/posts"
		posts, err := fetchPostsFromBloggingApp(url)
		if err != nil {
			log.Println("Error fetching posts:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	r.GET("/posts/paginated", func(c *gin.Context) {
		query := c.Request.URL.Query()
		pageNo := query.Get("pageNo")
		pageSize := query.Get("pageSize")
		url := "http://localhost:8080/posts/paginated?pageNo=" + pageNo + "&pageSize=" + pageSize
		posts, err := fetchPostsFromBloggingApp(url)
		if err != nil {
			log.Println("Error fetching paginated posts:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch paginated posts"})
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	r.GET("/posts/paginated-by-title", func(c *gin.Context) {
		query := c.Request.URL.Query()
		pageNo := query.Get("pageNo")
		pageSize := query.Get("pageSize")
		url := "http://localhost:8080/posts/paginated-by-title?pageNo=" + pageNo + "&pageSize=" + pageSize
		posts, err := fetchPostsFromBloggingApp(url)
		if err != nil {
			log.Println("Error fetching paginated posts by title:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch paginated posts by title"})
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	r.GET("/posts/paginated-by-author-id", func(c *gin.Context) {
		query := c.Request.URL.Query()
		pageNo := query.Get("pageNo")
		pageSize := query.Get("pageSize")
		url := "http://localhost:8080/posts/paginated-by-author-id?pageNo=" + pageNo + "&pageSize=" + pageSize
		posts, err := fetchPostsFromBloggingApp(url)
		if err != nil {
			log.Println("Error fetching paginated posts by author ID:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch paginated posts by author ID"})
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	r.Run(":8081") //running on port 8081 and consuming API from 8080
}
