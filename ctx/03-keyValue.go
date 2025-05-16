package ctx

import (
	"context"
	"fmt"
)

func KeyValue() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", "12345")
	bh(ctx)

}

func bh(ctx context.Context) {
	token := ctx.Value("id")
	fmt.Println("token", token)
}
