package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type response struct {
	Id string `json:"id"`
}

type userDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	app := chi.NewRouter()
	app.Use(middleware.Logger)
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowOriginFunc:  nil,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	client := setupS3()

	db := setupPg()
	db.AutoMigrate(&Recording{})

	app.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		getObjInput := s3.GetObjectInput{
			Bucket: aws.String("snakra-test"),
			Key:    aws.String(chi.URLParam(r, "id")),
		}
		output, err := client.GetObject(context.TODO(), &getObjInput)
		if err != nil {
			panic(err)
		}
		data, err := io.ReadAll(output.Body)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(data)
	})

	app.Post("/", func(w http.ResponseWriter, r *http.Request) {
		blob, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		reader := bytes.NewReader(blob)

		newRecording := Recording{ID: uuid.New(), CreatedAt: time.Now()}
		db.Create(&newRecording)

		putObjInput := &s3.PutObjectInput{
			Bucket: aws.String("snakra-test"),
			Key:    aws.String(newRecording.ID.String()),
			Body:   reader,
		}
		_, err = client.PutObject(context.TODO(), putObjInput)
		if err != nil {
			panic(err)
		}

		data := &response{Id: newRecording.ID.String()}
		render.JSON(w, r, data)
	})

	http.ListenAndServe(":3333", app)
}
