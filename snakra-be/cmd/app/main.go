package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/keshramjith/snakra/internal/server"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type response struct {
	Id string `json:"id"`
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

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	srv := server.NewServer()
	infoLog.Printf("Starting server on %s", srv.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

	// client := setupS3()

	// 	app.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 		getObjInput := s3.GetObjectInput{
	// 			Bucket: aws.String("snakra-test"),
	// 			Key:    aws.String(chi.URLParam(r, "id")),
	// 		}
	// 		output, err := client.GetObject(context.TODO(), &getObjInput)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		data, err := io.ReadAll(output.Body)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		w.Write(data)
	// 	})

	// app.Post("/", func(w http.ResponseWriter, r *http.Request) {

	// putObjInput := &s3.PutObjectInput{
	// 	Bucket: aws.String("snakra-test"),
	// 	Key:    aws.String(newRecording.ID.String()),
	// 	Body:   reader,
	// }
	// 	_, err = client.PutObject(context.TODO(), putObjInput)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	data := &response{Id: newRecording.ID.String()}
	// 		render.JSON(w, r, "post audio endpoint")
	// 	})

	// http.ListenAndServe(":3333", app)
}
