package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func workSlow(ctx context.Context) error {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Doing work ", i)
		case <-ctx.Done():
			fmt.Println("Cancel with context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("start")
	wg.Add(1)
	go workSlow(ctx)
	wg.Wait()
	fmt.Println("finish")
}
