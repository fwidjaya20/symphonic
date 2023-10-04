package event

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQDriver struct {
	ContractEvent.DriverArgs

	channel    *amqp091.Channel
	connection *amqp091.Connection
}

func NewRabbitMQDriver(args ContractEvent.DriverArgs) ContractEvent.QueueDriver {
	addr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/",
		args.Config.GetString("queue.connections.rabbitmq.protocol"),
		args.Config.GetString("queue.connections.rabbitmq.username"),
		args.Config.GetString("queue.connections.rabbitmq.password"),
		args.Config.GetString("queue.connections.rabbitmq.host"),
		args.Config.GetString("queue.connections.rabbitmq.port"),
	)

	connection, err := amqp091.Dial(addr)
	if nil != err {
		args.Logger.Errorf("RabbitMQ connection failed. %v\n", err.Error())
	}

	channel, err := connection.Channel()
	if nil != err {
		args.Logger.Errorf("RabbitMQ channel failed. %v\n", err.Error())
	}

	err = channel.ExchangeDeclare(
		args.Job.Signature(),
		amqp091.ExchangeFanout,
		true,
		false,
		false,
		false,
		nil,
	)
	if nil != err {
		args.Logger.Errorf("failed to declare '%s' exchange. %v", err.Error())
	}

	return &RabbitMQDriver{
		DriverArgs: args,
		channel:    channel,
		connection: connection,
	}
}

func (r *RabbitMQDriver) Driver() string {
	return DriverRabbitMQ
}

func (r *RabbitMQDriver) Publish() error {
	var err error

	payload, err := json.Marshal(r.Job.GetPayload())
	if nil != err {
		r.Logger.Errorf("failed to marshal Job payload. %v", err.Error())
		return err
	}

	err = r.channel.PublishWithContext(
		context.Background(),
		r.Job.Signature(),
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        payload,
		},
	)
	if nil != err {
		r.Logger.Errorf("failed to publish '%s'. %v", r.Job.Signature(), err.Error())
		return err
	}

	return nil
}

func (r *RabbitMQDriver) Subscribe(c context.Context) error {
	queue, err := r.channel.QueueDeclare(
		r.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if nil != err {
		r.Logger.Errorf("failed to declare '%s' queue. %v", r.QueueName, err.Error())
		return err
	}

	err = r.channel.QueueBind(
		queue.Name,
		"",
		r.Job.Signature(),
		false,
		nil,
	)
	if nil != err {
		r.Logger.Errorf("failed to bind '%s' to '%s' queue. %v", r.Job.Signature(), r.QueueName, err.Error())
		return err
	}

	messages, err := r.channel.ConsumeWithContext(
		c,
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if nil != err {
		r.Logger.Errorf("failed to consume messages on '%s' queue. %v", r.QueueName, err.Error())
	}

	var forever chan struct{}

	go func() {
		for msg := range messages {
			jobInstance := reflect.New(reflect.TypeOf(r.Job)).Interface().(ContractEvent.Job)

			if err = json.Unmarshal([]byte(msg.Body), jobInstance); nil != err {
				r.Logger.Infof("error unmarshalling payload: %v\n", err.Error())
				continue
			}

			for _, listener := range r.Listeners {
				if err = listener.Handle(jobInstance); nil != err {
					r.Logger.Errorf("error calling Handle method: %v\n", err.Error())
				}
			}
		}
	}()

	<-forever

	return nil
}

func (r *RabbitMQDriver) Flush() error {
	if err := r.channel.Close(); nil != err {
		r.Logger.Errorf("failed to close channel. %v", err.Error())
		return err
	}
	if err := r.connection.Close(); nil != err {
		r.Logger.Errorf("failed to close connection. %v", err.Error())
		return err
	}
	return nil
}
