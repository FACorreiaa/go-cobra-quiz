package server

import (
	"net/http"

	"github.com/FACorreiaa/go-cobra-quiz/api/handler"
)

func Router() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", handler.HelloHandler)
	return r
}
