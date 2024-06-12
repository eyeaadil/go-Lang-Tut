package main

import (
	"fmt"
	"time"
)

// time package exist
// in this package time format is most used


func main()  {
	fmt.Println("Welcome to time lecture")

	presentTime := time.Now()
	fmt.Println("The present time is",presentTime)

	// time format
	// 2006-01-02 15:04:05 Monday   ---->  Always gives this format value
	fmt.Println(presentTime.Format("2006-01-02 15:04:05 Monday"))


	// created date

	createdDate:=time.Date(2003,time.February,10,23,23,0,0,time.UTC)
	fmt.Println("Time is: ",createdDate)
}


