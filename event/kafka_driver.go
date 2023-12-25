package event

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/sirupsen/logrus"
)

type KafkaDriver struct {
	ContractEvent.DriverArgs

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
	if nil != err {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Errorf("Unable to marshal %v", d.Job.GetPayload())
		return err
	}

	if err := d.publisher.Publish(d.Job.Topic(), &message.Message{
		UUID:    watermill.NewUUID(),
		Payload: payload,
	}); err != nil {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Errorf("Unable to publish kafka message")
		return err
	}

	return nil
}

func (d *KafkaDriver) Subscribe(ctx context.Context) error {
	if err := d.provideSubscriber(); err != nil {
		return err
	}

	messages, err := d.subscriber.Subscribe(ctx, d.Job.Topic())
	if nil != err {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Errorf("failed to consume messages with '%s' topic", d.Job.Topic())
		return err
	}

	var forever chan struct{}

	go func(packets <-chan *message.Message) {
		for it := range packets {
			jobInstance := reflect.New(reflect.TypeOf(d.Job)).Interface().(ContractEvent.Job)

			if err = json.Unmarshal(it.Payload, jobInstance); nil != err {
				d.Logger.WithFields(logrus.Fields{
					logrus.ErrorKey: err,
					"payload":       string(it.Payload),
				}).Warn("error unmarshalling payload")
				continue
			}

			var wg sync.WaitGroup

			for _, listener := range d.Listeners {
				wg.Add(1)
				go func(fn ContractEvent.Listener) {
					defer wg.Done()

					if err := fn.Handle(jobInstance); err != nil {
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
		kafka.PublisherConfig{
			Brokers:   []string{d.addr},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if nil != err {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Error("Unable to initialize Kafka Publisher.")
		return err
	}
	d.publisher = publisher

	return nil
}

func (d *KafkaDriver) provideSubscriber() error {
	d.Logger.WithFields(logrus.Fields{
		"Offset":           d.InitialOffset,
		"To Sarama Offset": d.InitialOffset.SaramaOffset(),
		"Sarama Offset":    sarama.OffsetOldest,
	}).Info("FIND THIS")

	if d.subscriber != nil {
		return nil
	}

	config := kafka.DefaultSaramaSubscriberConfig()
	config.Consumer.Offsets.Initial = d.InitialOffset.SaramaOffset()

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{d.addr},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: config,
			ConsumerGroup:         d.ConsumerGroup,
		},
		watermill.NewStdLogger(false, false),
	)
	if nil != err {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Error("Unable to initialize Kafka Subscriber.")
		return err
	}

	d.subscriber = subscriber

	return nil
}

func NewKafkaDriver(args ContractEvent.DriverArgs) ContractEvent.QueueDriver {
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
