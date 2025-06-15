package config

import (
	"time"

	"github.com/caarlos0/env"
)

// Represents a structure with all env variables needed by the backend.
var cfg struct {
	DatabaseURL                string        `env:"DATABASE_URL"`
	DatabaseUser               string        `env:"DATABASE_USER"`
	DatabasePassword           string        `env:"DATABASE_PASSWORD"`
	DatabaseHost               string        `env:"DATABASE_HOST"`
	DatabasePort               int           `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseName               string        `env:"DATABASE_NAME"`
	DatabaseSSLMode            string        `env:"DATABASE_SSLMODE" envDefault:"require"`
	DatabaseMaxOpenConnections int           `env:"DATABASE_MAX_OPEN_CONNECTIONS" envDefault:"10"`
	DatabaseMaxIdleConnections int           `env:"DATABASE_MAX_IDLE_CONNECTIONS" envDefault:"2"`
	StatsdHost                 string        `env:"STATSD_HOST"`
	StatsdPort                 int           `env:"STATSD_PORT"`
	EnvName                    string        `env:"ENV_NAME,required"`
	Port                       int           `env:"PORT" envDefault:"8080"`
	MatchDisplayLowDelta       time.Duration `env:"MATCH_DISPLAY_LOW_DELTA" envDefault:"12h"`
	MatchDisplayHighDelta      time.Duration `env:"MATCH_DISPLAY_LOW_DELTA" envDefault:"96h"`
	MinCashOutAmount           float64       `env:"MIN_CASH_OUT_AMOUNT" envDefault:"15"`
	AmqpGamesExchange          string        `env:"AMQP_GAMES_EXCHANGE,required"`
	MaxPicks                   int           `env:"MAX_PICKS" envDefault:"4"`
	NameChangeInterval         time.Duration `env:"NAME_CHANGE_INTERVAL" envDefault:"720h"`
	RMQHost                    string        `env:"RMQ_HOST,required"`
	RMQPort                    int           `env:"RMQ_PORT,required"`
	RMQVHost                   string        `env:"RMQ_VHOST,required"`
	RMQUser                    string        `env:"RMQ_USER,required"`
	RMQPassword                string        `env:"RMQ_PASSWORD,required"`
	RMQGameUpdatesExchange     string        `env:"RMQ_GAME_UPDATES_EXCHANGE,required"`
	RMQFCMExchange             string        `env:"RMQ_FCM_EXCHANGE,required"`
	StripeSigningSecret        string        `env:"STRIPE_SIGNING_SECRET"`
	StripeSecret               string        `env:"STRIPE_SECRET"`
	RedisAddress               string        `env:"REDIS_ADDRESS,required"`
	RedisPassword              string        `env:"REDIS_PASSWORD,required"`
	RedisDb                    int           `env:"REDIS_DB,required"`
	DefaultChatRoomChannel     string        `env:"DEFAULT_CHAT_ROOM_CHANNEL" envDefault:"laliga"`
}

func init() {
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
}
