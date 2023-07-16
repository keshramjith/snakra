package s3service

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	c   *s3.Client
	sbn string
}

func NewS3Client(s3bn string) *S3Client {
	// loads profile and credentials from ~/.aws
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		fmt.Println(err)
	}
	client := s3.NewFromConfig(cfg)
	return &S3Client{c: client, sbn: s3bn}
}
