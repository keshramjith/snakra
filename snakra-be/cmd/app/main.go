package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/keshramjith/snakra/internal/server"
)

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
	defer srv.CloseDb()
	errorLog.Fatal(err)
}
