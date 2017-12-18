package main

import (
	"fmt"
)

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Start.")
		f(s)
		fmt.Println("End.")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

func main() {
	decorator(Hello)("hello world")
}
