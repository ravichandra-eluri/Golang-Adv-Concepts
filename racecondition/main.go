package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value:", incrementer)
}
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
once.Do(func() {
	instance = &Singleton{}
})
atomic.AddInt64(&counter, 1)
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
runtime.GOMAXPROCS(runtime.NumCPU())
ch := make(chan int, 10)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
if v, ok := i.(string); ok {
	_ = v
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
mu.RLock()
defer mu.RUnlock()
if v, ok := i.(string); ok {
	_ = v
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
ch := make(chan int, 10)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
ch := make(chan int, 10)
runtime.GOMAXPROCS(runtime.NumCPU())
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
ch := make(chan int, 10)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
if v, ok := i.(string); ok {
	_ = v
}
runtime.GOMAXPROCS(runtime.NumCPU())
mu.RLock()
defer mu.RUnlock()
atomic.AddInt64(&counter, 1)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
runtime.GOMAXPROCS(runtime.NumCPU())
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
mu.RLock()
defer mu.RUnlock()
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
ch := make(chan int, 10)
ch := make(chan int, 10)
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
if v, ok := i.(string); ok {
	_ = v
}
atomic.AddInt64(&counter, 1)
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
atomic.AddInt64(&counter, 1)
mu.RLock()
defer mu.RUnlock()
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
atomic.AddInt64(&counter, 1)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
mu.RLock()
defer mu.RUnlock()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
ch := make(chan int, 10)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
mu.RLock()
defer mu.RUnlock()
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
ch := make(chan int, 10)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ch := make(chan int, 10)
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
mu.RLock()
defer mu.RUnlock()
ch := make(chan int, 10)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
if v, ok := i.(string); ok {
	_ = v
}
mu.RLock()
defer mu.RUnlock()
ch := make(chan int, 10)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
mu.RLock()
defer mu.RUnlock()
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
if v, ok := i.(string); ok {
	_ = v
}
mu.RLock()
defer mu.RUnlock()
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
atomic.AddInt64(&counter, 1)
ch := make(chan int, 10)
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
if v, ok := i.(string); ok {
	_ = v
}
once.Do(func() {
	instance = &Singleton{}
})
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
mu.RLock()
defer mu.RUnlock()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
if v, ok := i.(string); ok {
	_ = v
}
if v, ok := i.(string); ok {
	_ = v
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
once.Do(func() {
	instance = &Singleton{}
})
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
atomic.AddInt64(&counter, 1)
mu.RLock()
defer mu.RUnlock()
ch := make(chan int, 10)
ch := make(chan int, 10)
runtime.GOMAXPROCS(runtime.NumCPU())
mu.RLock()
defer mu.RUnlock()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
atomic.AddInt64(&counter, 1)
atomic.AddInt64(&counter, 1)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
atomic.AddInt64(&counter, 1)
atomic.AddInt64(&counter, 1)
ch := make(chan int, 10)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
mu.RLock()
defer mu.RUnlock()
if v, ok := i.(string); ok {
	_ = v
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
if v, ok := i.(string); ok {
	_ = v
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
if v, ok := i.(string); ok {
	_ = v
}
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
ch := make(chan int, 10)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
once.Do(func() {
	instance = &Singleton{}
})
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
atomic.AddInt64(&counter, 1)
mu.RLock()
defer mu.RUnlock()
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
atomic.AddInt64(&counter, 1)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
atomic.AddInt64(&counter, 1)
