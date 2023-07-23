package main

import "fmt"

func main()  {
	fmt.Println("Hello World")

	// string
	var nameOne string = "Alex"
	var nameTwo = "Alice"
	var nameThree string
	nameThree = "Bob"
	nameFour := "Cat"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// int
	var numOne int = 20
	var numTwo = 30
	numThree :=  25
	fmt.Println(numOne, numTwo, numThree)

	// check bits & memory part (int8, int32...)

	// different ways of printing functions (Format specifier)
	fmt.Println("My name is ", nameOne, "\n and my age is ", numOne)
	fmt.Printf("My name is %v and my age is %v \n", nameOne, numOne)
	fmt.Printf("My name is %q \n", nameOne)

	// save formatted strings
	var str = fmt.Sprintf("My name is %v and my age is %v ", nameOne, numOne)
	fmt.Println(str)

	// arrays
	var num = [3]int {20, 30, 40}
	fmt.Println(num, len(num))
	name := [2]string {"Test", "Lax"}
	name[0] = "Alice"
	fmt.Println(name)

	// slice 
	var number = []int {20, 30}
	number = append(number, 40)
	fmt.Println(number)

	// slice ranges
	rangeOne := num[1:3]
	fmt.Println(rangeOne)

}