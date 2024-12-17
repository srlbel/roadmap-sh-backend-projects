package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type blogData struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var createdTime = time.Now().Local()
var updatedTime = time.Now().Local()

var blogs = []blogData{
	{
		ID:        "1",
		Title:     "How I Met Your Mother",
		Content:   "Some Content",
		Category:  "TV",
		Tags:      []string{"Comedy"},
		CreatedAt: createdTime,
		UpdatedAt: updatedTime,
	},
	{
		ID:        "2",
		Title:     "err != nil",
		Content:   "Some Content",
		Category:  "Technology",
		Tags:      []string{"Programming", "Go"},
		CreatedAt: createdTime,
		UpdatedAt: updatedTime,
	},
}

func main() {
	router := gin.Default()
	router.GET("/blogs", getBlogs)
	router.GET("/blogs/:id", getBlogById)
	router.GET("/posts", getBlogsByTag)
	router.POST("/blogs", postBlogs)
	router.PUT("/blogs/:id", updateBlogById)
	router.DELETE("/blogs/:id", deleteBlogById)

	router.Run(":3001")
}

func getBlogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, blogs)
}

func postBlogs(c *gin.Context) {
	time := time.Now().Local()
	var newBlog blogData

	newBlog.CreatedAt, newBlog.UpdatedAt = time, time

	if err := c.BindJSON(&newBlog); err != nil {
		return
	}

	blogs = append(blogs, newBlog)
	c.IndentedJSON(http.StatusCreated, newBlog)
}

func getBlogById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range blogs {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
}

func updateBlogById(c *gin.Context) {
	id := c.Param("id")
	var newBlog blogData

	if err := c.BindJSON(&newBlog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i := range blogs {
		if blogs[i].ID == id {
			blogs[i].Title = newBlog.Title
			blogs[i].Content = newBlog.Content
			blogs[i].Category = newBlog.Category
			blogs[i].Tags = newBlog.Tags
			blogs[i].UpdatedAt = time.Now().Local()

			c.IndentedJSON(http.StatusOK, newBlog)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"error": "Blog not found",
	})
}

func deleteBlogById(c *gin.Context) {
	id := c.Param("id")

	for i := range blogs {
		if blogs[i].ID == id {
			blogs = append(blogs[:i], blogs[i+1:]...)

			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Blog deleted successfully",
			})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"error": "Blog not found",
	})
}

func getBlogsByTag(c *gin.Context) {
	tag := c.Query("term")
	if tag == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "tag query parameter is required"})
		return
	}

	var filteredBlogs []blogData

	for _, blog := range blogs {
		for _, t := range blog.Tags {
			if strings.EqualFold(t, tag) {
				filteredBlogs = append(filteredBlogs, blog)
				break
			}
		}
	}

	if len(filteredBlogs) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No blogs found with the specified tags"})
		return
	}

	c.IndentedJSON(http.StatusOK, filteredBlogs)
}
