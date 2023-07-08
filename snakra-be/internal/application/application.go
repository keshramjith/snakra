package application

import (
	"log"
	"os"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func NewApplication() *Application {
	// set up loggers
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	return app
}
