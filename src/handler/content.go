package handler

import (
	"net/http"

	"PrivateZone/src/model"

	"github.com/gin-gonic/gin"
)

func CreateContentFile(c *gin.Context) {
	var content model.ContentFile
	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.contentFileService.CreateContentFile(content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create content file"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Content file created successfully"})
}

func CreateContentStream(c *gin.Context) {
	var content model.ContentStream
	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.contentStreamService.CreateContentStream(content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create content stream"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Content stream created successfully"})
}

// Adicione outras funções de controller conforme necessário
