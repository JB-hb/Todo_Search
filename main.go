package main

import (
	//"fmt"
	//"os"

	"todo/List"
)

func main(){

	var lst list.List_type
	var code int
	var err error
	node := list.Node_type{Path: "path/one", Line: 23, Todo: "test-node"}

	lst.Push_node(&node)
	lst.Print_list("path", nil)

	if(err == nil && code == 0){
		return
	}



	/*
	switch{
	case len(os.Args) == 2:
		path = os.Args[1]
		fmt.Println(path)
	case len(os.Args) > 2: 
		fmt.Println("Error: insert valid arguments")
	default:
		fmt.Println("Insert path of the project")
		fmt.Scanln(&path)
		fmt.Printf("the path is %s\n", path)
	}
	*/

}
