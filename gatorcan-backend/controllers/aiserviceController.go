package controllers

import (
	"context"
	"gatorcan-backend/errors"
	"gatorcan-backend/interfaces"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AIServiceController struct {
	aiServiceService interfaces.AiServiceService
	logger           *log.Logger
}

func NewAIServiceController(aiServiceService interfaces.AiServiceService, logger *log.Logger) *AIServiceController {
	return &AIServiceController{
		aiServiceService: aiServiceService,
		logger:           logger,
	}
}

func (ac *AIServiceController) GetCourseRecommendations(c *gin.Context) {
	ac.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Get username from JWT token
	username, exists := c.Get("username")
	if !exists {
		ac.logger.Printf("Unauthorized access attempt to get course recommendations")
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUnauthorized.Error()})
		return
	}

	// Call the service layer to fetch course recommendations
	courseRecommendations, err := ac.aiServiceService.GetCourseRecommendations(ctx, ac.logger, username.(string))
	if err != nil {
		if err == errors.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrUserNotFound.Error()})
		} else if err == errors.ErrFailedToFetch {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFailedToFetch.Error()})
		} else if err == errors.ErrMicroserviceNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrMicroserviceNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrMicroserviceDown.Error()})
		}
		ac.logger.Printf("error getting course recommendations: %v", err)
		// c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrMicroserviceDown})
		return
	}

	// Return the course recommendations as JSON response
	c.JSON(http.StatusOK, courseRecommendations)
}
