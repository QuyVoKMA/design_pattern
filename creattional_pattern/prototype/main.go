package main

import (
	"crea/prototype/a"
	"fmt"
)

func main() {
	file1 := &a.File{Name :"File 1"}
	file2 := &a.File{Name :"File 2"}
	file3 := &a.File{Name :"File 3"}

	folder1:=&a.Folder{
		Children:[]a.INode{file1},
		Name: "Folder 1",
	}

	folder2:=&a.Folder{
		Children:[]a.INode{folder1,file2,file3},
		Name: "Folder 2",
	}

	fmt.Println("\n Printing for folder 2")
	folder2.Print("   ")
	cloneFolder :=folder2.Clone()
	fmt.Println("\n PRinting for clone folder 2")
	cloneFolder.Print("   ")
}
