package ctx

import (
	"context"
	"fmt"
	"time"
)

func Intro() {
	// Contexto é um pacote que permite o cancelamento de uma operação
	// que está sendo executada em uma goroutine.
	// É um pacote que permite o cancelamento de uma operação que está sendo executada em uma goroutine.

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel: cancelado")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel: Hotel reservado")
	}
}
