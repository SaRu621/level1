package main

import (
	"context"
	"fmt"
	"sync"
)

// Первый способ: использование канала

func myGoroutine(done chan bool) {
	done <- true
}

// Второй способ: контекст

func myGoroutine2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("Выходим из горутины")
		return
	default:

	}
}

// Третий способ: внешний флаг (механизм похож на контекст)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	/*done := make(chan bool)

	go myGoroutine(done)

	<-done*/

	ctx, cancel := context.WithCancel(context.Background())
	go myGoroutine2(ctx, &wg)

	cancel()

	wg.Wait()
	fmt.Println("Высвободили горутину")
}
