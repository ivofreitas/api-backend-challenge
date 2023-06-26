package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/sword/api-backend-challenge/config"
	"github.com/sword/api-backend-challenge/log"
)

type Publisher struct {
	producer     sarama.AsyncProducer
	kafkaAddress []string
	kafkaTopic   string
	logger       *logrus.Entry
}

func NewPublisher() *Publisher {
	config := config.GetEnv().Kafka

	s := &Publisher{}
	s.kafkaTopic = config.Topic
	s.kafkaAddress = []string{config.Broker}
	s.logger = log.NewEntry()

	s.connect()

	go func() {
		for {
			select {
			case success := <-s.producer.Successes():
				s.logger.Printf("Produced message to topic %s, partition %d, offset %d\n",
					success.Topic, success.Partition, success.Offset)
			case err := <-s.producer.Errors():
				s.logger.Printf("Failed to produce message: %s\n", err.Err)
			}
		}
	}()

	return s
}

func (s *Publisher) Close() {
	s.producer.AsyncClose()
}

func (s *Publisher) Publish(str string) {
	message := &sarama.ProducerMessage{
		Topic: s.kafkaTopic,
		Value: sarama.StringEncoder(str),
	}

	s.producer.Input() <- message
}

func (s *Publisher) connect() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_1_0_0
	producer, err := sarama.NewAsyncProducer(s.kafkaAddress, config)
	if err != nil {
		s.logger.WithError(err).Fatal()
	}

	s.producer = producer
}
