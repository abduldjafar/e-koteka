package main

import "fmt"

func someFunc(a ...interface{}) {

	fmt.Println(a[0])
}

func main() {
	someFunc("q", 1)
}
