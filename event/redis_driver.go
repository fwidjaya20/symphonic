package event

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/redis/go-redis/v9"
)

type RedisDriver struct {
	ContractEvent.DriverArgs
}

func NewRedisDriver(args ContractEvent.DriverArgs) ContractEvent.QueueDriver {
	return &RedisDriver{
		args,
	}
}

func (r *RedisDriver) Driver() string {
	return DriverRedis
}

func (r *RedisDriver) Publish() error {
	client := r.getConnection()

	payload, err := json.Marshal(r.Job.GetPayload())
	if nil != err {
		return err
	}

	if err = client.Publish(context.Background(), r.Job.Signature(), payload).Err(); nil != err {
		return err
	}

	r.Logger.Infof("'%s' has been published: %v", r.Job.Signature(), r.Job.GetPayload())

	return nil
}

func (r *RedisDriver) Subscribe(c context.Context) error {
	client := r.getConnection()

	stream := client.Subscribe(context.Background(), r.Job.Signature())
	defer stream.Close()

	typeOfJob := reflect.TypeOf(r.Job)

	for {
		msg, err := stream.ReceiveMessage(context.Background())
		if nil != err {
			fmt.Printf("Error receiving message: %v\n", err)
			continue
		}

		jobInstance := reflect.New(typeOfJob).Interface()

		if err = json.Unmarshal([]byte(msg.Payload), jobInstance); nil != err {
			r.Logger.Infof("Error unmarshalling payload: %v\n", err.Error())
			continue
		}

		for _, listener := range r.Listeners {
			handleMethod := reflect.ValueOf(listener).MethodByName("Handle")

			if !handleMethod.IsValid() {
				r.Logger.Error("Handle method not found on listener")
				continue
			}

			if handleMethod.Type().NumIn() != 1 {
				r.Logger.Error("Handle method has an unexpected number of parameters\n")
				continue
			}

			result := handleMethod.Call([]reflect.Value{reflect.ValueOf(jobInstance)})

			if nil != result[0].Interface() {
				r.Logger.Errorf("Error calling Handle method: %v\n", result[0].Interface())
			}
		}
	}
}

func (r *RedisDriver) getConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%s",
			r.Config.GetString("queue.connections.redis.host"),
			r.Config.GetString("queue.connections.redis.port"),
		),
	})
}
