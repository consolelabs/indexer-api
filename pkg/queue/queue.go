package queue

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/queue/message"
)

type Queue struct {
	consumer *kafka.Consumer
	producer *kafka.Producer
	cfg      *config.Config
}

type Config struct {
	Producer bool
	Consumer bool
}

func New(cfg *config.Config, kafkaCfg *Config) *Queue {
	q := &Queue{
		cfg: cfg,
	}
	if kafkaCfg.Consumer {
		instanceId, err := os.Hostname()
		if err != nil {
			logger.L.Fatalf(err, "Failed to get hostname")
		}

		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers":  cfg.Kafka.Servers,
			"group.id":           cfg.Kafka.ConsumerGroup,
			"auto.offset.reset":  "earliest",
			"group.instance.id":  instanceId,
			"session.timeout.ms": 300000,
		})
		if err != nil {
			logger.L.Fatalf(err, "failed to create kafka consumer")
		}
		q.consumer = consumer
	}
	if kafkaCfg.Producer {
		producer, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers":       cfg.Kafka.Servers,
			"go.produce.channel.size": 1000000,
		})
		if err != nil {
			logger.L.Fatalf(err, "failed to create kafka producer")
		}
		q.producer = producer
	}

	return q
}

func (q *Queue) Consumer() *kafka.Consumer {
	return q.consumer
}

func (q *Queue) Producer() *kafka.Producer {
	return q.producer
}

func (q *Queue) ProducerReport() error {
	term := false
	for !term {
		select {
		case e := <-q.producer.Events():
			switch ev := e.(type) {
			case *kafka.Message:
				// Message delivery report
				m := ev
				if m.TopicPartition.Error != nil {
					logger.L.Infof("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					logger.L.Infof("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}

			case kafka.Error:
				// Generic client instance-level errors, such as
				// broker connection failures, authentication issues, etc.
				//
				// These errors should generally be considered informational
				// as the underlying client will automatically try to
				// recover from any errors encountered, the application
				// does not need to take action on them.

				e := ev
				if e.IsFatal() {
					// Fatal error handling.
					//
					// When a fatal error is detected by the producer
					// instance, it will emit kafka.Error event (with
					// IsFatal()) set on the Events channel.
					//
					// Note:
					//   After a fatal error has been raised, any
					//   subsequent Produce*() calls will fail with
					//   the original error code.
					logger.L.Infof("FATAL ERROR: %v: terminating\n", e)
				} else {
					logger.L.Infof("Error: %v\n", e)
				}

			default:
				logger.L.Infof("Ignored event: %s\n", ev)
			}
		}
	}
	return nil
}

func (q *Queue) Enqueue(value message.KafkaMessage) error {
	if q.producer == nil {
		return errors.New("producer is not initialized")
	}
	defer q.Producer().Flush(0)

	topic := q.cfg.Kafka.IndexerTopic

	// use json.Marshal to conver KafkaMessage to []byte
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	q.producer.ProduceChannel() <- &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: bytes,
	}

	return nil
}
