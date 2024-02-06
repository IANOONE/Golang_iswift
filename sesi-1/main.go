package main

import (
	"fmt"
	"math"
)

const (
	statusNow = iota + 1
	statusPending
	statusOld
)

func main() {
	// fmt.Println("Hello World")

	// initiate variable with datatypes
	// var title string = "Golang Cohort 3"
	// var totalMember int = 20

	// initiate variable without datatypes
	// var title1 string
	// title1 = "Golang Beginner"

	// fmt.Println("Course title", title)
	// fmt.Println("Total Member", totalMember)

	// with := golang will automatically detect what's data type of the value
	// title := "Golang cohort 3"
	// totalMember := 20

	// fmt.Printf("%T, %T \n", title, totalMember)
	// fmt.Println("Course title", title)

	// i := 0

	
	// for i := 1; i <= 3; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// i := 1

	// for i <= 3 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// for {
	// 	fmt.Println("Angka", i)

	// 	i++
	// 	if( i == 3){
	// 		break
	// 	}
	// }

	// for x := i; x <=3; x++ {
	// 	for j := x; j <= 5; j++ {
	// 		fmt.Println(i, " ")
	// 	}

	// 	fmt.Println()
	// }



	/* var dataa []string
	dataa = []string{"aaa", "bbb"}


	for _, data := range dataa {
		fmt.Println("Data", ":", data)
	} */

	// exp := 5

	// for i := 1; i <= exp; i++ {

	// 	if(math.Sqrt(i) === 5){
	// 		fmt.Println("cube")
	// 	}

	// 	fmt.Println(i)


	// }

	nilai := 64
	exp := math.Sqrt(float64(nilai))
	res := math.Pow(float64(exp), 2) == float64(nilai)  



	fmt.Println(res)
}
