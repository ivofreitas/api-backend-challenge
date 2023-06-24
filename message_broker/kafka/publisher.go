package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/sword/api-backend-challenge/config"
	"sync/atomic"
	"time"
)

type Publisher struct {
	producer     sarama.SyncProducer
	id           string
	running      int32
	kafkaAddress []string
	kafkaTopic   string
}

func NewPublisher() (*Publisher, error) {
	config := config.GetEnv().Kafka

	s := &Publisher{}
	atomic.StoreInt32(&s.running, 0)
	s.kafkaTopic = config.Topic
	s.kafkaAddress = config.Brokers

	var err error
	for i := 0; i < 60; i++ {
		err = s.connect()
		if err != nil {
			//logger.Get().Errorf("Failed to start consumer: %s", err)
		} else {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Publisher) IsRunning() bool {
	return atomic.LoadInt32(&s.running) != 0
}

func (s *Publisher) Close() {
	err := s.producer.Close()
	if err != nil {
		//logger.Get().Errorln(err)
	}
	atomic.StoreInt32(&s.running, 0)
}

func (s *Publisher) Publish(str string) error {
	var err error
	if !s.IsRunning() {
		return errors.New("unable to publish to kafka brokers - producer is not running")
	}
	message := &sarama.ProducerMessage{
		Topic: s.kafkaTopic,
		Value: sarama.StringEncoder(str),
	}
	_, _, err = s.producer.SendMessage(message)
	//logger.Get().Debugf("Publish to %s/%d/%d topic the message %s", s.kafkaTopic, p, o, str)
	return err
}

func (s *Publisher) connect() (err error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Version = sarama.V2_1_0_0
	s.producer, err = sarama.NewSyncProducer(s.kafkaAddress, config)
	if err != nil {
		//err = errors.Wrapf(err, "Unable to connect to kafka brokers - %s. Error - %s.", s.kafkaAddress, err)
	}
	atomic.StoreInt32(&s.running, 1)

	return
}
