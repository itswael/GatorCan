package utils

import (
	"context"
	"gatorcan-backend/config"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func testsns() {
	// Set up logging
	logger := log.New(os.Stdout, "[SNS Test] ", log.LstdFlags)

	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatalf("Error loading .env file: %v", env_err)
	}

	// Load config from environment
	cfg := config.LoadConfig()

	// Test message and subject
	message := "Hello from Go SNS test! 🚀"
	subject := "SNS Test Subject"

	// Publish SNS message
	err := PublishSNSMessage(context.Background(), logger, cfg, message, subject)
	if err != nil {
		logger.Fatalf("Failed to publish SNS message: %v", err)
	}

	logger.Println("✅ SNS message published successfully!")
}
