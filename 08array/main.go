package main

import "fmt"

// array is very less used in go lang

func main()  {

	// Array are very less in used,instead of array ,slice Data Structure is Used

	// array declaration
	var fruitList [5] string

	// array initialisation
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Grapes"
	fruitList[3] = "Mango"
	fruitList[4] = "Orange"

	// array print
	fmt.Println(fruitList)
	// print array element one by one

	// for length of array :----> len( ) function used
	for i := 0; i < len(fruitList); i++ {
		fmt.Println("Elements of FruitslIst :",fruitList[i])
	}


	// Decalartion and initilisation simultaneously
	var vegList = [4]string{"peas","carrot","Potato","Tomato"}
	fmt.Println("VegetableList: ",vegList);
	
}