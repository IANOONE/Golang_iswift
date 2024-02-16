package main

import (
	"fmt"
	"reflect"
)

// type Employee struct {
// 	name     string
// 	age      int
// 	position string
// }

// type Person struct {
// 	name string
// 	age int
// }

// type Hobby struct {
// 	name  string
// }

// this is embedded struct
// type Member struct {
// 	position   string
// 	person     Person
// 	hobby      Hobby
// }

// func main() {

// !! STRUCT
/* A data type which is collection of various kinds properties/fields, and methods */

// ! Giving value to struct
// we should collect struct to variable
// var employee Employee

// 1. initialize style

// A.
// var employee1 Employee
// employee1.name = "Brilian"
// employee1.age = 24
// employee1.position = "Backend"

// fmt.Printf("%+v \n", employee1)

// B.
// var employee2 = Employee{
// 	name:     "Brilian",
// 	age:      24,
// 	position: "Backend",
// }
// fmt.Printf("%+v \n", employee2)

// !! Pointer to Struct
// var employee3 *Employee = &employee2

// before
// fmt.Printf("%+v \n", employee2)
// fmt.Printf("%+v \n", employee3)
// fmt.Println(strings.Repeat("-", 20))

// employee3.name = "Ian"
// when we change employee3 value, the value of employee2 will be following the change values, bcs their shared memory address using pointer

// after
// fmt.Printf("Employee name :%+v \n", employee2.name)
// fmt.Printf("Employee name :%+v \n", employee3.name)

// ! Embedded Struct
//  embed a struct as a field into a struct

// var member1 = Member{}
// member1.position = "FrontEnd"
// member1.person.name = "Sebastian"
// member1.person.age = 30
// member1.hobby.name = "fishing"

// fmt.Printf("%+v",member1)

// ! Anonymous Struct
// this struct not declared first as a new struct data type, but declared while created new variable

// var member2 = struct {
// 	position string
// 	person Person
// }{}
// dont forget {} to initialize struct

// member2.position = "Data Analyst"
// member2.person.age = 35
// member2.person.name = "Komar"

// fmt.Printf("%+v",member2)

// var member3 = struct {
// 	position string
// 	person Person
// }{
// 	position: "UI/UX",
// 	person: Person{name : "Sam", age: 27},
// }
// fmt.Printf("%+v",member3)

// ! Slice to Struct

// var peoples = []Person{
// 	{name: "Sam", age : 15},
// 	{name: "Rick", age : 17},
// 	{name: "Ange", age : 16},
// }

// for _,value := range peoples {
// 	fmt.Printf("%+v",value)
// }

// ! Slice to Anonymous Struct

// var members = []struct{
// 	position string
// 	person Person
// }{
// 	{position: "UI/UX", person: Person{name: "Alex", age : 33}},
// 	{position: "FrontEnd", person: Person{name: "Smith", age : 35}},
// }

// for _,value := range members {
// 	fmt.Printf("%+v",value)
// }

// }

// !! FUNCTION
/*
	syntax function in go

	* func function_name(params)

	Initiation function

	* func expFunction(param1 string, param2 int)
*/

// func main() {
// var peoples = []struct{
// 	name string
// 	age int
// 	country string
// 	address string
// 	hobby string
// 	dream string
// }{
// 	{name : "Sebastian", age : 20, country: "United States", address : "New York", hobby : "Playing Games", dream: "Programmer"},
// 	{name : "Alexander", age : 23, country: "United Kingdom", address : "London", hobby : "Listening Music", dream: "Singer"},
// 	{name : "Angela", age : 20, country: "Canada", address : "Toronto", hobby: "Drawing", dream: "Architect"},
// }

// for _,value := range peoples {
// 	greet(value.name, value.age)
// 	secondGreet(value.country, value.address)
// 	fmt.Println(strings.Repeat("-", 50))
// }

// ! Function (Return)

// get return value
// printMsg := thirdGreet("Hello", "Sam")
// fmt.Println(printMsg)

//
// ! Function (Returning Multiple Value)
// get multiple return value
// multipleValue, dividedValue := calculate(5)
// fmt.Println(multipleValue)
// fmt.Println(dividedValue)

//
// ! Function (Predefined Return)
// newMultipleValue, newDividedValue := newCalculate(7)
// fmt.Println(newMultipleValue)
// fmt.Println(newDividedValue)

//
// ! Function (variadic func)
// get return variadic func
// studentList := printStudent("wati", "tono", "yandi", "beckto")
// fmt.Println(studentList)

