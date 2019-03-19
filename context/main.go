package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	contextWithCancel()
	contextWithTimeout()
}

func produceMessageAfter(ctx context.Context, msg string, duration time.Duration) {
	select {
	case <-time.After(duration):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func contextWithCancel() {
	ctx := context.Background()
	ctxC, cancel := context.WithCancel(ctx)
	time.AfterFunc(time.Second, cancel)
	produceMessageAfter(ctxC, "context with cancel", 5000*time.Millisecond)
}

func contextWithTimeout() {
	ctx := context.Background()
	ctxC, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	produceMessageAfter(ctxC, "context with timeout", 5000*time.Millisecond)
}
