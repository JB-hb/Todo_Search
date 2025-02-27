package search

import(
	"os"
	"bufio"
	"path/filepath"
	"fmt"
	"strings"

	"todo/List"
)

//todo : search

func Find_todos(path string, listp *list.List_type){
	files, err := os.ReadDir(path)
	if (err != nil){
		fmt.Printf("error reading: %s\n", path)	
	}
	for _, file := range files{
		if file.IsDir(){
			new_path := filepath.Join(path, file.Name()) 				
			Find_todos(new_path, listp)
			continue
		}else{
			value, errR := os.Open(filepath.Join(path, file.Name()))
			if(errR != nil){
				fmt.Printf("Error opening file %s\n", filepath.Join(path, file.Name()))
				continue
			}
			reader := bufio.NewReader(value)
			cont := 0
			for{

				line, errL := reader.ReadString('\n')
				if errL != nil{
					break
				}
				cont += 1
				if(len(line) > 1 && line[0] == '/' && line[1] == '/'){
					if (strings.Contains(line, "todo")){
						line,_ = strings.CutPrefix(line, "//")
						line = line[:len(line) - 1]
						node := list.Node_type {Todo: line, Path: filepath.Join(path, file.Name()), Line: cont}
						listp.Push_node(&node)
					}
				}
			}
		}
	}
}
