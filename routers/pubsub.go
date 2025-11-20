package routers

import (
	"strings"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func (s *ServiceRouter) InitKafkaConsumer(
	consumerEventFunc ConsumerEventFunc,
) (error) {
	var kafkaConfig KafkaConfig
	utils.GetEnv(&kafkaConfig)

	if kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": kafkaConfig.Brokers,
        "security.protocol": "SASL_SSL",
        "sasl.mechanisms":   "PLAIN",
        "group.id":          kafkaConfig.ConsumerGroup,
        "auto.offset.reset": "earliest",
	}); err != nil {
		return err
	} else {
		s.PubsubConsumer = kafkaConsumer
	}

	topicsArr := strings.Split(kafkaConfig.Topics, ",")
	s.pubsubDone = make(chan struct{})
	if err := s.PubsubConsumer.SubscribeTopics(topicsArr, nil); err != nil {
		return err
	}

	go s.ListenEvents(consumerEventFunc, kafkaConfig.Timeout)

	s.IsKafkaEnabled = true
	return nil
}

// if kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
// 		"bootstrap.servers": kafkaConfig.Brokers,
// 		"security.protocol": "SASL_SSL",
// 		"sasl.mechanisms":   "PLAIN",
// 		"acks":              "all",
// 	}); err != nil {
// 		return err
// 	} else {
// 		s.PubsubProducer = kafkaProducer
// 	}