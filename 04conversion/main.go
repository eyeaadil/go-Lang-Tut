package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app");
	fmt.Println("Please rate our pizza between 1 and 5");

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ",input);

	// numrating := input + 1;
	// here i want to add 1 in string,which are impossible because num can not directly added to string
	// so we have to convert string to int


	// here we are converting string to int
	numrating,err := strconv.ParseFloat(strings.TrimSpace(input),64);

	// here 64 is bitsize of int,float etc 

	if(err !=nil){
		fmt.Println("Please provide a number",err);
	} else{
		fmt.Println("Added 1 to string : ",numrating+1)
	}
}