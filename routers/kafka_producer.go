package routers

import (
	"encoding/json"
	"time"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var KafkaProducerInstance *KafkaProducer

type KafkaProducer struct {
	pubsubProducer *kafka.Producer
}

func NewKafkaProducer() (*KafkaProducer, error) {
	var kafkaConfig KafkaConfig
	if KafkaProducerInstance != nil {
		return KafkaProducerInstance, nil
	} else {
		utils.GetEnv(&kafkaConfig)
		if kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": kafkaConfig.Brokers,
			"security.protocol": "SASL_SSL",
			"sasl.mechanisms":   "PLAIN",
			"acks":              "all", 
		}); err != nil {
			return nil, err
		} else {
			KafkaProducerInstance = &KafkaProducer{
				pubsubProducer: kafkaProducer,
			}
			return KafkaProducerInstance, nil
		}
	}
}

func (s *KafkaProducer) ProduceAsyncWrite(topic, key string, value any) error {
	byteVal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.pubsubProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          byteVal,
		Timestamp:      time.Now(),
		TimestampType:  kafka.TimestampCreateTime,
	}, nil)
}
