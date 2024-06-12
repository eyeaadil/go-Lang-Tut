package main

import (
	"fmt"
	"reflect"
)


const LoginToken = "Adfdvdvndn"  //const can be used anywhere
// here capital L indicates that the the given variable is public
func main() {
	// fmt.Println("this is variables lecture");

	// if we declare and do not use then it show an error 
	var username string = "this is string"
	fmt.Println(username);

	// this line give the type of variable

	// for type %T with printf give type
	// one of the package used fro type is "reflect"

	fmt.Println("Type of variable username:", reflect.TypeOf(username))
	fmt.Printf("Variable is of Type: %T \n",username)
	
	
	var isLogIn bool = false 
	fmt.Println(isLogIn);
	fmt.Printf("Variable is of Type: %T \n",isLogIn)

	// default values and some aliases
	var anotherVariable int;
	fmt.Println(anotherVariable);
	fmt.Printf("Variable is of Type: %T \n",anotherVariable)

	// declare implicit type
	var website = "www.google.com"
	fmt.Println(website);

	// no var style   Outside any function and method this is not allowed 
	numberOfUser :=3000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken);
	fmt.Printf("Variable is of Type: %T \n",LoginToken)

	
	
}