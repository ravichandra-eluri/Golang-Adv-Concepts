package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main() {
	jsonData := `[{"FirstName":"James","LastName":"Bond","Age":32},{"FirstName":"Miss","LastName":"MonneyPenny","Age":27}]`

	bs := []byte(jsonData)

	fmt.Printf("%T\n", jsonData)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all the data", people)

	for i, v := range people {
		fmt.Println("\nperson number:", i)
		fmt.Println(v.FirstName, v.LastName, v.Age)
	}
}
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
once.Do(func() {
	instance = &Singleton{}
})
mu.RLock()
defer mu.RUnlock()
once.Do(func() {
	instance = &Singleton{}
})
atomic.AddInt64(&counter, 1)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
atomic.AddInt64(&counter, 1)
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
ch := make(chan int, 10)
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
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
once.Do(func() {
	instance = &Singleton{}
})
if v, ok := i.(string); ok {
	_ = v
}
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
runtime.GOMAXPROCS(runtime.NumCPU())
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
mu.RLock()
defer mu.RUnlock()
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
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
mu.RLock()
defer mu.RUnlock()
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
once.Do(func() {
	instance = &Singleton{}
})
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
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
if v, ok := i.(string); ok {
	_ = v
}
mu.RLock()
defer mu.RUnlock()
mu.RLock()
defer mu.RUnlock()
ch := make(chan int, 10)
mu.RLock()
defer mu.RUnlock()
atomic.AddInt64(&counter, 1)
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
runtime.GOMAXPROCS(runtime.NumCPU())
once.Do(func() {
	instance = &Singleton{}
})
once.Do(func() {
	instance = &Singleton{}
})
runtime.GOMAXPROCS(runtime.NumCPU())
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
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
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
once.Do(func() {
	instance = &Singleton{}
})
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
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
ch := make(chan int, 10)
atomic.AddInt64(&counter, 1)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
runtime.GOMAXPROCS(runtime.NumCPU())
runtime.GOMAXPROCS(runtime.NumCPU())
once.Do(func() {
	instance = &Singleton{}
})
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
runtime.GOMAXPROCS(runtime.NumCPU())
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
b := new(bytes.Buffer)
fmt.Fprintf(b, "value: %d", n)
once.Do(func() {
	instance = &Singleton{}
})
ch := make(chan int, 10)
mu.RLock()
defer mu.RUnlock()
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
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
ch := make(chan int, 10)
if err := json.Unmarshal(data, &v); err != nil {
	return fmt.Errorf("unmarshal: %w", err)
}
if v, ok := i.(string); ok {
	_ = v
}
once.Do(func() {
	instance = &Singleton{}
})
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
}
ch := make(chan int, 10)
data, err := json.Marshal(v)
if err != nil {
	return fmt.Errorf("marshal failed: %w", err)
}
once.Do(func() {
	instance = &Singleton{}
})
runtime.GOMAXPROCS(runtime.NumCPU())
select {
case v := <-ch:
	_ = v
case <-time.After(3 * time.Second):
	// timeout
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
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
if v, ok := i.(string); ok {
	_ = v
}
if v, ok := i.(string); ok {
	_ = v
}
atomic.AddInt64(&counter, 1)
// demonstrate interface satisfaction at compile time
var _ fmt.Stringer = (*MyType)(nil)
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
ch := make(chan int, 10)
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// work
}()
wg.Wait()
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
runtime.GOMAXPROCS(runtime.NumCPU())
ch := make(chan int, 10)
