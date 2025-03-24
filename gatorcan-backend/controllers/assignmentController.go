package controllers

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/interfaces"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AssignmentController struct {
	assignmentService interfaces.AssignmentService
	logger            *log.Logger
}

func NewAssignmentController(service interfaces.AssignmentService, logger *log.Logger) *AssignmentController {
	return &AssignmentController{
		assignmentService: service,
		logger:            logger,
	}
}

func (ac *AssignmentController) GetAssignments(c *gin.Context) {
	panic("implement me")
}
func (ac *AssignmentController) GetAssignment(c *gin.Context) {
	panic("implement me")
}

func (ac *AssignmentController) CreateAssignment(c *gin.Context) {
	panic("implement me")
}

func (ac *AssignmentController) UpdateAssignment(c *gin.Context) {
	panic("implement me")
}

func (ac *AssignmentController) DeleteAssignment(c *gin.Context) {
	panic("implement me")
}

func (ac *AssignmentController) SubmitAssignment(c *gin.Context) {
	panic("implement me")
}
func (ac *AssignmentController) GradeAssignment(c *gin.Context) {
	panic("implement me")
}

func (ac *AssignmentController) GetSubmission(c *gin.Context) {
	panic("implement me")
}

func (ac *AssignmentController) UploadFilesToAssignment(c *gin.Context) {
	ac.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	usernameStr, ok := username.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve username"})
		return
	}

	var request dtos.UploadFileToAssignmentDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// Call the service to upload the file
	response, err := ac.assignmentService.UploadFileToAssignment(ctx, ac.logger, usernameStr, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}
	c.JSON(http.StatusOK, response)
	return
}
