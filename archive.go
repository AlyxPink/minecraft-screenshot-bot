package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/charmbracelet/log"
)

type Metadata struct {
	GeneratedTime string `json:"generated_time"`
	Description   string `json:"description"`
}

func archiveScreenshot(screenshot *os.File) {
	// Create S3 client
	var accountId = os.Getenv("S3_ACCOUNT_ID")
	var accessKeyId = os.Getenv("S3_ACCESS_KEY_ID")
	var accessKeySecret = os.Getenv("S3_ACCESS_KEY_SECRET")

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	// Upload screenshot to S3
	var bucketName = os.Getenv("S3_BUCKET_NAME")
	var path = fmt.Sprintf("CraftViews/%s", screenshot.Name())

	_, err = client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
		Body:   screenshot,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Info(fmt.Sprintf("Successfully uploaded %s to %s/%s", screenshot.Name(), bucketName, path))
}
