package controllers

import (
    "net/http"
    "backend/services"
    "github.com/gin-gonic/gin"
)

func GetChatByID(c *gin.Context) {
    id := c.Param("id")
    chat, err := services.GetChat(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"chat": chat})
}