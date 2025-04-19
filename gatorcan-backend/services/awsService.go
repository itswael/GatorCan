package services

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/adapters"
	"gatorcan-backend/config"
	"gatorcan-backend/interfaces"
	"log"
)

type AWSService struct {
	httpClient interfaces.HTTPClient
	config     *config.AppConfig
}

func NewAWSService(httpClient interfaces.HTTPClient, config *config.AppConfig) interfaces.AWSService {
	return &AWSService{
		httpClient: httpClient,
		config:     config,
	}
}

func (s *AWSService) PushNotificationToSNS(ctx context.Context, logger *log.Logger, message string) error {
	input := dtos.SNSMessageDTO{
		Subject: "Notification", // or dynamically set this if needed
		Message: message,
	}

	err := adapters.PublishSNSMessage(ctx, logger, s.config, input)
	if err != nil {
		logger.Printf("error pushing notification to SNS: %v", err)
		return err
	}

	logger.Println("✅ PushNotificationToSNS: SNS message sent successfully")
	return nil
}
