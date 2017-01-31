package main

import "fmt"
import "math/rand"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Code below this line:")
	fmt.Println("---------------------")
	fmt.Println("Hello, World")
	
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	a = ":) øæå TEST"
	fmt.Println(a)
	fmt.Println("My favorite number is", rand.Intn(10))
}
