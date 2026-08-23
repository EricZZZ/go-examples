package main

import (
	"database/sql"
	"log"
	"sql-in-go/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbConnection, err := sql.Open("sqlite3", "./shop.db")

	if err != nil {
		log.Fatal(err)
	}

	defer dbConnection.Close()

	orderRepository := &database.OrderRespository{Db: dbConnection}

	err = orderRepository.CreateTable()

	if err != nil {
		log.Fatal("Error creating orders table:", err)
	}

	err = orderRepository.Insert(database.Order{Product: "MacBook", Amount: 10000})
	if err != nil {
		log.Fatal("Error inserting order", err)
	}

	err = orderRepository.Insert(database.Order{Product: "IPhone", Amount: 789})
	if err != nil {
		log.Fatal("Error inserting order", err)
	}

	orders, err := orderRepository.GetAll()
	if err != nil {
		log.Fatal("Error getting orders", err)
	}

	log.Println(orders)

	order, err := orderRepository.GetById(1)
	if err != nil {
		log.Fatal("Error getting order by id", err)
	}

	log.Println("------update order------")
	order.Product = "MacBook Pro"
	order.Amount = 20000
	err = orderRepository.Update(order)
	if err != nil {
		log.Fatal("Error updating order", err)
	}

	orders, err = orderRepository.GetAll()
	if err != nil {
		log.Fatal("Error getting orders", err)
	}

	log.Println(orders)

	log.Println("------delete order------")
	err = orderRepository.Delete(1)
	if err != nil {
		log.Fatal("Error deleting order", err)
	}

	orders, err = orderRepository.GetAll()
	if err != nil {
		log.Fatal("Error getting orders", err)
	}

	log.Println(orders)

}
