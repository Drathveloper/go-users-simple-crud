package main

import (
	"log"
	"log/slog"
	"net/http"
	//_ "net/http/pprof"
	"os"
)

func main() {
	initLogger()
	container := initializeDependencies()
	initDatabase(container)
	router := initializeRoutes(container)
	/*go func() {
		slog.Info("starting pprof server on :6060")
		log.Fatal(http.ListenAndServe(":6060", nil))
	}()*/
	slog.Info("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func initLogger() {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})
	slog.SetDefault(slog.New(logHandler))
}
