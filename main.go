package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ordres := generateOrders(20)

	go func() {
		defer wg.Done()
		processOrders(ordres)
	}()

	go func() {
		defer wg.Done()
		updateOrderStatuses(ordres)
	}()

	reportOrderStatus(ordres)

	wg.Wait()

	fmt.Println("All operations completed. Exiting")

}

func updateOrderStatuses(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		status := []string{"Processing", "Shipped", "Delivered"}[rand.Intn(3)]
		order.Status = status
		fmt.Printf("Updated order %d status: %s\n", order.ID, status)
	}
}

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := range count {
		orders[i] = &Order{ID: i + 1, Status: "Pending"}
	}
	return orders
}

func reportOrderStatus(orders []*Order) {
	for range 5 {
		time.Sleep(1 * time.Second)
		fmt.Println("\n--- Order Status Report ---")
		for _, order := range orders {
			fmt.Printf("Order %d: %s\n", order.ID, order.Status)
		}
		fmt.Println("-----------------------")
	}
}
