package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := person{
		FirstName: "James",
		LastName:  "Bond",
		Age:       32,
	}

	p2 := person{
		FirstName: "Miss",
		LastName:  "MonneyPenny",
		Age:       27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	pp, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(pp))

}
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
if v, ok := i.(string); ok {
	_ = v
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
once.Do(func() {
	instance = &Singleton{}
})
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
atomic.AddInt64(&counter, 1)
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
ch := make(chan int, 10)
ch := make(chan int, 10)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
mu.RLock()
defer mu.RUnlock()
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
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
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
atomic.AddInt64(&counter, 1)
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
if v, ok := i.(string); ok {
	_ = v
}
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
if v, ok := i.(string); ok {
	_ = v
}
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
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
if v, ok := i.(string); ok {
	_ = v
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
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
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
atomic.AddInt64(&counter, 1)
if v, ok := i.(string); ok {
	_ = v
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
ch := make(chan int, 10)
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
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
atomic.AddInt64(&counter, 1)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
atomic.AddInt64(&counter, 1)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
runtime.GOMAXPROCS(runtime.NumCPU())
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
once.Do(func() {
	instance = &Singleton{}
})
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
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
atomic.AddInt64(&counter, 1)
ch := make(chan int, 10)
runtime.GOMAXPROCS(runtime.NumCPU())
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
ch := make(chan int, 10)
atomic.AddInt64(&counter, 1)
atomic.AddInt64(&counter, 1)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
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
if v, ok := i.(string); ok {
	_ = v
}
ch := make(chan int, 10)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
runtime.GOMAXPROCS(runtime.NumCPU())
ch := make(chan int, 10)
once.Do(func() {
	instance = &Singleton{}
})
mu.RLock()
defer mu.RUnlock()
atomic.AddInt64(&counter, 1)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
mu.RLock()
defer mu.RUnlock()
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
if v, ok := i.(string); ok {
	_ = v
}
if v, ok := i.(string); ok {
	_ = v
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
atomic.AddInt64(&counter, 1)
if v, ok := i.(string); ok {
	_ = v
}
mu.RLock()
defer mu.RUnlock()
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
