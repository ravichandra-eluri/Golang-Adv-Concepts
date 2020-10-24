package main

import "fmt"

// Defers run in LIFO order after the function returns
func deferOrder() {
	fmt.Println("--- defer order ---")
	for i := 1; i <= 3; i++ {
		defer fmt.Println("deferred:", i) // prints 3, 2, 1
	}
	fmt.Println("function body done")
}

// recover() catches a panic; the deferred function runs before unwinding continues
func safeDiv(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, nil
}

// defer runs even on an early return — useful for cleanup
func openResource(name string) error {
	fmt.Println("opening:", name)
	defer fmt.Println("closing:", name) // always runs
	if name == "" {
		return fmt.Errorf("empty name")
	}
	fmt.Println("using:", name)
	return nil
}

// panic propagates up the stack until a recover() catches it
func mustPositive(n int) int {
	if n <= 0 {
		panic(fmt.Sprintf("expected positive, got %d", n))
	}
	return n
}

func main() {
	deferOrder()

	fmt.Println()
	r1, err := safeDiv(10, 2)
	fmt.Printf("10/2 = %d, err = %v\n", r1, err)

	r2, err := safeDiv(10, 0)
	fmt.Printf("10/0 = %d, err = %v\n", r2, err)

	fmt.Println()
	openResource("db-conn")
	fmt.Println()
	openResource("")

	fmt.Println()
	fmt.Println(mustPositive(5))

	// Panic caught from a goroutine's deferred recover
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("goroutine recovered:", r)
			}
		}()
		mustPositive(-1)
	}()
	<-done
}
