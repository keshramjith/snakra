package s3service

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *S3Client) AddObject(vnString, key string) error {

	putObjInput := &s3.PutObjectInput{
		Bucket: aws.String(c.sbn),
		Key:    aws.String(key),
		Body:   strings.NewReader(vnString),
	}
	_, err := c.c.PutObject(context.TODO(), putObjInput)
	if err != nil {
		return err
	}
	fmt.Printf("Added object to s3 bucket: %s\n", c.sbn)
	return nil
}

func (c *S3Client) RemoveObject(key string) error {
	fmt.Printf("Remove object from s3 bucket")
	return nil
}

func (c *S3Client) GetObject(key string) error {
	fmt.Printf("Get object from s3 bucket")
	return nil
}
