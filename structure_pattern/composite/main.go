package main

import (
	p "design_pattern/structure_pattern/composite/a"
)

func main(){
	file1 :=&p.File{Name:"File1"}
	file2 :=&p.File{Name:"File2"}
	file3 :=&p.File{Name:"File3"}
	folder1 :=&p.Folder{Name: "Folder1"}
	folder1.Add(file1)
	folder2 :=&p.Folder{Name: "Folder2"}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")

	// Mong muốn là nó đi tìm kiếm tất cả các file1, file2, file3
	// Nó sẽ tìm vào tròng Folder 2 -> file2 ->file 3 -> folder1->file1
	//Thiet lap cac composite, rồi cài đặt hết vào folder2 nó sẽ gọi ra tất cả theo trình tự.
}
