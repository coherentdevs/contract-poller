package cache

type Config struct {
	Host     string `env:"REDIS_HOST" envDefault:"localhost:6379"`
	Username string `env:"REDIS_USERNAME" envDefault:""`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}
