package main

import "fmt"

func main() {
	// just like we used make( ) function to create a "slice" ,we also used it to make "Maps"
	fmt.Println("Maps in Go lang")
	var languages = make(map[string]string)
	languages["JS"]="JavaScript"
	languages["RB"]="Ruby"
	languages["PY"]="Python"
	fmt.Println("List of all languages:",languages)

	// printing one element from MAP
	fmt.Println("Details of a language:",languages["RB"])

	// deleting an element from MAP
	delete(languages,"JS")
	fmt.Println("List of all languages:",languages)

	//looping through MAP
	for key,value := range languages{
		fmt.Printf("For key %v,value is %v\n",key,value)
	}

}