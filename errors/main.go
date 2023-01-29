package main

import (
	"errors"
	"fmt"
)

// Sentinel error — compared with errors.Is
var ErrNotFound = errors.New("not found")

// Custom error type — extracted with errors.As
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: field %q — %s", e.Field, e.Message)
}

func findUser(id int) error {
	if id <= 0 {
		return &ValidationError{Field: "id", Message: "must be positive"}
	}
	if id > 100 {
		return fmt.Errorf("findUser: %w", ErrNotFound)
	}
	return nil
}

// Wrapping adds caller context while preserving the original error
func getUser(id int) error {
	if err := findUser(id); err != nil {
		return fmt.Errorf("getUser(%d): %w", id, err)
	}
	return nil
}

func main() {
	// errors.Is: unwraps the entire chain looking for a sentinel
	err := getUser(200)
	fmt.Println(err)
	fmt.Println("is ErrNotFound:", errors.Is(err, ErrNotFound))

	// errors.As: unwraps to find a concrete type
	err2 := getUser(-1)
	var ve *ValidationError
	if errors.As(err2, &ve) {
		fmt.Printf("field %q failed: %s\n", ve.Field, ve.Message)
	}

	// Multi-level wrapping still traces back to root
	base := errors.New("base error")
	layer1 := fmt.Errorf("layer 1: %w", base)
	layer2 := fmt.Errorf("layer 2: %w", layer1)
	fmt.Println(layer2)
	fmt.Println("unwraps to base:", errors.Is(layer2, base))

	// errors.Unwrap: manually peel one layer
	fmt.Println("unwrap layer2:", errors.Unwrap(layer2))
}
