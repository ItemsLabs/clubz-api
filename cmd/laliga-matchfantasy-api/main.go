package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/amqp"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/config"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/dbstore"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/event_bus"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/listeners"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/types"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	echoTrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/olahol/melody.v1"
)

func openDB() (*sql.DB, error) {
	logrus.Info("connecting to database")

	// Start a database connection.
	db, err := sql.Open("pgx", config.DatabaseURL())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.DatabaseMaxOpenConnections())
	db.SetMaxIdleConns(config.DatabaseMaxIdleConnections())

	logrus.Info("pinging db connection")
	// Actually test the connection against the database, so we catch
	// problematic connections early.
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for using Swagger with Echo.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	db, err := openDB()
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	defer db.Close()

	b := event_bus.NewDefaultBus()

	// start datadog tracing
	tracer.Start()
	defer tracer.Stop()

	// start event listener
	listeners.StartGameEventsListener(
		context.Background(),
		config.RMQConnectionURL(),
		config.RMQGameUpdatesExchange(),
		amqp.RandomQueueName("laliga-matchfantasy-api"),
		b,
	)

	// start http handler
	r := echo.New()

	r.Use(echoTrace.Middleware(echoTrace.WithServiceName("laliga-matchfantasy-api")))
	m := melody.New()
	chatWS := melody.New()
	leaderboardWS := melody.New()
	wsConfigs := map[string]types.WebsocketConfig{
		"game": {
			Manager:  m,
			EventBus: b,
		},
		"chat": {
			Manager:  chatWS,
			EventBus: b,
		},
		"leaderboard": {
			Manager:  leaderboardWS,
			EventBus: b,
		},
	}
	routes.InitRouter(r, dbstore.New(db), wsConfigs) // m, b, chatWS)
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	logrus.Info("starting laliga-matchfantasy-api service")
	if err := r.Start(fmt.Sprintf(":%d", config.Port())); err != nil {
		logrus.WithError(err).Error("cannot run service")
	}
}
