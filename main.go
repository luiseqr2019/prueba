package main

import (
	"fmt"
	"prueba/stack"
)

func main() {
	st := stack.NewStack[int]()

	st.Push(10)
	st.Push(20)
	fmt.Println(st.Size())
	fmt.Println(st.Top())
	value, _ := st.Pop()
	fmt.Println(value)
	fmt.Println(st.Size())
	fmt.Println(st.Top())

}
