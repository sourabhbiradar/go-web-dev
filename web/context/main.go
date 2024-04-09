package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// context
	fmt.Println("\tContext")

	// context WithTimeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(2 * time.Second) // simulating main() tasks

	fmt.Println("Cancelling contexting...")
	cancel()

	// wait for 'go worker' to finish (although it should be cancelled by now)
	time.Sleep(1 * time.Second)

	// context withValue
	getUserIDOrderID(321)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled")
		default:
			fmt.Println("worker still active...")
			time.Sleep(1 * time.Second)
		}
	}
}

// context withValue
func order(ctx context.Context, orderID int) {
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		fmt.Println("userID not found in context..")
		return
	}
	fmt.Printf("order %d for user %d\n", orderID, userID)
}
func getUserIDOrderID(orderID int) {
	ctx := context.WithValue(context.Background(), "userID", 123)

	order(ctx, orderID)
}
