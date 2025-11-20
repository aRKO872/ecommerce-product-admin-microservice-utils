package utils

import (
	"context"
	"log"
	"time"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/literals"
	"github.com/go-playground/validator/v10"
)

func DummyServiceSetup(servName string) {
	log.Printf("Set up %s service.\n", servName)
}

func GetEnv(cfg interface{}) {
	if err := ParseEnv(cfg); err != nil {
		log.Fatal("Error parsing environment variables: ", err)
	}
	validator := validator.New()
	if err := validator.Struct(cfg); err != nil {
		log.Fatal("Config validation failed: ", err)
	}
}

func GetCorrelationID(ctx context.Context) string {
	if correlationId, ok := ctx.Value(literals.ContextKeyCorrelationID).(string); !ok {
		return ""
	} else {
		return correlationId
	}
}

func GetCurrentTime() string {
	if loc, err := time.LoadLocation("Asia/Kolkata"); err != nil {
		return ""
	} else {
		currentTime := time.Now().UTC()
		currentTimeIndia := currentTime.In(loc)
		return currentTimeIndia.Format(time.RFC3339)
	}
}