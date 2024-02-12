package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/charmbracelet/log"
)

func uploadToS3(screenshot_path string) url.URL {
	var bucketName = os.Getenv("S3_BUCKET_NAME")

	client := getCloudFlareR2Client()

	// Open screenshot file
	screenshot, err := os.Open(screenshot_path)
	if err != nil {
		log.Fatal(err)
	}

	filename := filepath.Base(screenshot.Name())
	path := fmt.Sprintf("craftviews/%s", filename)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
		Body:   screenshot,
	})

	if err != nil {
		log.Fatal("Error while uploading screenshot to s3", err)
	}

	objectURL := url.URL{
		Scheme: "https",
		Host:   os.Getenv("S3_PUBLIC_DOMAIN"),
		Path:   fmt.Sprintf("/%s", path),
	}

	log.Info("Screenshot uploaded to R2.", "URL", objectURL.String())

	return objectURL
}

func getCloudFlareR2Client() *s3.Client {
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

	return s3.NewFromConfig(cfg)
}
