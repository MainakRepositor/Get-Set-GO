package main

import "fmt"

func main(){
	var name string = "Mainak" // single line declaration with data type
	var name2 = "Lucifer" // single line declaration without data type
	var name3 string // mentioning data type and 
	name3 = "Rocket" // assigning afterwards
	name4 := "Cat" // initializing without specifying data type. However := can be used only once while declaration
	
	fmt.Println(name, name2, name3, name4) // print the data
	fmt.Printf("Variable is of type %T \n",name2) // print the type of data



	var number1 int = 21	
	var number2 = 30
	var number3 int
	number3 = 40
	number4 := 50
	fmt.Println(number1, number2, number3, number4)

	var num1 int8 = -128 //only till -128 to 127
	var num2 int8 = 127
	var num3 uint = 250000 // only positive number
	var num4 uint8 = 255 // from 0 - 255
	fmt.Println(num1,num2,num3,num4)

	var score1 float32 = -25.98
	var score2 float32 = 99899878787744784854447888878887784445.4444448888889999997777
	score3 := 1578.64 // by default of type float64
	fmt.Println(score1,score2,score3)


}
