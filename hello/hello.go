package main

import "fmt"

func main() {
	fmt.Println(Hello("Frankie"))
}

func Hello(name string) string {
	return "Hello, " + name
}
