package list

import (
	"fmt"
	"errors"
	"strings"
)

//node struct and methods
type Node_type struct{
	Path string
	Line int 
	Todo string 
	Next *Node_type
}

func (node Node_type) Has_next() bool{
	if(node.Next == nil){
		return false
	}
	return true
}

func (node Node_type) Get_next() *Node_type{
	return node.Next
}

func (node Node_type) Set_next(next *Node_type){
	node.Next = next
}

func (node Node_type) Get_values() (string, string, int){
	return node.Path, node.Todo, node.Line
}

func (node Node_type) Get_path() (string){
	return node.Path
}

//list struct and methods

type List_type struct{
	Head *Node_type
	name string
}

func (list *List_type) Is_empty() bool{
	if(list.Head == nil){
		return true
	}

	return false
}

func (list *List_type) Get_len(pos int, node *Node_type) int{
	if(list.Is_empty() == true){
		return 0
	}

	if(pos == 0){
		return list.Get_len(pos + 1, list.Head)	
	}

	if(node.Has_next() == true){
		return list.Get_len(pos + 1, (*node).Get_next())
	}

	return pos + 1
}

func (list *List_type) Push_node(node *Node_type)(int, error){
	if(list.Is_empty() == true){
		list.Head = node
		return 0, nil
	}	

	var temp *Node_type = list.Head
	list.Head = node
	list.Head.Set_next(temp)
	return 0, nil

} 

func (list *List_type) Select_node_index(index int, pos int, node *Node_type) (*Node_type, error){
	if(list.Is_empty() == true){
		return nil, errors.New("the list is empty")	 
	}

	if(index > list.Get_len(0, nil)){
		return nil, errors.New("index higher than the list length")
	}

	if(index == 0){
		return list.Head, nil
	}

	if(index == pos){
		return node, nil
	}

	if(pos == 0){
		return list.Select_node_index(index, pos+1, list.Head.Get_next())
	}

	return list.Select_node_index(index, pos+1, (*node).Get_next())
}

func (list *List_type) Delete_node_index(index int, pos int, node *Node_type, prev *Node_type) (int, error){
	if(list.Is_empty() == true){
		return 1, errors.New("The list is empty")
	}

	if(list.Get_len(0, nil) < index){
		return 2, errors.New("Index higher than list length") 
	}
	
	if(index == 0){
		list.Head = list.Head.Get_next()
		return 0, nil
	}

	if(index == pos){
		prev.Set_next(node.Get_next())	
		return 0, nil
	}

	if(pos == 0){
		return list.Delete_node_index(index, pos + 1, list.Head.Get_next(), list.Head)
	}

	return list.Delete_node_index(index, pos + 1, node.Get_next(), node)
}

func (list *List_type) Print_list(base_path string, node *Node_type, cont int) (int, error){
	if(list.Is_empty() == true){
		fmt.Println("aja")
		return 1, errors.New("the list is empty")
	}

	
	if(node == nil){
		fmt.Printf("%-40s|%-40s|%-5s\n","To-do", "Path", "Line")
		path, todo, line := list.Head.Get_values() 
		new_path, found := strings.CutPrefix(path, base_path)
		if(found == true){
			fmt.Printf("%-40s|%-40s|%-5d(%d)\n", todo, new_path, line,cont)
		}
		if(list.Head.Has_next() == true){
			return list.Print_list(base_path, list.Head.Get_next(), cont+1)
		}
		return 0, nil
	}

	path, todo, line := node.Get_values() 
	new_path, found := strings.CutPrefix(path, base_path)
	if(found == true){
		fmt.Printf("%-40s|%-40s|%-5d(%d)\n", todo, new_path, line,cont)
	}

	if((*node).Has_next() == true){
		return list.Print_list(base_path, list.Head.Get_next(),cont+1)
	}

	return 0, nil
}
