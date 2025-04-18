package utils

import (
	"context"
	appconfig "gatorcan-backend/config" // renamed import
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

// PublishSNSMessage sends a message to AWS SNS using the provided config and logger.
func PublishSNSMessage(ctx context.Context, logger *log.Logger, cfg *appconfig.AppConfig, message string, subject string) error {
	awsCfg, err := InitAWSSession(ctx, cfg)
	if err != nil {
		logger.Printf("Failed to initialize AWS session: %v", err)
		return err
	}

	client := sns.NewFromConfig(awsCfg)

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	input := &sns.PublishInput{
		Message:  &message,
		Subject:  &subject,
		TopicArn: &cfg.SNSConfig.TopicARN,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"Environment": {
				DataType:    aws.String("String"),
				StringValue: aws.String(cfg.Environment),
			},
		},
	}

	_, err = client.Publish(ctxWithTimeout, input)
	if err != nil {
		logger.Printf("Failed to publish SNS message: %v", err)
		return err
	}

	logger.Println("SNS message published successfully.")
	return nil
}
