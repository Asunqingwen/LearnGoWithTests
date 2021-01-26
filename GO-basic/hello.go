package main

import "fmt"

//常量
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
