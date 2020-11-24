package a

import "fmt"

/*Hàn hóa trong 1 cửu hàng*/

type Item struct {
	ObserverList []Observer
	Name string
	InStock bool
}

/*Hàm khỏi tạo*/
/*NewItem is function*/
func NewItem(name string) *Item {
	return &Item{
		Name:name,
	}
}
/*UpdateAvailability is function*/
func (i *Item) UpdateAvailability(){
	fmt.Printf("Item %s is now in stock\n", i.Name)
	i.InStock=true
	i.NotifyAll()
 }
/*Register is function*/
 func (i *Item) Register(o Observer){
 	i.ObserverList = append(i.ObserverList, o)

 }
/*DeRegister is function*/
 func (i *Item) DeRegister(o Observer){
 	i.ObserverList = removeFromSlice(i.ObserverList, o)
 }

 func (i *Item) NotifyAll(){
 	for _,observer :=range i.ObserverList{
 		observer.Update(i.Name)
	}
 }


 func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer{
 	l :=len(observerList)
 	for i, observer:=range observerList{
 		if observerToRemove.GetID() == observer.GetID() {
 			observerList[l-1], observerList[i]=observerList[i],observerList[l-1]
 			return observerList[:l-1]
		}
	}
	return observerList
 }