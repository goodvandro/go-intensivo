package main

import (
	"database/sql"
	"fmt"

	"github.com/goodvandro/go-intensivo/internal/infra/database"
	usecase "github.com/goodvandro/go-intensivo/internal/useCase"
	_ "github.com/mattn/go-sqlite3" // the underline is because the package is not used directly but when it is necessary
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	defer db.Close() // wait all processes to execute than close the database

	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "1234",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)

	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
