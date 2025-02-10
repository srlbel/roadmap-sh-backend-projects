package controllers

import (
	"net/http"
	"time"

	"blogging-platform-api/src/database"
	"blogging-platform-api/src/models"

	"github.com/gin-gonic/gin"
)

func GetBlogs(c *gin.Context) {
	var blogs []models.Blog
	database.DB.Find(&blogs)

	c.IndentedJSON(http.StatusOK, blogs)
}

func GetBlogById(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog

	if err := database.DB.First(&blog, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, blog)
}

func PostBlogs(c *gin.Context) {
	var newBlog models.Blog

	if err := c.BindJSON(&newBlog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newBlog)
	c.IndentedJSON(http.StatusCreated, newBlog)
}

func UpdateBlogById(c *gin.Context) {
	id := c.Param("id")
	var updatedBlog models.Blog

	if err := database.DB.First(&updatedBlog, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
		return
	}

	if err := c.BindJSON(&updatedBlog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBlog.UpdatedAt = time.Now()
	database.DB.Save(&updatedBlog)

	c.IndentedJSON(http.StatusOK, updatedBlog)

}

func DeleteBlogById(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog

	if err := database.DB.First(&blog, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
	}

	if err := database.DB.Delete(&blog).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func GetBlogsByTag(c *gin.Context) {
	tag := c.Query("term")
	if tag == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "tag query parameter is required"})
		return
	}

	var filteredBlogs []models.Blog
	if err := database.DB.Where("tag = ?", tag).Find(&filteredBlogs).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}

	if len(filteredBlogs) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No blogs found with the specified tags"})
		return
	}

	c.IndentedJSON(http.StatusOK, filteredBlogs)
}
