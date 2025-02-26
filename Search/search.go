package search

import(
	"os"
	"buffio"
	"path/filepath"
	"fmt"
	"strings"

	"todo/List"
)

func Find_todos(path string, list *list.List_type){
	files, err := os.ReadDir(path)
	for i, file := range files{
		if file.IsDir(){
			new_path string := filepath.Join(path, file.Name) 				
			Find_todos(new_path, list)
			continue
		}else{
			value, errR := os.Open(filepath.Join(path, file.Name))
			if(errR != nil){
				fmt.Printf("Error opening file %s\n", filepath.Join(path, file.Name))
				continue
			}
			scanner := buffio.NewScanner(value)
			cont int := 0
			for scanner.Scan(){
				cont += 1
				line := strings.ToLower(scanner.Text())
				if(strings.Contains(line, "todo")){
					node list.Node_type := {Todo: Line, Path: filepath.Join(path, file.Name), Line: cont}
					list.Push_node(node)
				}
			}
		}
	}
}
