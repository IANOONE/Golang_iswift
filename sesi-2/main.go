package main

func main() {

	// ARRAY

	// initiations:

	// defined length **
	// var animals = [3]string{"cat", "dog", "mouse"}

	// inferred length, based on how many value on array that we add **
	// animals := [...]string{"cat", "dog", "mouse"}

	// reassign
	// animals[2] = "bird"

	// fmt.Printf("%#v\n", len(animals))

	// breakdown key index and value
	// for key, animal := range animals {
	// 	fmt.Printf("Cara pertama, index : %d, value %s\n", key, animal)
	// }

	// breakdown key and call the index to show the data of the indexes
	// for key := range animals {
	// 	fmt.Printf("Cara pertama, index : %d, value %s\n", key, animals[key])
	// }

	// used data length to loop
	// for i := 0; i < len(animals); i++ {
	// 	fmt.Printf("Cara pertama, index : %d, value %s\n", i, animals[i])``
	// }

	// multidimentionals array

	// balances := [2][3]string{{"1", "2", "3"}, {"4", "5", "6"}}

	// loop outside
	// for _, arr := range balances {
	// loop inside
	// 	for _, value := range arr {
	// 		fmt.Printf("%s ", value)
	// 	}
	// 	fmt.Println()
	// }

	//
	//
	// SLICE

	// initiations:
	// var fruits = []string{"apple", "mango", "cocoa"}
	// when used slice we shouldn't add parameter how long long value will store, it's dynamically

	//fmt.Println(len(fruits)) // return the length of slice (the numbers of elements in the slice)
	//fmt.Println(cap(fruits))  // return capacity of the slice (the number of elements the slice can grow or shrink to)

	// data manipulation using slice

	// 1. make function

	// ?? slice_name := make([]type, length, capacity)
	// fruits := make([]string, 4)

	// fruits[0] = "mango"
	// fruits[1] = "strawberry"
	// fruits[2] = "banana"
	// fruits[3] = "banana"

	// fmt.Println(fruits)
	// fmt.Println(len(fruits)) // how much elements stored in the slice
	// fmt.Println(cap(fruits)) // capacity of the slice (the number of elements the slice can grow or shrink to)

	// 2. append function
	// for added new elements in the slice, the elements will added on last index
	// when we make slice and set the len of variable, the value is not strict, it would be dynamic if we add new value using append

	// fruits = append(fruits, "melon")
	// fmt.Println(fruits)

	// 3. append function with ellipses
	// same idea like spread array or object in js

	// forestAnimals := []string{"Monkey", "Lion", "Giraffe"}
	// seaAnimals := []string{"Shark", "Whale", "Penguin"}

	// allAnimals := append(forestAnimals, seaAnimals...)
	// fmt.Println("%v",forestAnimals)
	// fmt.Println("%v",seaAnimals)
	// fmt.Printf("%v", allAnimals)

	// 4. copy function
	// this function will copying elements from a slice to anothers slice. when we do this the elements at the slice that we want as places copied data or value.

	// animals1 := []string{"Monkey", "Lion", "Giraffe"}
	// animals2 := []string{"Shark", "Whale", "Penguin"}

	// copyFunc := copy(animals1, animals2)
	// declare
	// copyFunc will get return how many values change

	// fmt.Println("Animals 1 =>", animals1)
	// fmt.Println("Animals 2 =>", animals2)
	// fmt.Println("Copied Elements =>", copyFunc)

	// slicing function

	// [start:stop]
	// start = first index slice that we want to slice
	// stop = sequences data

	// animals1 := []string{"Monkey", "Lion", "Giraffe"}

	// animals2 := animals1[1:3]
	// slice data from index no 1, until sequence np 3
	// fmt.Printf("%#v\n", animals2)

	// animals3 := animals1[1:]
	// slice data from index no1, until last sequence data
	// fmt.Printf("%#v\n", animals3)

	// animals4 := animals1[:1]
	// slice data from index no 0, until sequence data no 1
	// fmt.Printf("%#v\n", animals4)

	// animals5 := animals1[:]
	// equal to animals1[:len(animals1)]
	// fmt.Printf("%#v\n", animals5)

	// 5. combining slicing & append

	// colors := []string{"red", "orange", "yellow", "green", "biru"}

	// colors = append(colors[:2], "pink")
	// fmt.Printf("%#v\n", colors)

	// 6. backing array
	/* every we create a slice, go will make a hidden array automatically that we called BACKING ARRAY. the reason is to saved an elements in slice. GoLang implement slice as an data stucture that called header slice. header slice consist of :
	1. address memory from backing array
	2. length of slice that can get from function len
	4. capacity of slice that can get from function cap
	*/

	// lang1 := []string{"phyton", "ruby", "java", "go", "javascript"}
	// lang2 := lang1[2:4]

	// lang2[0] = "php"

	// 8. cap function

	// newFruits1 := []string{"strawberry", "mango", "apple", "pineapple", "melon"}

	// fmt.Println("len =>", len(newFruits1)) // how much elements stored in the slice
	// fmt.Println("cap =>" ,cap(newFruits1)) // capacity of the slice (the number of elements the slice can grow or shrink to or how much elements can store)
	// fmt.Println(strings.Repeat("-",20))

	// newFruits2 := newFruits1[0:3]

	// fmt.Println("len =>", len(newFruits2))
	// fmt.Println("cap =>" ,cap(newFruits2))
	// fmt.Println(strings.Repeat("-",20))

	// newFruits3 := newFruits1[1:]

	// fmt.Println("len =>", len(newFruits3))
	// fmt.Println("cap =>" ,cap(newFruits3))
	// fmt.Println(strings.Repeat("-",20))

	//
	//
	// MAP
	// same idea like array of object in js, but  data_types key or value is not dynamic, exp: if we set key as an string, all of key data should be a string they can't afford others data types

	// variable_name := map[data_type]data_type

	// exp iniations

	// person1 := map[string]string{}
	// person1["name"] = "Brilian"
	// person1["birthplace"] = "Sleman"
	// person1["age"] = "24"
	// person1["address"] = "Pontianak"

	// fmt.Printf("%#v\n", person1)
	// fmt.Println(strings.Repeat("-", 5-0))
	// person2 := map[string]int64{
	// 	"nik":        3201131005990001,
	// 	"birthdate": 10051999 ,
	// }
	// fmt.Printf("%#v\n", person2)

	// Data manipulation with map

	// 1. Deleting value
	// delete by key
	// delete(person1, "address")
	// fmt.Println(person1)
	// fmt.Println(strings.Repeat("-", 5-0))

	// 2. detecting value

	// value , exist := person1["age"]
	// if(exist) {
	// 	fmt.Println(value)
	// 	fmt.Println(strings.Repeat("-", 5-0))
	// } else {
	// 	fmt.Println("value doesnt exist")
	// 	fmt.Println(strings.Repeat("-", 5-0))
	// }

	// 3. combining slice and map

	// biodata := []map[string]string{
	// 	{"name" : "ian", "address" : "pontianak"},
	// 	{"name" : "brilian", "address" : "jakarta"},
	// 	{"name" : "brian", "address" : "bogor"},
	// }

	// fmt.Println(biodata)

	// for key, person := range biodata {
	// 	fmt.Printf("index : %d, name : %s, address : %s\n", key, person["name"], person["address"])
	// }

	//
	//
	// POINTER
	/*  a variable for saving memory address variable
	using pointer in the conditions :
	# to avoid repeating/doubling data
	# empty value meaningfull // comparation value variable with nill
	# change value outside function
	*/

	// how to declare pointer we should add asterisk (*) before writing data type
	// variable_name := *data_type

	// initiation pointer :
	// firstExample  *int
	// secondExample  *string
	// _ , _ = firstExample, secondExample

	// 1. initiations

	// var firstNumber int = 4

	// when we want to get memory address from value we should add ampersand (&) sign
	// we should equalizing data_type pointer with the data_type which we initialize
	// var secondNumber *int = &firstNumber

	// fmt.Println()
	// fmt.Println("firstNumber (value)", firstNumber)
	// fmt.Println("firstNumber (address memory)", &firstNumber)
	// fmt.Println(strings.Repeat("-", 50))

	// we should convert pointer value to know what's the real value
	// fmt.Println("secondNumber (value)", *secondNumber)
	// fmt.Println("secondNumber (address memory)", secondNumber)

	// 2. changing value through pointer
	/* pointer used for saved memory address, if we change the value of a pointer, other variable that have a same memory address will be following changes the value */

	// var firstGroup string = "BlackPink"
	// var secondGroup *string = &firstGroup

	// fmt.Println("First Group (value) : ", firstGroup)
	// fmt.Println("First Group (memory address) :", &firstGroup)

	// fmt.Println("Second Group (value) : ", *secondGroup)
	// fmt.Println("Second Group (memory address) :" ,secondGroup)

	// *secondGroup = "NewJeans"

	// fmt.Println("First Group (value) : ", firstGroup)
	// fmt.Println("First Group (memory address) :", &firstGroup)

	// fmt.Println("Second Group (value) : ", *secondGroup)
	// fmt.Println("Second Group (memory address) :" ,secondGroup)

	// 3. pointer as parameter

	// var numero string = "numero"

	// fmt.Println("Before:", numero)
	//
	// fmt.Println( strings.Repeat("-", 20))
	//
	// change the value variable age using address memory using ampersand (&)
	// changeValue(&numero)
	// fmt.Println("After:", numero)

	// 3. nill is not zero value

	// nill is a value not a data types
	// basic data types have a zero value or default value. that's means if we're declare variable without set default value, it still have a default value

	// nill can't used on basic data types,
	//  some data types can use nill as a value :
	// Pointer
	// Slice
	// Map
	// Channel
	// Empty Interface / Interface[]

	// exp :

	// var name = "ian"

	// var student1 *string
	// student1 = &name

	// if student1 != nil {
	// 	fmt.Println("Student name:", *student1)
	// } else  {
	// 	fmt.Println("Student doesn't exist")
	// }

	// age := 20

	// if(age != nil){}     => checked int value, alert will showing nill can't used for int value

}

// func changeValue(someWord *string){
// 	*someWord = "word"
// }
