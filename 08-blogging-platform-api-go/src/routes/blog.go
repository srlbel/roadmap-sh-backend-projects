package routes

import (
	"blogging-platform-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine) {
	blogGroup := router.Group("/blogs")
	{
		blogGroup.GET("/", controllers.GetBlogs)
		blogGroup.GET("/:id", controllers.GetBlogById)
		blogGroup.POST("/", controllers.PostBlogs)
		blogGroup.PUT("/:id", controllers.UpdateBlogById)
		blogGroup.DELETE("/:id", controllers.DeleteBlogById)
	}

	postGroup := router.Group("/posts")
	{
		postGroup.GET("/", controllers.GetBlogsByTag)
	}
}
