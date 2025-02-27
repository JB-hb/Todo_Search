package main

import (
	"os"
	"fmt"

	"todo/List"
	"todo/Search"

)


func main(){

	var lst list.List_type
	var path string

	switch{
	case len(os.Args) == 2:
		path = os.Args[1]
		fmt.Println(path)
	case len(os.Args) > 2: 
		fmt.Println("Error: insert valid arguments")
	default:
		fmt.Println("Insert path of the project")
		fmt.Scanln(&path)
	}

	search.Find_todos(path, &lst)

	lst.Print_list(path, nil, 0)


}
