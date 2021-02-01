package main

import "fmt"

func main() {
	//fmt.Println(isPalindrome2("  #$%^45675676%&*  "))
	minStack := Constructor()
	minStack.Push(0)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	fmt.Println("----------------------")

	minStack.Push(-3)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	fmt.Println("----------------------")

	minStack.Push(-2)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	fmt.Println("----------------------")

	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	fmt.Println("----------------------")

	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	fmt.Println("----------------------")

	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	fmt.Println("----------------------")
}
