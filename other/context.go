// trying context
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	// thirdPartFunc()
	ctx := context.Background()
	timeout := flag.Duration("t", 200, "timeout in milliseconds")
	flag.Parse()
	fmt.Println(*timeout)
	res, err := businessLogic(ctx, *timeout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func businessLogic(ctx context.Context, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	ch := make(chan struct{})
	go func(ch chan<- struct{}) {
		thirdPartFunc(ch)
	}(ch)

	select {
	case <-ctx.Done():
		return "", fmt.Errorf("took more than %s: %v", timeout, ctx.Err())
	case <-ch:
		return "done", nil
	}
}

func thirdPartFunc(ch chan<- struct{}) {
	start := time.Now().UnixMilli()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("took", time.Now().UnixMilli()-start, "ms")
	ch <- struct{}{}
}
