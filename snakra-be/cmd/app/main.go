package main

import (
	"flag"
	"log"
	"os"

	"github.com/keshramjith/snakra/internal/server"
)

type config struct {
	port int
	env  string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 3001, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|staging|prod)")

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	srv := server.NewServer()
	infoLog.Printf("Starting server on %s", cfg.port)
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
