package routes

import (
    "backend/controllers"
    "github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
    apiGroup := router.Group("/api")
    {
        apiGroup.GET("/chat/:id", controllers.GetChatByID)
    }
}