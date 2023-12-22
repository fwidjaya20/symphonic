package main

import (
	SysLog "log"
	"time"

	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	ContractFoundation "github.com/fwidjaya20/symphonic/contracts/foundation"
	"github.com/fwidjaya20/symphonic/event"
	ExampleEvent "github.com/fwidjaya20/symphonic/example/event"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/fwidjaya20/symphonic/log"
)

func main() {
	facades.Config().Add("app.providers", []ContractFoundation.ServiceProvider{
		&event.ServiceProvider{},
		&log.ServiceProvider{},
	})

	facades.Config().Add("database", map[string]any{
		"connections": map[string]any{
			"redis": map[string]any{
				"driver":   "redis",
				"host":     facades.Config().Get("REDIS_HOST", "localhost"),
				"port":     facades.Config().Get("REDIS_PORT", 6379),
				"database": facades.Config().Get("REDIS_DATABASE", "0"),
				"password": facades.Config().Get("REDIS_PASSWORD", ""),
			},
		},
	})

	facades.Config().Add("queue", map[string]any{
		"connections": map[string]any{
			"kafka": map[string]any{},
			"rabbitmq": map[string]any{
				"protocol": "amqp",
				"username": "guest",
				"password": "guest",
				"host":     "localhost",
				"port":     "5672",
			},
			"redis": map[string]any{
				"host":     facades.Config().Get("database.connections.redis.host"),
				"port":     facades.Config().Get("database.connections.redis.port"),
				"database": facades.Config().Get("database.connections.redis.database"),
				"password": facades.Config().Get("database.connections.redis.password"),
			},
		},
		"default": "sync",
	})

	facades.App().Boot()

	kernel := ExampleEvent.Kernel{}

	facades.Event().Register(kernel.Listen())

	go func() {
		if err := facades.Event().Run(ContractEvent.RunEvent{
			Connection:    event.DriverRedis,
			ConsumerGroup: "symphonic-example-group",
			Job:           ExampleEvent.PostCreated{},
		}); nil != err {
			SysLog.Fatalln(err.Error())
		}
	}()

	time.Sleep(5 * time.Second)

	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			if err := facades.Event().Job(&ExampleEvent.PostCreated{
				Id:        int64(i),
				Author:    "Fredrick Widjaya",
				CreatedAt: time.Now(),
			}).OnConnection(event.DriverRedis).Publish(); nil != err {
				SysLog.Fatalln(err.Error())
			}
		}
	}()

	select {}
}
