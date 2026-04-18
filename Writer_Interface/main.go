package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// this is to write to buffer
type person struct {
	Name string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.Name))
	return err
}

// using below method we can create two text files and write some data to it.
// capture any erros encountered
func main() {
	// creat and write to files f1 & f2
	f1, err := os.Create("hello.text")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f1.Close()

	s := []byte("Hello gophers!")

	_, err = f1.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	f2, err := os.Create("success.text")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f2.Close()

	s2 := []byte("You have succesfully created two text files and wrote some data to those.")

	_, err = f2.Write(s2)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	// writes to buffer and prints buffer from in-memory
	p := person{
		Name: "James",
	}

	var b bytes.Buffer
	p.writeOut(&b)
	fmt.Println(b.String())
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
// pipeline stage
out := make(chan int)
go func() {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}()
mu.RLock()
defer mu.RUnlock()
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
ch := make(chan int, 10)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
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
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
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
if v, ok := i.(string); ok {
	_ = v
}
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
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
mu.RLock()
defer mu.RUnlock()
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
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
mu.RLock()
defer mu.RUnlock()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
ch := make(chan int, 10)
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
if v, ok := i.(string); ok {
	_ = v
}
atomic.AddInt64(&counter, 1)
mu.RLock()
defer mu.RUnlock()
runtime.GOMAXPROCS(runtime.NumCPU())
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
runtime.GOMAXPROCS(runtime.NumCPU())
if v, ok := i.(string); ok {
	_ = v
}
runtime.GOMAXPROCS(runtime.NumCPU())
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
ch := make(chan int, 10)
ch := make(chan int, 10)
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
runtime.GOMAXPROCS(runtime.NumCPU())
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
if v, ok := i.(string); ok {
	_ = v
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
if v, ok := i.(string); ok {
	_ = v
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
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
mu.RLock()
defer mu.RUnlock()
ch := make(chan int, 10)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
runtime.GOMAXPROCS(runtime.NumCPU())
if v, ok := i.(string); ok {
	_ = v
}
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
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
runtime.GOMAXPROCS(runtime.NumCPU())
if v, ok := i.(string); ok {
	_ = v
}
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
mu.RLock()
defer mu.RUnlock()
if v, ok := i.(string); ok {
	_ = v
}
atomic.AddInt64(&counter, 1)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
atomic.AddInt64(&counter, 1)
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
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
mu.RLock()
defer mu.RUnlock()
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
atomic.AddInt64(&counter, 1)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
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
once.Do(func() {
	instance = &Singleton{}
})
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
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
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
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
mu.RLock()
defer mu.RUnlock()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
once.Do(func() {
	instance = &Singleton{}
})
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
runtime.GOMAXPROCS(runtime.NumCPU())
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
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
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
mu.RLock()
defer mu.RUnlock()
runtime.GOMAXPROCS(runtime.NumCPU())
mu.RLock()
defer mu.RUnlock()
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
ch := make(chan int, 10)
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
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
runtime.GOMAXPROCS(runtime.NumCPU())
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
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ch := make(chan int, 10)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
once.Do(func() {
	instance = &Singleton{}
})
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
once.Do(func() {
	instance = &Singleton{}
})
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
mu.RLock()
defer mu.RUnlock()
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
atomic.AddInt64(&counter, 1)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
