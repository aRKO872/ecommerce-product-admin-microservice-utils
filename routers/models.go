package routers

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type ServiceRouter struct {
	AppID       string
	Middlewares []mux.MiddlewareFunc
	Routes      []CommonRouter
	wg          sync.WaitGroup
	Shutdown    func() error
}

type CommonRouter struct {
	Method      string
	Endpoint    string
	Middlewares []mux.MiddlewareFunc
	Handler     http.HandlerFunc
}

type httpConfig struct {
	Port string `env:"PORT" validate:"required,numeric"`
	Host string `env:"HOST" validate:"required"`
}

type grpcConfig struct {
	Port string `env:"PORT" validate:"required,numeric"`
	Host string `env:"HOST" validate:"required"`
	ReqRespSize string `env:"GRPC_REQ_RESP_SIZE"`
}
