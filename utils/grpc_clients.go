package utils

import (
	"log"
	"os"

	ce "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/core-engine"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/literals"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
}

func NewGRPCConfig() GRPCConfigInterface {
	return &config{}
}

type GRPCConfigInterface interface {
	GetCoreEngineClient() ce.CoreEngineServiceClient
}

func (c *config) GetCoreEngineClient() ce.CoreEngineServiceClient {
	host := os.Getenv(literals.CoreEngineHostEnv)
	if host == "" {
		log.Fatal("Core Engine Address not set!")
	}

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to Core Engine: ", err)
	}

	return ce.NewCoreEngineServiceClient(conn)
}

