package main

import (
	"context"
	"fmt"
	"sync"
)

// Need a key type.
type myKey int

// Need a key value.
const key myKey = 0

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			ctx := context.WithValue(ctx, key, id)

			<-ctx.Done()
			fmt.Println("cancelled: ", id)
		}(i)
	}

	cancel()
	wg.Wait()

}
