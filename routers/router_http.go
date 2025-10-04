package routers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/gorilla/mux"
)

func (s *ServiceRouter) ServeHTTP() {
	var cfg httpConfig
	getEnv(&cfg)

	router := mux.NewRouter()

	sRouter := router.PathPrefix("/" + s.AppID).Subrouter()
	for _, cr := range s.Routes {
		r := sRouter.NewRoute().Subrouter()
		r.Use(cr.Middlewares...)
		r.Use(s.Middlewares...)
		r.HandleFunc(cr.Endpoint, cr.Handler).Methods(cr.Method)
	}

	srv := s.serveHTTP(cfg, router)
	s.Shutdown = func() error {
		return srv.Shutdown(context.Background())
	}

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		<-stop
		log.Printf("Shutting down %s server\n", s.AppID)
		if err := s.Shutdown(); err != nil {
			log.Printf("Error shutting down %s server: %v\n", s.AppID, err)
		}
	}()
	
	log.Printf("%s server is running on %s\n", s.AppID, srv.Addr)
	s.wg.Wait()
}

func (s *ServiceRouter) serveHTTP(cfg httpConfig, handler *mux.Router) *http.Server {
	addr := cfg.Host + ":" + cfg.Port
	httpServ := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := httpServ.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("HTTP server ListenAndServe error: ", err)
		}
	}()

	return httpServ
}