package app

import (
	"fmt"
	"github.com/TatarinAlba/WBTest/config"
	v1 "github.com/TatarinAlba/WBTest/internal/controller/http/v1"
	"github.com/TatarinAlba/WBTest/internal/controller/nats"
	"github.com/TatarinAlba/WBTest/internal/usecase"
	"github.com/TatarinAlba/WBTest/internal/usecase/cache"
	"github.com/TatarinAlba/WBTest/internal/usecase/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"sync"

	_ "github.com/lib/pq"
)

func Run(cfg *config.Config) error {
	// Can change to 'DebugLevel'
	logrus.SetLevel(logrus.InfoLevel)
	// For html rendering we are using engine
	engine := html.New("./view", ".html")
	postgresConfigLine := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		cfg.Postgres.Username,
		cfg.Postgres.DatabaseName,
		cfg.Postgres.Password,
	)
	db, err := sqlx.Open("postgres", postgresConfigLine)
	if err != nil {
		return err
	}
	orderRepository := repo.NewOrderRepositoryPostgres(db)
	// Cache initialization
	ordersCache, err := cache.NewCache(orderRepository)
	if err != nil {
		return err
	}
	orderService := usecase.NewOrderUsecase(orderRepository, ordersCache)
	orderController := v1.NewOrderController(orderService)
	// Nats client connection
	orderSubscriber := nats.NewNatsController(orderService)
	group := sync.WaitGroup{}
	group.Add(1)
	go orderSubscriber.ConnectToServer(cfg)
	// Connection to the http controller
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	group.Add(1)
	orderController.Route(app)
	go func() {
		serverConfig := fmt.Sprintf(":%s", cfg.Server.Port)
		err := app.Listen(serverConfig)
		if err != nil {
			logrus.Fatal(err)
		}
	}()
	logrus.Info("Successfully started both servers (nats, fiber) client!")
	group.Wait()
	return nil
}
