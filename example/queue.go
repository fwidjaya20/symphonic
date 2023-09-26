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
			"kafka":    map[string]any{},
			"rabbitmq": map[string]any{},
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
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			if err := facades.Event().Job(&ExampleEvent.PostCreated{
				Id:        int64(i),
				Author:    "Fredrick Widjaya",
				CreatedAt: time.Now(),
			}).OnConnection("redis").Publish(); nil != err {
				SysLog.Fatalln(err.Error())
			}
		}
	}()

	go func() {
		if err := facades.Event().Run(ContractEvent.RunEvent{
			Connection: event.DriverRedis,
			Job:        ExampleEvent.PostCreated{},
		}); nil != err {
			SysLog.Fatalln(err.Error())
		}
	}()

	//go func() {
	//	client := redis.NewClient(&redis.Options{
	//		Addr: "localhost:6379", // Update with your Redis server address
	//	})
	//
	//	pubsub := client.Subscribe(context.Background(), ExampleEvent.PostCreated{}.Signature())
	//	defer pubsub.Close()
	//
	//	for {
	//		msg, err := pubsub.ReceiveMessage(context.Background())
	//		if err != nil {
	//			fmt.Printf("Error receiving message: %v\n", err)
	//			continue
	//		}
	//
	//		// Process the message here
	//		fmt.Printf("Received: %s\n", msg.Payload)
	//	}
	//}()

	select {}
}
