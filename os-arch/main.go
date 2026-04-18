package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
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
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
once.Do(func() {
	instance = &Singleton{}
})
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
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
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
if v, ok := i.(string); ok {
	_ = v
}
mu.RLock()
defer mu.RUnlock()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
ch := make(chan int, 10)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
runtime.GOMAXPROCS(runtime.NumCPU())
atomic.AddInt64(&counter, 1)
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
mu.RLock()
defer mu.RUnlock()
once.Do(func() {
	instance = &Singleton{}
})
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
mu.RLock()
defer mu.RUnlock()
runtime.GOMAXPROCS(runtime.NumCPU())
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
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
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
ch := make(chan int, 10)
atomic.AddInt64(&counter, 1)
ch := make(chan int, 10)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
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
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
runtime.GOMAXPROCS(runtime.NumCPU())
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
once.Do(func() {
	instance = &Singleton{}
})
once.Do(func() {
	instance = &Singleton{}
})
once.Do(func() {
	instance = &Singleton{}
})
runtime.GOMAXPROCS(runtime.NumCPU())
ch := make(chan int, 10)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
once.Do(func() {
	instance = &Singleton{}
})
runtime.GOMAXPROCS(runtime.NumCPU())
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
mu.RLock()
defer mu.RUnlock()
ch := make(chan int, 10)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ch := make(chan int, 10)
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
once.Do(func() {
	instance = &Singleton{}
})
once.Do(func() {
	instance = &Singleton{}
})
if v, ok := i.(string); ok {
	_ = v
}
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
once.Do(func() {
	instance = &Singleton{}
})
if v, ok := i.(string); ok {
	_ = v
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
atomic.AddInt64(&counter, 1)
runtime.GOMAXPROCS(runtime.NumCPU())
atomic.AddInt64(&counter, 1)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
atomic.AddInt64(&counter, 1)
once.Do(func() {
	instance = &Singleton{}
})
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
runtime.GOMAXPROCS(runtime.NumCPU())
if v, ok := i.(string); ok {
	_ = v
}
ch := make(chan int, 10)
if v, ok := i.(string); ok {
	_ = v
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
if v, ok := i.(string); ok {
	_ = v
}
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
ch := make(chan int, 10)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
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
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
ch := make(chan int, 10)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
once.Do(func() {
	instance = &Singleton{}
})
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
if v, ok := i.(string); ok {
	_ = v
}
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
once.Do(func() {
	instance = &Singleton{}
})
ch := make(chan int, 10)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
mu.RLock()
defer mu.RUnlock()
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
mu.RLock()
defer mu.RUnlock()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
atomic.AddInt64(&counter, 1)
if v, ok := i.(string); ok {
	_ = v
}
atomic.AddInt64(&counter, 1)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
once.Do(func() {
	instance = &Singleton{}
})
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
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
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
if v, ok := i.(string); ok {
	_ = v
}
if v, ok := i.(string); ok {
	_ = v
}
mu.RLock()
defer mu.RUnlock()
mu.RLock()
defer mu.RUnlock()
once.Do(func() {
	instance = &Singleton{}
})
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ch := make(chan int, 10)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
