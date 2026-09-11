package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("Hello, Go! The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	fmt.Println(math.Pi)
	
	fmt.Println(add(42, 13))
	addStr("Hello", 42)

	a, b := swap("hello", "world")
	fmt.Println(a, b)

}

// или x int y int ( если одного типа можно написать один раз)
func add(x, y int) int { 
	return x + y
}

func addStr(x string, y int) {
	fmt.Printf("%s %d \n", x, y)
}

func swap(x, y string) (string, string){
	return y, x
}