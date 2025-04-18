package utils

import (
	"fmt"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/config"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env: %v", err)
	}

	logger := log.New(os.Stdout, "[SNS Test] ", log.LstdFlags)
	cfg := config.LoadConfig()

	input := dtos.SNSMessageDTO{
		Message: "This is a test SNS message from adapter ✉️",
		Subject: "Test SNS Subject",
	}

	fmt.Printf("cfg.SNSConfig: %v\n", cfg.SNSConfig)
	fmt.Printf("Input: %+v\n", input)
	fmt.Printf("Logger: %v\n", logger)
	// if err := adapters.PublishSNSMessage(context.Background(), logger, cfg, input); err != nil {
	// 	logger.Fatalf("❌ Failed to publish SNS message: %v", err)
	// }
}
