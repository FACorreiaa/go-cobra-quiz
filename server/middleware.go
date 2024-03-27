package server

import (
	"log"
	"net/http"
	"time"
)

type wrappedWritter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWritter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &wrappedWritter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(ww, r)
		log.Println(
			ww.statusCode,
			r.Method,
			r.URL.Path,
			time.Since(start))
	})
}
