package main

import (
	"fmt"
)
var x int=5
	var y int 


	func main() {
	var t bool= true
	var f bool


	//names := "Thanpat Apiwong"
	
	fmt.Println("t is ", t)
	fmt.Println("f is ", f)
	var age int=20
	var favNum float64 = 1.6180339
	var complex128 complex128=5+5i
	var r rune = 10
	fmt.Println("age is ", age, "favNum is", favNum)
	fmt.Println("complex128 is ", complex128)
	fmt.Println("rune is", r)
	var myName string= "Thanapat Apiwong" 
	fmt.Println(myName + "is a robot")
	fmt.Println(myName[3])
	fmt.Println("Length of myName is", len(myName))
	
	var arry5[5] float64
	arry5[0]=98.7
	arry5[1]=93.2
	arry5[2]=77.2
	arry5[3]=83.7
	arry5[4]=88.2
	fmt.Println(arry5)
	fmt.Println("Length of arry5 is", len(arry5))
	fmt.Println("arry5[3]) is", arry5[3])
	arry3 :=[3] float64 {98, 93, 77}
	fmt.Println(arry3)


	var s [] float64 = arry5[0:5]
	fmt.Println(s)

	type Person struct{
		name string
		age int
	} 
	var p Person
	p.name = "Thanapat"
	p.age=20
	fmt.Println(p)



	var x int=5
	var xPtr*int=&x
	fmt.Println(xPtr)
}
