package event

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/sirupsen/logrus"
)

type RabbitMQDriver struct {
	ContractEvent.DriverArgs

	addr       string
	publisher  *amqp.Publisher
	subscriber *amqp.Subscriber
}

func (d *RabbitMQDriver) Driver() string {
	return DriverRabbitMQ
}

func (d *RabbitMQDriver) Flush() error {
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

func (d *RabbitMQDriver) Publish() error {
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
		}).Errorf("Unable to publish rabbitmq message")
		return err
	}

	return nil
}

func (d *RabbitMQDriver) Subscribe(ctx context.Context) error {
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

func (d *RabbitMQDriver) providePublisher() error {
	if d.publisher != nil {
		return nil
	}

	config := amqp.NewDurableQueueConfig(d.addr)

	publisher, err := amqp.NewPublisher(config, watermill.NewStdLogger(false, false))
	if nil != err {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Error("Unable to initialize Kafka Publisher.")
		return err
	}
	d.publisher = publisher

	return nil
}

func (d *RabbitMQDriver) provideSubscriber() error {
	if d.subscriber != nil {
		return nil
	}

	config := amqp.NewDurableQueueConfig(d.addr)

	subscriber, err := amqp.NewSubscriber(config, watermill.NewStdLogger(false, false))
	if nil != err {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Error("Unable to initialize Kafka Subscriber.")
		return err
	}

	d.subscriber = subscriber

	return nil
}

func NewRabbitMQDriver(args ContractEvent.DriverArgs) ContractEvent.QueueDriver {
	return &RabbitMQDriver{
		DriverArgs: args,
		addr: fmt.Sprintf(
			"%s://%s:%s@%s:%s/",
			args.Config.GetString("queue.connections.rabbitmq.protocol"),
			args.Config.GetString("queue.connections.rabbitmq.username"),
			args.Config.GetString("queue.connections.rabbitmq.password"),
			args.Config.GetString("queue.connections.rabbitmq.host"),
			args.Config.GetString("queue.connections.rabbitmq.port"),
		),
		publisher:  nil,
		subscriber: nil,
	}
}
