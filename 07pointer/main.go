package main
 import "fmt"
func main()  {
	fmt.Println("Welcome to the class of pointers")
// Here only creating a pointer
	// var ptr *int
	// fmt.Println("The value of ptr is", ptr)

	myNumber:=23

	// here creating a pointer that are referencing some varaible
	var ptr = &myNumber
	// this ptr pointer indicating myNumber and in ptr address of myNumber is stored
	fmt.Println("The address of myNumber is", ptr)
	
	// here we are printing the value of myNumber using ptr
	fmt.Println("The value of myNumber is", *ptr)

	*ptr = *ptr + 3  //this line indicating ,we do operation on value ,not the address,therefore value of myNumber will change


	// Both of these two line give the same ouput ie.---26
	fmt.Println("The new value of myNumber is", *ptr)
	fmt.Println("The new value of myNumber is", myNumber)
}