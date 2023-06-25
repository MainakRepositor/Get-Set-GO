package main

import "fmt"

var age int = 22
var name string = "Mainak"

func main(){
	// Print
	fmt.Print("Hello ") // same line
	fmt.Print("Mainak \n") // new line return
	fmt.Print("A new line")

	// Print Line
	fmt.Println("Hello")
	fmt.Println("New Line")

	fmt.Println("My age is",age)

	// Formatted String
	fmt.Printf("My age is %v and my name is %v \n",age,name)
	fmt.Printf("My age is %v and my name is %q \n",age,name)
	fmt.Printf("Data type of age is %T and name is %T\n",age,name)
	fmt.Printf("Your score is %0.1f points\n",25.55)

	//Sprintf or Save Printf
	var str = fmt.Sprintf("My age is %v and my name is %v \n",age,name)
	fmt.Println("The saved string is :",str)
}
