package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ultrafenrir/go-devops/internal/pkg/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	r := chi.NewRouter()
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Post("/update/{metricType}/{metricName}/{metricValue}", WriteMetrics(ctx))

	// http.HandleFunc("/", handler)
	// err := http.ListenAndServe("localhost:8080", nil)
	fmt.Println("Serving on ")
	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Fatal(err)
	}
	cancel()

}
