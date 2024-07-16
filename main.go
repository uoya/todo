package main

import (
	"fmt"
)

func main() {
	name := "John"
	fmt.Println(Hello(name))
}

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
