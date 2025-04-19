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
	ac.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Extract courseID from URL parameters
	courseIDParam := c.Param("cid")
	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		ac.logger.Printf("Invalid course ID: %s", courseIDParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Call the service layer to fetch assignments
	assignments, err := ac.assignmentService.GetAssignmentsByCourseID(ctx, courseID)
	if err == errors.ErrAssignmentNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignments not found"})
		return
	} else if err != nil {
		ac.logger.Printf("error getting assignments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting assignments"})
		return
	}
	// Return the assignments as JSON response
	c.JSON(http.StatusOK, gin.H{"assignments": assignments})
}
func (ac *AssignmentController) GetAssignment(c *gin.Context) {
	ac.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Extract courseID from URL parameters
	courseIDParam := c.Param("cid")
	assignmentIDParam := c.Param("aid")
	assignmentID, err := strconv.Atoi(assignmentIDParam)
	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		ac.logger.Printf("Invalid course ID: %s", courseIDParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	// Call the service layer to fetch assignments
	assignments, err := ac.assignmentService.GetAssignmentByIDAndCourseID(ctx, assignmentID, courseID)
	if err == errors.ErrAssignmentNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	} else if err != nil {
		ac.logger.Printf("error getting assignments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting assignments"})
		return
	}
	// Return the assignments as JSON response
	c.JSON(http.StatusOK, gin.H{"assignments": assignments})
}

func (ac *AssignmentController) CreateOrUpdateAssignment(c *gin.Context) {
	ac.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	_, exists := c.Get("username")
	if !exists {
		ac.logger.Println("Unauthorized access")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	courseIDParam := c.Param("cid")
	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		ac.logger.Printf("Invalid course ID: %s", courseIDParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var assignment dtos.CreateOrUpdateAssignmentRequestDTO
	if err := c.ShouldBindJSON(&assignment); err != nil {
		ac.logger.Printf("Invalid body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	assignment.CourseID = uint(courseID)

	response, err := ac.assignmentService.UpsertAssignment(ctx, ac.logger, &assignment)
	if err != nil {
		switch err {
		case errors.ErrCourseNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		case errors.ErrFailedToUpdateAssignment, errors.ErrFailedToCreateAssignment:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Assignment operation failed"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ac *AssignmentController) DeleteAssignment(c *gin.Context) {
	panic("implement me")
}

func (ac *AssignmentController) SubmitAssignment(c *gin.Context) {
	panic("implement me")
}

func (ac *AssignmentController) UploadFileToAssignment(c *gin.Context) {
	ac.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Get username from JWT token
	username, exists := c.Get("username")
	if !exists {
		ac.logger.Printf("Unauthorized access attempt to upload file to assignment")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract courseID and assignmentID from URL parameters
	courseIDParam := c.Param("cid")
	assignmentIDParam := c.Param("aid")

	// You can convert them to integer if needed (e.g., using strconv.Atoi)
	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		ac.logger.Printf("Invalid course ID: %s", courseIDParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	assignmentID, err := strconv.Atoi(assignmentIDParam)
	if err != nil {
		ac.logger.Printf("Invalid assignment ID: %s", assignmentIDParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	// Parse the rest of the request body (e.g., file_url, filename, file_type)
	var uploadData dtos.UploadFileToAssignmentDTO
	// Override the IDs in uploadData with the ones from the URL
	uploadData.CourseID = uint(courseID)
	uploadData.AssignmentID = uint(assignmentID)

	if err := c.ShouldBindJSON(&uploadData); err != nil {
		ac.logger.Printf("Failed to parse request body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Call the service to handle the business logic
	response, err := ac.assignmentService.UploadFileToAssignment(ctx, ac.logger, username.(string), &uploadData)
	if err == errors.ErrUserNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err == errors.ErrCourseNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	} else if err == errors.ErrAssignmentNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to assignment"})
		return
	}

	// Return the response
	c.JSON(http.StatusCreated, response)
}
