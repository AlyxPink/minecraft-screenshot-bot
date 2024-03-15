package uploader

import (
	"bytes"
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

type R2 struct {
	objectURL  url.URL
	objectPath string
}

func (r2 *R2) Upload(ctx context.Context, upload Upload) (error, string) {
	// Set the object fields
	r2.setObjectFields(upload)

	// Upload the screenshot to R2
	client := newR2Client(ctx)
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key:         aws.String(r2.objectPath),
		Body:        bytes.NewReader(upload.Screenshot.File),
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		log.Fatal("Error while uploading screenshot to R2", err)
		return err, ""
	}

	return nil, r2.objectURL.String()
}

func newR2Client(ctx context.Context) *s3.Client {
	var (
		accountId       = os.Getenv("R2_ACCOUNT_ID")
		accessKeyId     = os.Getenv("R2_ACCESS_KEY_ID")
		accessKeySecret = os.Getenv("R2_ACCESS_KEY_SECRET")
	)

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

func (r2 *R2) setObjectFields(upload Upload) {
	r2.objectPath = r2.getObjectPath(upload)
	r2.objectURL = r2.getObjectURL()
}

func (r2 *R2) getObjectPath(upload Upload) string {
	return filepath.Join(
		os.Getenv("R2_PATH_PREFIX"),
		fmt.Sprint(upload.Screenshot.ID.String(), filepath.Ext(upload.Screenshot.Name)),
	)
}

func (r2 *R2) getObjectURL() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   os.Getenv("R2_PUBLIC_DOMAIN"),
		Path:   r2.objectPath,
	}
}
