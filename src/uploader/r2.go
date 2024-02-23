package uploader

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/charmbracelet/log"
)

type R2 struct{}

func (r2 *R2) Upload(ctx context.Context, upload Upload) (error, string) {
	log.SetPrefix("R2 uploader")
	var bucketName = os.Getenv("S3_BUCKET_NAME")

	client := newR2Client(ctx)

	path := fmt.Sprintf("craftviews/%s", upload.Screenshot.ID)

	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
		Body:   upload.Screenshot.File,
	})

	if err != nil {
		log.Fatal("Error while uploading screenshot to s3", err)
		return err, ""
	}

	objectURL := url.URL{
		Scheme: "https",
		Host:   os.Getenv("S3_PUBLIC_DOMAIN"),
		Path:   fmt.Sprintf("/%s", path),
	}

	log.Info("Screenshot uploaded to R2.", "URL", objectURL.String())

	return nil, objectURL.String()
}

func newR2Client(ctx context.Context) *s3.Client {
	var accountId = os.Getenv("S3_ACCOUNT_ID")
	var accessKeyId = os.Getenv("S3_ACCESS_KEY_ID")
	var accessKeySecret = os.Getenv("S3_ACCESS_KEY_SECRET")

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}

	return s3.NewFromConfig(cfg)
}
