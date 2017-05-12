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
	fmt.Println("Commita 90% til alt annet, men glemte å commite til denne fila. Så da må jeg gjøre det.")
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	a = ":) øæå TEST"
	fmt.Println(a)
	fmt.Println("Test 1 2 3, 3 2 1")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("532 + 18 = ", add(532, 18))
	
	sum := 0
	for index := 0; index < 10; index++ {
		sum += 1
	}
	fmt.Println(sum)
}



