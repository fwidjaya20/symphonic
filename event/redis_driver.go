package event

import (
	"context"
	"encoding/json"
	"fmt"

	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/redis/go-redis/v9"
)

type RedisDriver struct {
	*ContractEvent.DriverArgs

	connection *redis.Client
}

func (r *RedisDriver) Driver() string {
	return DriverRedis
}

func (r *RedisDriver) Flush() error {
	return r.connection.Close()
}

func (r *RedisDriver) Publish() error {
	payload, err := json.Marshal(r.Job.GetPayload())
	if err != nil {
		return err
	}

	if err := r.connection.Publish(context.Background(), r.Job.Topic(), payload).Err(); err != nil {
		return err
	}

	r.Logger.Infof("'%s' has been published: %v", r.Job.Topic(), r.Job.GetPayload())

	return nil
}

func (r *RedisDriver) Subscribe(c context.Context) error {
	stream := r.connection.Subscribe(c, r.Topic)
	defer stream.Close()

	for {
		msg, err := stream.ReceiveMessage(c)
		if err != nil {
			fmt.Printf("Error receiving message: %v\n", err)
			continue
		}

		for _, listener := range r.Listeners {
			if err = listener.Handle([]byte(msg.Payload)); err != nil {
				r.Logger.Errorf("Error calling Handle method: %v\n", err.Error())
			}
		}
	}
}

func NewRedisDriver(args *ContractEvent.DriverArgs) ContractEvent.QueueDriver {
	return &RedisDriver{
		DriverArgs: args,
		connection: redis.NewClient(
			&redis.Options{ //nolint:exhaustruct // ignore due to redis configuration
				Addr: fmt.Sprintf(
					"%s:%s",
					args.Config.GetString("queue.connections.redis.host"),
					args.Config.GetString("queue.connections.redis.port"),
				),
			},
		),
	}
}
