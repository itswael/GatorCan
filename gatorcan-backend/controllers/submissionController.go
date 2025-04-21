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
	awsService        interfaces.AWSService
	logger            *log.Logger
}

func NewSubmissionController(service interfaces.SubmissionService, userService interfaces.UserService, awsService interfaces.AWSService, logger *log.Logger) *SubmissionController {
	return &SubmissionController{
		submissionService: service,
		userService:       userService,
		awsService:        awsService,
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

	notificationMessage := "Submission graded: " + strconv.Itoa(int(submissionData.AssignmentID)) + " for user: " + usernameStr
	err = sc.awsService.PushNotificationToSNS(ctx, sc.logger, notificationMessage)
	if err != nil {
		sc.logger.Printf("Failed to send SNS notification: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send SNS notification"})
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

func (sc *SubmissionController) SubmitAssignment(c *gin.Context) {
	sc.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	// ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	// defer cancel()
	ctx := c

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

	err = sc.submissionService.Submit(ctx, sc.logger, user.ID, uint(assignmentID), uint(courseID))
	if err == errors.ErrAssignmentFileNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrAssignmentFileNotFound.Error()})
		return
	} else if err != nil {
		sc.logger.Printf("Error submitting assignment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrSubmittingAssignment})
		return
	}

	// Return the response as JSON
	c.JSON(http.StatusOK, gin.H{"message": "Assignment submitted successfully"})
}

func (sc *SubmissionController) GetSubmissions(c *gin.Context) {
	sc.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Get username from JWT token
	_, exists := c.Get("username")
	if !exists {
		sc.logger.Printf("Unauthorized access")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
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

	submissions, err := sc.submissionService.GetSubmissions(ctx, courseID, assignmentID)
	if err == errors.ErrSubmissionNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrSubmissionNotFound.Error()})
		return
	} else if err != nil {
		sc.logger.Printf("Error fetching submission: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFetchingSubmissions})
		return
	}

	// Return the response as JSON
	c.JSON(http.StatusOK, submissions)
}

func (sc *SubmissionController) GetGrades(c *gin.Context) {
	sc.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Get username from JWT token
	username, exists := c.Get("username")
	if !exists {
		sc.logger.Printf("Unauthorized access")
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUnauthorized.Error()})
		return
	}

	// get userid from user service
	user, err := sc.userService.GetUserDetails(ctx, username.(string))
	if err != nil {
		sc.logger.Printf("Error fetching user ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFailedToFetch.Error()})
		return
	}

	// Extract courseID from URL parameters
	courseIDParam := c.Param("cid")
	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		sc.logger.Printf("Invalid course ID: %s", courseIDParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidCourseID.Error()})
		return
	}

	grades, err := sc.submissionService.GetGrades(ctx, sc.logger, courseID, user.ID)
	if err == errors.ErrSubmissionNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrSubmissionNotFound.Error()})
		return
	} else if err != nil {
		sc.logger.Printf("Error fetching submission: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFailedToFetch.Error()})
		return
	}

	// Return the response as JSON
	c.JSON(http.StatusOK, grades)
}
