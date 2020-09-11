package main

import "fmt"

func main() {
	fmt.Println("main function")
}

func hello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
