package main

import (
	"bufio"
	"fmt"
	"os"
)
func main()  {
	// bufio package
	// os package

	// var welcome string = "Adil"
	// fmt.Println("Welcome to Golang",welcome);

	fmt.Println("take input from user")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:",reader)

	// comma ok || error ok  syntax

	// try catch is not present in the go lang
	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ",input)

	// _ this indiacating that if there is an error occur then neglect the error
	// instaed of _ we can use error to show the error 

}