package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/keshramjith/snakra/internal/server"
)

func enrichContext(ctx context.Context, key, val string) context.Context {
	return context.WithValue(ctx, key, val)
}

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Failed to find .env file")
		os.Exit(1)
	}
	port := os.Getenv("PORT")
	// env := os.Getenv("ENV")
	s3bucketName := os.Getenv("S3_BUCKET")

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	srv := server.NewServer(infoLog, errorLog, s3bucketName, port)
	infoLog.Printf("Starting server on %s\n", port)
	err := srv.ListenAndServe()
	defer srv.Close()
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
