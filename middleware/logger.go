package middleware

import (
	"log/slog"
	"net/http"
)

func RequestResponseLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &ResponseWriterWrapper{w, http.StatusOK}
		slog.Default().Info("received request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(wrappedWriter, r)
		slog.Default().Info("sent response", "status", wrappedWriter.StatusCode)
	})
}

type ResponseWriterWrapper struct {
	http.ResponseWriter
	StatusCode int
}

func (w *ResponseWriterWrapper) WriteHeader(statusCode int) {
	if statusCode != http.StatusOK {
		w.StatusCode = statusCode
		w.ResponseWriter.WriteHeader(statusCode)
	}
}
