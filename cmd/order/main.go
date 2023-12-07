package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/goodvandro/go-intensivo/internal/infra/database"
	usecase "github.com/goodvandro/go-intensivo/internal/useCase"
	"github.com/goodvandro/go-intensivo/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3" // the underline is because the package is not used directly but when it is necessary
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	defer db.Close() // wait all processes to execute than close the database

	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel) // escutando a fila // T2
	rabbitmqWorker(msgRabbitmqChannel, uc)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")

	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)

		if err != nil {
			panic(err)
		}

		msg.Ack(false)

		fmt.Println("Mensagem processada e salva no banco", output)
	}
}
