package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	requestID := ctx.Value("requestID")
	fmt.Println("requestID:", requestID)
}

func main() {
	ctx := context.WithValue(context.Background(), "requestID", "12345")
	process(ctx)
}
