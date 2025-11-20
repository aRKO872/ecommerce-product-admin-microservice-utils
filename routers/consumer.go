package routers

import (
	"log"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// timeout is -1for indefinite wait
func (s *ServiceRouter) ListenEvents(consumerEventFunc func(ev *kafka.Message) error, timeout int) {
	defer s.PubsubConsumer.Close()

	run := true
	for run {
		select {
		case <-s.pubsubDone: 
			log.Println("Shutting down Kafka consumer listener")
			run = false
		default:
			ev, err := s.PubsubConsumer.ReadMessage(time.Duration(timeout) * time.Millisecond); if err != nil {
				log.Println("error reading consume event message: ", err)
			} else {
				if err := consumerEventFunc(ev); err != nil {
					log.Println("error processing consumer event: ", err)
				}
			}
		}
	}

}