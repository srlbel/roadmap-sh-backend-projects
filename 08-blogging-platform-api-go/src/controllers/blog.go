package controllers

import (
	"net/http"
	"strings"
	"time"

	"blogging-platform-api/src/models"

	"github.com/gin-gonic/gin"
)

var blogs = []models.Blog{
	{
		ID:        "1",
		Title:     "How I Met Your Mother",
		Content:   "Some Content",
		Category:  "TV",
		Tags:      []string{"Comedy"},
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	},
	{
		ID:        "2",
		Title:     "err != nil",
		Content:   "Some Content",
		Category:  "Technology",
		Tags:      []string{"Programming", "Go"},
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	},
}

func GetBlogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, blogs)
}

func GetBlogById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range blogs {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
}

func PostBlogs(c *gin.Context) {
	time := time.Now().Local()
	var newBlog models.Blog

	newBlog.CreatedAt, newBlog.UpdatedAt = time, time

	if err := c.BindJSON(&newBlog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogs = append(blogs, newBlog)
	c.IndentedJSON(http.StatusCreated, newBlog)
}

func UpdateBlogById(c *gin.Context) {
	id := c.Param("id")
	var updatedBlog models.Blog

	if err := c.BindJSON(&updatedBlog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range blogs {
		if blogs[i].ID == id {
			blogs[i].Title = updatedBlog.Title
			blogs[i].Content = updatedBlog.Content
			blogs[i].Category = updatedBlog.Category
			blogs[i].Tags = updatedBlog.Tags
			blogs[i].UpdatedAt = time.Now().Local()

			c.IndentedJSON(http.StatusOK, blogs[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
}

func DeleteBlogById(c *gin.Context) {
	id := c.Param("id")

	for i := range blogs {
		if blogs[i].ID == id {
			blogs = append(blogs[:i], blogs[i+1:]...)

			c.IndentedJSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
}

func GetBlogsByTag(c *gin.Context) {
	tag := c.Query("term")
	if tag == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "tag query parameter is required"})
		return
	}

	var filteredBlogs []models.Blog

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
