package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main(){

	a := "123"
	a_sc := &a //копируем его ссылку
	*a_sc = "456" //вставляем новое значение по ссылки в памяти

	//либо так

	var b string= "789"
	var b_sc = &b
	*b_sc = "1011"

	fmt.Println(a)
	fmt.Println(b)

	//структуры

	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{})
	fmt.Println(Vertex{X:1})

	//ссылки на структуры
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	
}
