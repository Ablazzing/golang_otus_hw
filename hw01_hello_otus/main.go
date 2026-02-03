package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse" //nolint:all
)

func main() {
	fmt.Println(reverse.String("Hello, OTUS!"))
}
