package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/JSainsburyPLC/issasdk/httplog"
	"github.com/JSainsburyPLC/issasdk/logging"

	"github.com/gordonrehling2/certavs/service/handlers"
)

func Start(ctx context.Context, port int, logger httplog.Logger) {
	addr := ":" + fmt.Sprintf("%d", port)
	logging.Info(ctx, "Starting certavs service")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/healthcheck", handlers.Healthcheck)
	//router.HandleFunc("/rfe", handlers.RfeList)

	//router.Handle("GET", "/rfe/:productID", handlers.RfeProductList())
	//router.Handle("POST", "/rfe/:productID", handlers.RfeProductCreate())

	http.ListenAndServe(addr, router)
	logging.Fatal(ctx, "Server failed")
}
