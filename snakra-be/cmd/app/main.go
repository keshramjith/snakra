package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"

	"github.com/keshramjith/snakra/internal/server"
)

func main() {
	env := os.Getenv("ENV")
	if env == "dev" {
		envErr := godotenv.Load(".env")
		if envErr != nil {
			fmt.Println("Failed to find .env file")
			os.Exit(1)
		}
	}

	port := os.Getenv("PORT")
	// env := os.Getenv("ENV")
	s3bucketName := os.Getenv("S3_BUCKET")

	var logger *zap.SugaredLogger
	logger = zap.NewExample().Sugar()

	srv := server.NewServer(logger, s3bucketName, port)
	logger.Infof("Starting server on %s", port)
	err := srv.ListenAndServe()
	defer srv.CloseDb()
	logger.Fatal(err)
}
