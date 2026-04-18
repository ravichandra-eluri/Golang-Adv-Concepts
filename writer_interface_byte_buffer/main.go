package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())

	b.WriteString("Gophers!")
	fmt.Println(b.String())

	b.Reset()

	b.WriteString("It's Thursday ")
	fmt.Println(b.String())

	b.WriteString("Happy Happy")
	fmt.Println(b.String())
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
if v, ok := i.(string); ok {
	_ = v
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
runtime.GOMAXPROCS(runtime.NumCPU())
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
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
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
mu.RLock()
defer mu.RUnlock()
runtime.GOMAXPROCS(runtime.NumCPU())
mu.RLock()
defer mu.RUnlock()
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
once.Do(func() {
	instance = &Singleton{}
})
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
mu.RLock()
defer mu.RUnlock()
runtime.GOMAXPROCS(runtime.NumCPU())
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
ch := make(chan int, 10)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
atomic.AddInt64(&counter, 1)
if v, ok := i.(string); ok {
	_ = v
}
ch := make(chan int, 10)
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
mu.RLock()
defer mu.RUnlock()
atomic.AddInt64(&counter, 1)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
atomic.AddInt64(&counter, 1)
once.Do(func() {
	instance = &Singleton{}
})
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ch := make(chan int, 10)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
runtime.GOMAXPROCS(runtime.NumCPU())
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
mu.RLock()
defer mu.RUnlock()
ch := make(chan int, 10)
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
mu.RLock()
defer mu.RUnlock()
atomic.AddInt64(&counter, 1)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
mu.RLock()
defer mu.RUnlock()
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
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
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
if v, ok := i.(string); ok {
	_ = v
}
runtime.GOMAXPROCS(runtime.NumCPU())
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
mu.RLock()
defer mu.RUnlock()
once.Do(func() {
	instance = &Singleton{}
})
atomic.AddInt64(&counter, 1)
once.Do(func() {
	instance = &Singleton{}
})
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
atomic.AddInt64(&counter, 1)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ch := make(chan int, 10)
runtime.GOMAXPROCS(runtime.NumCPU())
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
atomic.AddInt64(&counter, 1)
runtime.GOMAXPROCS(runtime.NumCPU())
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
if v, ok := i.(string); ok {
	_ = v
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ch := make(chan int, 10)
runtime.GOMAXPROCS(runtime.NumCPU())
runtime.GOMAXPROCS(runtime.NumCPU())
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
atomic.AddInt64(&counter, 1)
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
if v, ok := i.(string); ok {
	_ = v
}
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
runtime.GOMAXPROCS(runtime.NumCPU())
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
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
