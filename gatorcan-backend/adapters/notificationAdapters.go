package adapters

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/config"
	"gatorcan-backend/utils"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

func PublishSNSMessage(ctx context.Context, logger *log.Logger, cfg *config.AppConfig, input dtos.SNSMessageDTO) error {
	awsCfg, err := utils.InitAWSSession(ctx, cfg)
	if err != nil {
		logger.Printf("failed to initialize AWS session: %v", err)
		return err
	}

	client := sns.NewFromConfig(awsCfg)

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	params := &sns.PublishInput{
		Message:  &input.Message,
		Subject:  &input.Subject,
		TopicArn: &cfg.SNSConfig.TopicARN,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"Environment": {
				DataType:    awsString("String"),
				StringValue: awsString(cfg.Environment),
			},
		},
	}

	_, err = client.Publish(ctxWithTimeout, params)
	if err != nil {
		logger.Printf("failed to publish SNS message: %v", err)
		return err
	}

	logger.Println("✅ SNS message published")
	return nil
}

func awsString(s string) *string {
	return &s
}
