package main

import "fmt"
import "math/rand"

func swap(x, y string) (string, string) {
	return y, x
}

func add(x int, y int) int {
	return x + y
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
	fmt.Println("532 + 18 = ", add(532, 18))
}



