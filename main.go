package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	initLogger()
	container := initializeDependencies()
	initDatabase(container)
	router := initializeRoutes(container)
	slog.Info("Server started")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func initLogger() {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})
	slog.SetDefault(slog.New(logHandler))
}