// numberList := []int{1,2,3,4,5,6,7,8,9}
// spread slice
// totalSum := sumNum(numberList...)

// fmt.Println("Total Sum :", totalSum

// averageValue("sam", 50,75,80,90)

// }

// if the data type between params is different we declare data type after params
// func greet(name string, age int){
// 	fmt.Println("Hello there, My name is",name, "and I'm",age, "years old")
// }

// if the data type between params is similar we declare data type in the end after declare all param
// func secondGreet(country, address string){
// 	fmt.Println("I'm from", country)
// 	fmt.Println("I live in", address)
// }

// ! Function (return)

// if we want give a return we should add data types of return when we declare the function
// func thirdGreet(message, name string) string {
//  	var result string = fmt.Sprintf("%s %s", message, name)
// 	return result
// }

// ! Function (returning multiple value)

// we can returning multiple value with go, and we should add multiple data types of return data that we want when we declare the function, and we can return multiple value with different data types
// func calculate(number float64) (float64, float64) {
// 	multiple := number * number
// 	divided := number / number

// 	return multiple, divided
// }

//
// ! Function (predefined return values)
// we can defined the values that we want to return while we declare the function

// func newCalculate(number float64) (multiple float64, divided float64) {
// exp
// resMulti := number * number

// we should reassigned return value to store the return of function
// multiple = resMulti
// divided = number / number

// return
// }

//
// ! Function (variadic func)
// variadic func in go can store infinite argument

// exp #1
// func printStudent(names ...string) []string {
// 	var result []string

// 	for i := 0; i < len(names); i++ {
// 		result = append(result, names[i])
// 	}

// 	return result
// }

// exp #2
//  func sumNum(number ...int) int {
//  	var result int

//  	for _, value := range number {
//  		result += value
//  	}

//  	return result
//  }

// exp #3
// reguler parameters and variadic parameter
// to use this we should add variadic parameter in the last of parameters function

// func averageValue(name string, testScores ...float64){
// 	var totalScore float64

// 	for _, value := range testScores {
// 		totalScore += value
// 	}

// 	var avgVal float64 = totalScore / float64(len(testScores))

// 	fmt.Println("The student with name :", name)
// 	fmt.Println("Avg Score :", avgVal)
// }

//  !! METHOD
/*
	method is a function a attached to a datatypes, but method mostly used in struct data type
*/
// syntax
// * func(reciver_name Type) method_name(parameter_list)(return_type){}

//
// ! Method usage

// type Person struct {
// 	name string
// 	age int
// }

// // declare method introduce	from struct person
// func (p Person) Introduce(message string) string {
// 	return fmt.Sprintf("%s My name is %s and I'm %d years old", message, p.name, p.age)
// }

// func main(){
// accessing person struct and giving value to struct
// person := Person{name: "Samuel", age : 15}

// we can reach this method after accessing the struct
// 	fmt.Println(person.Introduce("Hii!"))
// }

//
// ! Method usage with pointer
/*
	we can used pointer in a method to change the real value from a	data type which was that pointer
*/

// type Person struct {
// 	name string
// 	age int
// }

// func (p Person) ChangeName1() {
// 	p.name = "Dinar"
// }

// change the value with memory address pointer
// func (p *Person) ChangeName2() {
// 	p.name = "Dinar"
// }

// func main() {
// 	person := Person{name : "Alex", age: 23}

// 	person.ChangeName1()
// 	fmt.Println("change value with ChangeName1 :", person.name )

// change value with pointer
// 	person.ChangeName2()
// 	fmt.Println("change value with ChangeName2 :", person.name )
// }

//
//
// !! REFLECT
/*
for inspection variable, take information from variable or manipulate it

from that package, they have 2 packages the most important to know :

1. reflect.ValueOf() to return object in reflect.Value types, which contains information related to value of variable we are looking for
2. reflect.TypeOf() to return object in reflect.Type types, that object contains information related to data types variable we are looking for
*/

func main() {
	var angka int = 64

	valueCheck := reflect.ValueOf(angka)
	typeCheck := reflect.TypeOf(angka)

	fmt.Println("value :", valueCheck)
	fmt.Println("type :", typeCheck)

}

//
//
// !! EXPORTED & UNEXPORTED
/*
in GO lang every different folder will considered as a package

we can use variable or data type from other packages if that variable and data types has been exported from the packages

how to we export a variable or a data types is start write a variable or a data types with uppercase
*/

//  !! INIT FUNCTION

/*
we can running a function before func main with init function

mostly we used this for config a program before main program is running

*/

// func main() {

// }