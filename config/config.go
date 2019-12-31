package config

type Config struct {
	OrdersConfig       OrdersConfig
	RedisReserveConfig RedisConfig
	RedisOrdersConfig  RedisConfig
}

type OrdersConfig struct {
	Host string
	Port int
}

type RedisConfig struct {
	Channel, URI, Password string
}
