package s3config

import "fmt"

func (c *S3Client) AddObject(obj any, key string) error {
	fmt.Printf("Add object to s3 bucket: %s\n", c.sbn)
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
