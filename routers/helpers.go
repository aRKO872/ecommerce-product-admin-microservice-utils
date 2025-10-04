package routers

import (
	"log"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
	"github.com/go-playground/validator/v10"
)

func getEnv(cfg interface{}) {
	if err := utils.ParseEnv(&cfg); err != nil {
		log.Fatal("Error parsing environment variables: ", err)
	}
	validator := validator.New()
	if err := validator.Struct(cfg); err != nil {
		log.Fatal("Config validation failed: ", err)
	}
}

func (s *ServiceRouter) Wait() {
	s.wg.Wait()
}