package web

import (
	"go-rest-template/internal/app/utils"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w}
		handler.ServeHTTP(sw, r)
		end := time.Now()
		responseTime := end.Sub(start)
		utils.Print(r.RemoteAddr, r.Method, r.URL, responseTime, sw.status)
	})
}
