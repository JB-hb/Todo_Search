package list

import (
	"fmt"
	"errors"
	"strings"
)

//node struct and methods
type Node_type struct{
	next *node
	path string
	line int 
	todo string 
}

func (node Node_type) Has_next() bool{
	if(node.next == nil){
		return false
	}
	return true
}

func (node Node_type) Get_next() *Node_type{
	return node.next
}

func (node Node_type) Set_next(next *Node_type){
	node.next = next
}

func (node Node_type) Get_values() (string, string, int){
	return node.path, node.todo, node.line
}

//list struct and methods

type List_type struct{
	head *Node_type
	name string
}

func (list List_type) Is_empty() bool{
	if(list.head == nil){
		return true
	}

	return false
}

func (list List_type) Get_len(pos int, node *Node_type) int{
	if(list.Is_empty() == true){
		return 0
	}

	if(pos == 0){
		return list.Get_len(pos + 1, list.head)	
	}

	if(node.Has_next() == true){
		return list.Get_len(pos + 1, (*node).Get_next())
	}

	return pos + 1
}

func (list List_type) Push_node(node *Node_type)int, error{
	if(list.Is_empty() == true){
		list.head = node
		return 1, errors.New("the list is empty")
	}	

	temp *node := list.head
	list.head = node
	list.(*head).Set_next(temp)
	return 0, nil

} 

func (list List_type) Select_node_index(index int, pos int, node *Node_type) (*Node_type, error){
	if(list.Is_empty() == true){
		return nil, errors.New("the list is empty")	 
	}

	if(index > list.Get_len()){
		return nil, errors.New("index higher than the list length")
	}

	if(index == 0){
		return list.head, nil
	}

	if(index == pos){
		return node, nil
	}

	if(pos == 0){
		return list.Select_node_index(index, pos+1, list.(*head).Get_next())
	}

	return list.Select_node_index(index, pos+1, (*node).Get_next())
}

func (list List_type) Delete_node_index(index int, pos int, node *Node_type, prev *Node_type) (int, error){
	if(list.Is_empty() == true){
		return 1, errors.New("The list is empty")
	}

	if(list.Get_len() < index){
		return 2, errors.New("Index higher than list length") 
	}
	
	if(index == 0){
		list.head = (*head).Get_next()
		return 0, nil
	}

	if(index == pos){
		(*prev).Set_next((*node).Get_next())	
		return 0, nil
	}

	if(pos == 0){
		return Delete_node_index(index, pos + 1, list.(*head).Get_next(), list.head)
	}

	return Delete_node_index(index, pos + 1, (*node).Get_next(), node)
}

func (list List_type) Print_list(base_path string, node *Node_type) int, error{
	if(list.Is_empty() == true){
		return 1, errors.New("the list is empty")
	}
	
	if(node == nil){
		fmt.Printf("%-40s|%-40s|%-5s\n","To-do", "Path", "Line")
		path string, todo string, line int := list.(*head).Get_values() 
		new_path, found := strings.CutPrefix(path, base_path)
		if(found == true){
			fmt.Printf("%-40s|%-40s|%-5d\n", todo, new_path, line)
		}
		if(list.(*head).Has_next() == true){
			return list.Print_list(base_path, list.(*head).Get_next())
		}
		return 0, nil
	}

	path string, todo string, line int := (*node).Get_values() 
	new_path, found := strings.CutPrefix(path, base_path)
	if(found == true){
		fmt.Printf("%-40s|%-40s|%-5d\n", todo, new_path, line)
	}

	if((*node).Has_next() == true){
		return list.Print_list(base_path, list.(*head).Get_next())
	}

	return 0, nil
}
