package helpers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)


func CheckArg(arg ...string)  string {
	var argStored string

	 if(len(arg[1:]) < 1 ){
		fmt.Println("The argument can't be empty, you should add name or id number!")
		fmt.Println("exp : 'go run main.go Alex' or 'go run main.go 1'")
	 } else  if(len(arg[1:]) > 1 ){
		fmt.Println("You can't add more than 1 argument, please reassign argument!")
		fmt.Println("exp : 'go run main.go Alex' or 'go run main.go 1'")
	 }  else {
		argStored = strings.Join(arg[1:], "")
	 }

	 return argStored
}

func FetchData(check string)  (result map[int]Student) {
	// get data students
	students := Students()

	// initiate map for result
	result = make(map[int]Student)
	// 

		for key, entry := range students {
			if check == strconv.Itoa(key) || strings.EqualFold(check, entry.Name) {
				result[key] = Student{
					Id:         key,
					Name:       entry.Name,
					Address:    entry.Address,
					Profession: entry.Profession,
					Reason:     entry.Reason,
				}
			} 
		}
	if(len(result) < 1){ 
		fmt.Printf("Data with id or name %s not found \n", check)
	}
	return result
}

func PrintResult(result map[int]Student) {
	for _, res := range result {
		v := reflect.ValueOf(res)
		f := reflect.TypeOf(res)
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("%s : %v\n", f.Field(i).Name, v.Field(i))
		}
		if len(result) > 1 {
			fmt.Println(strings.Repeat("-", 50))
		}
	}
}