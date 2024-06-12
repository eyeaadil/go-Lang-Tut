package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slice")
	var fruitList = []string{"Apple","Mango","papaya"}
	// In Array size is fixed
// Slice is like a vector size is not fixed,size may be increases or decreases
	fmt.Printf("Type of fruitList is %T",fruitList)

	// for adding element to the slice ,we use append( ) function
	fruitList = append(fruitList, "banana")
	fruitList = append(fruitList, "grapes", "orange")
	fmt.Println(fruitList)

	// colon syntax
	// This syntax is used to make slice of slice 

	// [Apple,Mango,papaya,banana,grapes,orange]
	// Now make slice of this given slice
	fruitList=append(fruitList[2:4])
	// here 4 is excluded

	// fruitList=append(fruitList[2:] ) start from 2 and goes to last 
	// fruitList=append(fruitList[:4]) start from 0 and goes to 3 because 4 is excluded
	
	fmt.Println(fruitList)

	highScore:=make([]int, 4)
	highScore[0]=234
	highScore[1]=945
	highScore[2]=465
	highScore[3]=867
	// by using this it stastically allocate the memory so if we want to access "highScore[4]",it will give an error

	highScore=append(highScore, 555,666,321)  //As this line of code run new memory allocation for all these value
	fmt.Println(highScore)



	// fmt.Println(sort.IntsAreSorted(highScore))

	sort.Ints(highScore) //Ints sorts a slice of ints in increasing order.
	fmt.Println(highScore)  
	
	fmt.Println(sort.IntsAreSorted(highScore))  // give boolean value


	// How to remove value from slice based on the index
	var courses=[]string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println(courses)

	// with the help of append we add the data to slice,but we can also delete using this function
	// Here we want to delete "swift from slice"
	var index int=2
	courses=append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
	// append method reallocate the memory

	

}