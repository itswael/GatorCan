package routes

import (
	"gatorcan-backend/controllers"
	"gatorcan-backend/middleware"
	"gatorcan-backend/models"
	"log"

	"github.com/gin-gonic/gin"
)

func UserRoutes(userController *controllers.UserController, courseController *controllers.CourseController, assignmentController *controllers.AssignmentController, submissionController *controllers.SubmissionController, router *gin.Engine, logger *log.Logger) {

	//  Public Routes
	router.POST("/login", func(c *gin.Context) {
		userController.Login(c, logger)
	})

	// Admin-only Routes
	adminGroup := router.Group("/admin")
	adminGroup.Use(middleware.AuthMiddleware(logger, string(models.Admin)))
	{
		adminGroup.POST("/add_user", func(c *gin.Context) {
			userController.CreateUser(c, logger)
		})
		adminGroup.DELETE("/:username", func(c *gin.Context) {
			userController.DeleteUser(c, logger)
		})
		adminGroup.PUT("/update_role", func(c *gin.Context) {
			userController.UpdateRoles(c, logger)
		})

	}
	userGroup := router.Group("/user")
	userGroup.Use(middleware.AuthMiddleware(logger, string(models.Student), string(models.Admin)))
	{
		userGroup.GET("/:username", func(c *gin.Context) {
			userController.GetUserDetails(c, logger)
		})
		userGroup.PUT("/update", func(c *gin.Context) {
			userController.UpdateUser(c, logger)
		})

	}

	// Instructor-only Routes
	instructorRoutes := router.Group("/instructor")
	instructorRoutes.Use(middleware.AuthMiddleware(logger, string(models.Instructor)))
	{
		//instructorRoutes.POST("/upload-assignment", UploadAssignmentHandler)
		instructorRoutes.POST("/courses/:cid/assignments/:aid/grade", func(c *gin.Context) {
			submissionController.GradeSubmission(c)
		})
	}

	courseGroup := router.Group("/courses")
	courseGroup.Use(middleware.AuthMiddleware(logger, string(models.Student), string(models.Admin)))
	{
		courseGroup.GET("/enrolled", func(c *gin.Context) {
			courseController.GetEnrolledCourses(c)
		})
		//courseGroup.POST("/enroll", controllers.EnrollCourse)
		courseGroup.GET("/", func(c *gin.Context) {
			courseController.GetCourses(c)
		})

		courseGroup.POST("/enroll", func(c *gin.Context) {
			courseController.EnrollInCourse(c)
		})

		courseGroup.GET("/:cid", func(c *gin.Context) {
			courseController.GetCourse(c)
		})
		assignmentGroup := courseGroup.Group("/:cid/assignments")
		assignmentGroup.Use(middleware.AuthMiddleware(logger, string(models.Student), string(models.Instructor), string(models.TA)))
		{
			assignmentGroup.GET("/", func(c *gin.Context) {
				assignmentController.GetAssignments(c)
			})

			assignmentGroup.GET("/:aid", func(c *gin.Context) {
				assignmentController.GetAssignment(c)
			})

			assignmentGroup.POST("/:aid/upload", func(c *gin.Context) {
				assignmentController.UploadFileToAssignment(c)
			})

		}
	}

}
