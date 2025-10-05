package utils

import (
	"log"
	"os"

	ce "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/core-engine"
	invMsc "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/inventory-msc"
	ordersMsc "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/orders-msc"
	productsMsc "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/products-msc"
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
	GetInventoryMscClient() invMsc.InventoryServiceClient
	GetOrdersMscClient() ordersMsc.OrdersServiceClient
	GetProductsMscClient() productsMsc.ProductsServiceClient
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

func (c *config) GetInventoryMscClient() invMsc.InventoryServiceClient {
	host := os.Getenv(literals.InventoryMscHostEnv)
	if host == "" {
		log.Fatal("Inventory Microservice Address not set!")
	}

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to Inventory Microservice: ", err)
	}

	return invMsc.NewInventoryServiceClient(conn)
}

func (c *config) GetOrdersMscClient() ordersMsc.OrdersServiceClient {
	host := os.Getenv(literals.OrdersMscHostEnv)
	if host == "" {
		log.Fatal("Orders Microservice Address not set!")
	}

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to Orders Microservice: ", err)
	}

	return ordersMsc.NewOrdersServiceClient(conn)
}

func (c *config) GetProductsMscClient()	productsMsc.ProductsServiceClient {
	host := os.Getenv(literals.ProductsMscHostEnv)
	if host == "" {
		log.Fatal("Products Microservice Address not set!")
	}

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to Products Microservice: ", err)
	}

	return productsMsc.NewProductsServiceClient(conn)
}