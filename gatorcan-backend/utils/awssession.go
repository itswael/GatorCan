package utils

import (
	"context"
	"fmt"
	"gatorcan-backend/config"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func InitAWSSession(ctx context.Context, cfg *config.AppConfig) (aws.Config, error) {
	fmt.Println("cfg.SNSConfig:", cfg.SNSConfig)
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var opts []func(*awsconfig.LoadOptions) error

	opts = append(opts, awsconfig.WithRegion(cfg.SNSConfig.Region))

	if cfg.SNSConfig.AccessKeyID != "" && cfg.SNSConfig.SecretAccessKey != "" {
		customCreds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			cfg.SNSConfig.AccessKeyID,
			cfg.SNSConfig.SecretAccessKey,
			cfg.SNSConfig.SessionToken,
		))
		opts = append(opts, awsconfig.WithCredentialsProvider(customCreds))
	}

	return awsconfig.LoadDefaultConfig(ctxWithTimeout, opts...)
}
