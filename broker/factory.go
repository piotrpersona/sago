package broker

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/piotrpersona/saga/config"
	"github.com/piotrpersona/saga/store"
)

func New(broker string, config config.Config) (b Broker, err error) {
	switch broker {
	case "redis":
		var client *redis.Client
		client, err = store.NewRedisClient(config.RedisOrdersConfig)
		b = redisBroker{client: client, channel: config.RedisReserveConfig.Channel}
	default:
		err = fmt.Errorf("Broker %s not supported", broker)
	}
	return
}
