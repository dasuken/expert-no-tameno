package main

import (
	"errors"
	"fmt"
)

func main() {
	e1 := errors.New("e1")
	e2 := errors.New("e1")
	fmt.Println(errors.Is(e1, e2))
}

// server
