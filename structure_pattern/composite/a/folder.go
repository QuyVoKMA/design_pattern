package a

import "fmt"
/*Folder is struct*/
type Folder struct {
	components []Componet
	Name string
}
/*Search is function*/
func (f *Folder) Search(keyword string){
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite :=range f.components{
		composite.Search(keyword)
	}
}

/*Add is function*/
func (f *Folder) Add(c Componet){
	f.components =append(f.components, c)
}
