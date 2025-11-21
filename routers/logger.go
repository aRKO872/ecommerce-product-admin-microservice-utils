package routers

import (
	"context"
	"log"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/literals"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
)

type LogMessage struct {
	CorrelationID string `json:"correlation_id"`
	Message       string `json:"message"`
	AppID         string `json:"app_id"`
	Level         string `json:"level"`
	Timestamp     string `json:"timestamp"`
}

type Logger struct {
	LogProducer *KafkaProducer
	AppID       string
}

func NewLogger(
	logProducer *KafkaProducer,
	appID string,
) *Logger {
	return &Logger{
		LogProducer: logProducer,
		AppID:       appID,
	}
}

func (l *Logger) Log(
	ctx context.Context,
	logMsg string,
	level string,
) {
	timestamp := utils.GetCurrentTime()

	log.Printf(literals.LogTemplate, l.AppID, level, logMsg, timestamp)
	logObj := LogMessage{
		CorrelationID: utils.GetCorrelationID(ctx),
		Message:       logMsg,
		AppID:         l.AppID,
		Level:         level,
		Timestamp:     timestamp,
	}
	if err := l.LogProducer.ProduceAsyncWrite(
		literals.TopicLogMessages,
		literals.TopicLogMessages,
		logObj,
	); err != nil {
		log.Println("error logging message from service to log-service, ", err.Error())
	}
}