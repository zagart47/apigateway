package router

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

const RequestIdKey = "request_id"

func (r *Router) CheckRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestId := req.URL.Query().Get(RequestIdKey)
		if len(requestId) < 6 {
			requestId = uuid.NewString()
		}
		req.Header.Set(RequestIdKey, requestId)
		next.ServeHTTP(w, req)
	})
}

func (r *Router) RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		type responseWriter struct {
			http.ResponseWriter
			statusCode int
		}
		requestId := req.Header.Get("request_id")
		remoteAddr := req.RemoteAddr
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rw, req)
		fmt.Printf("%s %s %d %s\n", time.Now().Format("2006-01-02 15:04:05"), remoteAddr, rw.statusCode, requestId)
	})
}

func (r *Router) SetHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, req)
	})
}
