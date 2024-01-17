package nats

import (
	"context"
	"encoding/json"
	"github.com/TatarinAlba/WBTest/config"
	"github.com/TatarinAlba/WBTest/internal/entity"
	"github.com/TatarinAlba/WBTest/internal/usecase"
	"github.com/TatarinAlba/WBTest/pkg/validator"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
	"time"
)

type Controller struct {
	*usecase.OrderUsecase
}

func NewNatsController(orderUsecase *usecase.OrderUsecase) *Controller {
	return &Controller{orderUsecase}
}

func (controller *Controller) ConnectToServer(cfg *config.Config) {
	natsClient, err := nats.Connect(cfg.NatsServer.Url)
	if err != nil {
		logrus.Fatal(err)
	}
	defer natsClient.Close()
	ctx := context.Background()
	jetStream, err := jetstream.New(natsClient)
	if err != nil {
		logrus.Fatal(err)
	}
	stream, err := jetStream.Stream(ctx, "ORDERS")
	if err != nil {
		logrus.Fatalf("Error during connection to the stream in nats from the client: %s", err)
	}
	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:    "NEW",
		Durable: "NEW",
	})
	if err != nil {
		logrus.Fatalf("Error during creating consumer in nats from the client: %s", err)
	}
	for {
		if _, err := consumer.Consume(controller.addOrder); err != nil {
			logrus.Fatal(err)
		}
		// Can be changed by yourself
		time.Sleep(1 * time.Second)
	}
}

func (controller *Controller) addOrder(msg jetstream.Msg) {
	var order entity.Order
	err := msg.Ack()
	if err != nil {
		logrus.Fatalf("Error during message processing: %s", err)
	}
	err = json.Unmarshal(msg.Data(), &order)
	if err != nil {
		logrus.Errorf("Cannot parse given json: %s", err)
		return

	}
	err = validator.ValidateEntity(&order)
	if err != nil {
		logrus.Errorf("Cannot take this entity: %s", err)
		return
	}
	err = controller.OrderUsecase.CreateOrder(order)
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Infof("Added order with uid [%s]", order.OrderUID)
	}
}
