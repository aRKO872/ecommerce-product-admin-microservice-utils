package routers

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type ConsumerEventFunc func(ev *kafka.Message) error

type ServiceRouter struct {
	AppID          string
	Middlewares    []mux.MiddlewareFunc
	Routes         []CommonRouter
	wg             sync.WaitGroup
	Shutdown       func() error
	PubsubProducer *KafkaProducer
	PubsubConsumer *kafka.Consumer
	IsKafkaEnabled bool
	pubsubDone     chan struct{}
}

type CommonRouter struct {
	Method      string
	Endpoint    string
	Middlewares []mux.MiddlewareFunc
	Handler     http.HandlerFunc
}

type httpConfig struct {
	Port              string `env:"PORT" validate:"required,numeric"`
	Host              string `env:"HOST" validate:"required"`
	KafkaFlushTimeout int    `env:"KAFKA_FLUSH_TIMEOUT_MS" validate:"numeric"`
}

type grpcConfig struct {
	Port              string `env:"PORT" validate:"required,numeric"`
	Host              string `env:"HOST" validate:"required"`
	ReqRespSize       string `env:"GRPC_REQ_RESP_SIZE"`
	KafkaFlushTimeout int    `env:"KAFKA_FLUSH_TIMEOUT_MS" validate:"numeric"`
}

type KafkaConfig struct {
	Brokers       string `env:"KAFKA_BROKERS" validate:"required"`
	ConsumerGroup string `env:"KAFKA_CONSUMER_GROUP"`
	Topics        string `env:"KAFKA_TOPICS"`
	Timeout       int    `env:"KAFKA_TIMEOUT_MS"`
}
