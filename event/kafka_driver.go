package event

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/sirupsen/logrus"
)

type KafkaDriver struct {
	*ContractEvent.DriverArgs

	addr       string
	publisher  *kafka.Publisher
	subscriber *kafka.Subscriber
}

func (d *KafkaDriver) Driver() string {
	return DriverKafka
}

func (d *KafkaDriver) Flush() error {
	if d.publisher != nil {
		if err := d.publisher.Close(); err != nil {
			return err
		}
	}

	if d.subscriber != nil {
		if err := d.subscriber.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (d *KafkaDriver) Publish() error {
	if err := d.providePublisher(); err != nil {
		return err
	}

	payload, err := json.Marshal(d.Job.GetPayload())
	if err != nil {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Errorf("Unable to marshal %v", d.Job.GetPayload())

		return err
	}

	if err := d.publisher.Publish(d.Job.Topic(), &message.Message{
		UUID:     watermill.NewUUID(),
		Metadata: nil,
		Payload:  payload,
	}); err != nil {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Errorf("Unable to publish kafka message")

		return err
	}

	return nil
}

//nolint:dupl // ignore duplication because implementation depends on the provider used.
func (d *KafkaDriver) Subscribe(ctx context.Context) error {
	if err := d.provideSubscriber(); err != nil {
		return err
	}

	messages, err := d.subscriber.Subscribe(ctx, d.Topic)
	if err != nil {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Errorf("failed to consume messages with '%s' topic", d.Topic)

		return err
	}

	var forever chan struct{}

	go func(packets <-chan *message.Message) {
		for it := range packets {
			var wg sync.WaitGroup

			for _, listener := range d.Listeners {
				wg.Add(1)

				go func(fn ContractEvent.Listener) {
					defer wg.Done()

					if err := fn.Handle(it.Payload); err != nil {
						d.Logger.WithFields(logrus.Fields{
							logrus.ErrorKey: err,
						}).Warn("error calling Handle method")
					}
				}(listener)
			}

			wg.Wait()

			it.Ack()
		}
	}(messages)

	<-forever

	return nil
}

func (d *KafkaDriver) providePublisher() error {
	if d.publisher != nil {
		return nil
	}

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{ //nolint:exhaustruct // ignore due to kafka publisher configuration
			Brokers:   []string{d.addr},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Error("Unable to initialize Kafka Publisher.")

		return err
	}

	d.publisher = publisher

	return nil
}

func (d *KafkaDriver) provideSubscriber() error {
	if d.subscriber != nil {
		return nil
	}

	config := kafka.DefaultSaramaSubscriberConfig()
	config.Consumer.Offsets.Initial = d.InitialOffset.SaramaOffset()

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{ //nolint:exhaustruct // ignore due to kafka subscriber configuration
			Brokers:               []string{d.addr},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: config,
			ConsumerGroup:         d.ConsumerGroup,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Error("Unable to initialize Kafka Subscriber.")

		return err
	}

	d.subscriber = subscriber

	return nil
}

func NewKafkaDriver(args *ContractEvent.DriverArgs) ContractEvent.QueueDriver {
	return &KafkaDriver{
		DriverArgs: args,
		addr: fmt.Sprintf(
			"%s:%s",
			args.Config.GetString("queue.connections.kafka.host"),
			args.Config.GetString("queue.connections.kafka.port"),
		),
		publisher:  nil,
		subscriber: nil,
	}
}
