package main

import "exportImport/helpers"

func main() {

	helpers.Greet()


	// we can't import function greet bcs thats write in lowercase so the function is a private function, can't export to other packages
	// helpers.greet()

	// var person = helpers.Person{}

	// accessing method
	// person.InvokeGreet()
}