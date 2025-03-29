package controllers

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/errors"
	"gatorcan-backend/interfaces"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SubmissionController struct {
	submissionService interfaces.SubmissionService
	userService       interfaces.UserService
	logger            *log.Logger
}

func NewSubmissionController(service interfaces.SubmissionService, userService interfaces.UserService, logger *log.Logger) *SubmissionController {
	return &SubmissionController{
		submissionService: service,
		userService:       userService,
		logger:            logger,
	}
}

func (sc *SubmissionController) GradeSubmission(c *gin.Context) {
	sc.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Get username from JWT token
	username, exists := c.Get("username")
	if !exists {
		sc.logger.Printf("Unauthorized access attempt to upload file to assignment")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract courseID from URL parameters
	courseIDParam := c.Param("cid")
	_, err := strconv.Atoi(courseIDParam)
	if err != nil {
		sc.logger.Printf("Invalid course ID: %s", courseIDParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Bind the request body to the GradeSubmissionRequestDTO struct
	var submissionData dtos.GradeSubmissionRequestDTO
	if err := c.ShouldBindJSON(&submissionData); err != nil {
		sc.logger.Printf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call the service layer to grade the submission
	usernameStr, ok := username.(string)
	if !ok {
		sc.logger.Printf("Invalid username type")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username type"})
		return
	}
	response, err := sc.submissionService.GradeSubmission(ctx, sc.logger, usernameStr, &submissionData)
	if err == errors.ErrSubmissionNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	} else if err != nil {
		sc.logger.Printf("Error grading submission: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error grading submission"})
		return
	}

	// Return the response as JSON
	c.JSON(http.StatusOK, response)
}

func (sc *SubmissionController) GetSubmission(c *gin.Context) {
	sc.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Get username from JWT token
	username, exists := c.Get("username")
	if !exists {
		sc.logger.Printf("Unauthorized access")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// get userid from user service
	user, err := sc.userService.GetUserDetails(ctx, username.(string))
	if err != nil {
		sc.logger.Printf("Error fetching user ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user ID"})
		return
	}

	// Extract courseID and assignmentID from URL parameters
	courseIDParam := c.Param("cid")
	assignmentIDParam := c.Param("aid")

	// You can convert them to integer if needed (e.g., using strconv.Atoi)
	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		sc.logger.Printf("Invalid course ID: %s", courseIDParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	assignmentID, err := strconv.Atoi(assignmentIDParam)
	if err != nil {
		sc.logger.Printf("Invalid assignment ID: %s", assignmentIDParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	submissions, err := sc.submissionService.GetSubmission(ctx, courseID, assignmentID, user.ID)
	if err == errors.ErrSubmissionNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	} else if err != nil {
		sc.logger.Printf("Error fetching submission: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching submission"})
		return
	}

	// Return the response as JSON
	c.JSON(http.StatusOK, submissions)
}
