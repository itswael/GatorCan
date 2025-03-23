package controllers

import (
	"gatorcan-backend/interfaces"
	"log"

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
