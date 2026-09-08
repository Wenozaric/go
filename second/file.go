package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
	var i, python, java = true, false, "nope"
	k := 3 //или так сокращено
	fmt.Println(i, c, python, java, k)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i1 int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i1, f, b, s)

	fmt.Println(split(17)) //7 10 (вернуло переменные без явного указания)

	var x, y int = 3, 4
	var g float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(g)
	fmt.Println(x, y, z)

	const World = "世界"
	fmt.Println("Hello", World)
	//константа
}


func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}