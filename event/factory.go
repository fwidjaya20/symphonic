package event

import "github.com/fwidjaya20/symphonic/contracts/event"

func GetQueueDriver(driver string, args event.DriverArgs) event.QueueDriver {
	switch driver {
	case DriverKafka:
		return NewKafkaDriver(args)
	case DriverRabbitMQ:
		return NewRabbitMQDriver(args)
	case DriverRedis:
		return NewRedisDriver(args)
	default:
		return NewSyncDriver(args)
	}
}
