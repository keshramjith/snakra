package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type response struct {
	Id string `json:"id"`
}

type Recording struct {
	gorm.Model
	ID        uuid.UUID `sql:"AUTO_INCREMENT" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
}

func setupPg() *gorm.DB {
	dsn := "postgres://kesh:password@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect to db")
	}
	return db
}

func setupS3() *s3.Client {
	// loads profile and credentials from ~/.aws
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		fmt.Println(err)
	}
	client := s3.NewFromConfig(cfg)
	return client
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	client := setupS3()

	db := setupPg()
	db.AutoMigrate(&Recording{})

	app.Get("/:id", func(c *fiber.Ctx) error {
		getObjInput := s3.GetObjectInput{
			Bucket: aws.String("snakra-test"),
			Key:    aws.String(c.Params("id")),
		}
		output, err := client.GetObject(context.TODO(), &getObjInput)
		if err != nil {
			panic(err)
		}
		return c.SendStream(output.Body)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		blob := c.Body()
		reader := bytes.NewReader(blob)

		newRecording := Recording{ID: uuid.New(), CreatedAt: time.Now()}
		db.Create(&newRecording)

		putObjInput := &s3.PutObjectInput{
			Bucket: aws.String("snakra-test"),
			Key:    aws.String(newRecording.ID.String()),
			Body:   reader,
		}
		_, err := client.PutObject(context.TODO(), putObjInput)
		if err != nil {
			panic(err)
		}
		return c.JSON(response{Id: newRecording.ID.String()})
	})

	app.Listen(":3001")
}
