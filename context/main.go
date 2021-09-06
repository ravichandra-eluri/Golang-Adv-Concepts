package main

import (
	"context"
	"fmt"
	"time"
)

type ctxKey string

func doWork(ctx context.Context, name string) {
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Printf("%s: finished\n", name)
	case <-ctx.Done():
		fmt.Printf("%s: cancelled — %v\n", name, ctx.Err())
	}
}

func fetchUserID(ctx context.Context) string {
	if id, ok := ctx.Value(ctxKey("userID")).(string); ok {
		return id
	}
	return "unknown"
}

func main() {
	// WithCancel: cancel manually
	ctx1, cancel1 := context.WithCancel(context.Background())
	go doWork(ctx1, "task-cancel")
	time.Sleep(50 * time.Millisecond)
	cancel1() // fires before task finishes
	time.Sleep(10 * time.Millisecond)

	// WithTimeout: auto-cancel after duration
	ctx2, cancel2 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel2()
	doWork(ctx2, "task-timeout") // 200ms work > 100ms timeout → cancelled

	// WithDeadline: cancel at a specific wall-clock time
	deadline := time.Now().Add(150 * time.Millisecond)
	ctx3, cancel3 := context.WithDeadline(context.Background(), deadline)
	defer cancel3()
	doWork(ctx3, "task-deadline")

	// WithValue: pass request-scoped data down the call stack
	ctx4 := context.WithValue(context.Background(), ctxKey("userID"), "user-42")
	fmt.Println("user id from context:", fetchUserID(ctx4))

	// Propagation: cancelling a parent cancels all children
	parent, cancelParent := context.WithCancel(context.Background())
	child, cancelChild := context.WithCancel(parent)
	defer cancelChild()

	go doWork(child, "child-task")
	time.Sleep(30 * time.Millisecond)
	cancelParent() // cancels child too
	time.Sleep(10 * time.Millisecond)
}
