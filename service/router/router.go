package router

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/julienschmidt/httprouter"

	"github.com/gordonrehling2/certavs/service"
	"github.com/gordonrehling2/certavs/service/handlers"
	"github.com/gordonrehling2/certavs/config"

)

type Router struct {
	Config config.Config
}

func NewRouter(config config.Config) *Router {
	return &Router{
		Config: config,
	}
}

func (r *Router) Run(rfeService service.IRfeService) {
	route := httprouter.New()
	handler := handlers.NewHandler(rfeService)

	route.GET("/healthcheck", handler.HealthCheck())
	route.GET("/rfe", handler.RfeList())
	route.GET("/rfe/:productID", handler.RfeProductList())
	route.POST("/rfe/:productID", handler.RfeProductCreate())

	address := fmt.Sprintf(":%v", r.Config.Api.Port)

	if err := http.ListenAndServe(address, route); err != nil {
		log.Fatalf("Failed to start server %s", err)
	}
}