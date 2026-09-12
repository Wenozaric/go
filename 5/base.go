package main

import "fmt"

func main() {
	//a n[Type] 
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)


	//flexible array
	var s []int = primes[1:4] //include fisrt, exclude last
	fmt.Println(s)

	//slice doest not store any data
	names := [4]string{"Misha", "Pasha", "Dasha", "Masha"}
	fmt.Println(names)

	a2 := names[0: 2]
	b2 := names[0: 3]
	fmt.Println(a2, b2) //print data in array

	a3 := names[0:2]

	b2[0] = "New value" //modify data in original array
	fmt.Println(a2, b2) //new values
	fmt.Println(names)  //new value in original array

	fmt.Printf("%p\n %p\n", a2, a3) //identical values

	//u can also do that with slices
	p := [5]int{1, 2, 3, 4, 5}
	newP := p[:]//extend all array with slice [:]
	newP = append(newP, 6)
	//if new len > max len in original array, its different arrays, else it's the same array (pointer to same memory)


	//different
	fmt.Println(p) //len 5
	fmt.Println(newP) //len 6

	//another example
	p2 := [3]int{1, 2}
	newP2 := p2[:]
	newP2[2] = 3
	//same arrays
	fmt.Println(p2) //len 3
	fmt.Println(newP2) //len3
	
	//len return len of array cap return capacity of array (max len)

							//Reminder of printf
	//%v - value, %+v for structres, %T - type <= universal flags
	//%d - default int, %x - hex, %b - binary, %f - float <= ints
	//%s - string, %q - string with quotes "", %c - char (Unicode) <= strings
	//%t - bool (true/false), %p - pointer
	//\n - new line, %% - just print "%"

	fmt.Printf("len: %d cap: %d\narray:%v\n", len(p2), cap(p2), p2)

	//nil arrays
	fmt.Println("\n--nil arrays--")
	var s1 []int
	fmt.Println(s, len(s), cap(s))
	if s1 == nil { fmt.Println("nil!")}

	//make()
	fmt.Println("\n--make--")

	dinam_a := make([]int, 5) //len = 5, cap = 5
	printSlice("dinam_a", dinam_a)

	dinam_b := make([]int, 0, 5) //len = 0, cap = 5
	printSlice("dinam_b", dinam_b)

	dinam_c := dinam_b[:2] //len = 2, cap = 5
	printSlice("dinam_c", dinam_c)

	dinam_d := dinam_c[2:5] //len = 3, cap = 3
	printSlice("dinam_d", dinam_d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
