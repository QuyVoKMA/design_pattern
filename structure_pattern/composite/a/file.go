package a

import "fmt"
/*File is struct*/
type File struct {
	Name string

}
/*Search is function*/
func (f *File) Search(keyword string){
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.Name)
}
