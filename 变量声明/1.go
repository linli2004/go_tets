package main

import "fmt"
func main(){
	const name string = "Alice"
	fmt.Println(name)

	const age=20
	fmt.Println(age)

	const name_1,name_2 string="Bob","Charlie"
	fmt.Println(name_1,name_2)

	const name_3,age_1="David",30
	fmt.Println(name_3,age_1)

	var age_2 uint8=25
	var age_3=32
	age_4:=40
	fmt.Println(age_2,age_3,age_4)

	var age_5,age_6 int=50,60
	fmt.Println(age_5,age_6)

	var name_4,age_7="jay",30
	fmt.Println(name_4,age_7)

	name_5,age_8:= "tom",40
	fmt.Println(name_5,age_8)
}