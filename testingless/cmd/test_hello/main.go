package main

import "fmt"

const helloStr = "Hello, "

func Hello(name string) string {
	return helloStr + name
}

func main() {
	fmt.Println(Hello("Mark"))
}
