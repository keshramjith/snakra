package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type response struct {
	Msg string `json:"msg"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// aws s3 setup
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		fmt.Println(err)
	}

	client := s3.NewFromConfig(cfg)

	app.Get("/", func(c *fiber.Ctx) error {
		resp := response{Msg: "connection test success"}
		return c.JSON(resp)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		blob := c.Body()
		reader := bytes.NewReader(blob)
		putObjInput := &s3.PutObjectInput{
			Bucket: aws.String("snakra-test"),
			Key:    aws.String("anidnian"),
			Body:   reader,
		}
		putObjectOuput, err := client.PutObject(context.TODO(), putObjInput)
		if err != nil {
			panic(err)
		}
		fmt.Println(putObjectOuput)
		return c.JSON(response{Msg: "recording received!"})
	})

	app.Listen(":3001")
}
