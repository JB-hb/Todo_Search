package main

import "fmt"
import "os"
import "github.com/JB-hb/Todo_Search/List/list"

func main(){

	var path string;
	
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

	fmt.Println(list.TestPack())

}
