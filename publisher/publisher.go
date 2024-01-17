package main

import (
	"context"
	"encoding/json"
	"github.com/TatarinAlba/WBTest/publisher/entityJSON"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"os"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	ctx := context.Background()
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}
	_, err = js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
		MaxBytes: 1024 * 1024,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		myStruct := entityJSON.Order{}
		err = gofakeit.Struct(&myStruct)
		// Can be changed by yourself
		myStruct.Delivery.Email = "some_site@gmail.com"
		myStruct.Delivery.Phone = "+74444444444"
		if err != nil {
			log.Fatal(err)
		}
		toWrite, err := json.Marshal(myStruct)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile("./mock_data.json", toWrite, 0666)
		if err != nil {
			log.Fatal(err)
		}
		jsonData, err := os.ReadFile("./mock_data.json")
		if err != nil {
			log.Fatal(err)
		}
		if _, err := js.Publish(ctx, "ORDERS.ADD", jsonData); err != nil {
			log.Fatal(err)
		}
		log.Print("Sent the message successfully!")
		// Can be changed by yourself
		time.Sleep(1 * time.Second)
	}
}
