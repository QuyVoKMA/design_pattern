package main

import "fmt"

func main(){
	var arr =[]int32{1,2,3}
	var n int32
	var i int32
	if n>0 && n<=1000{
		for i=n-1;i>=0;i--{
			arr[i]+=1
			fmt.Printf("%v ",arr[i])
		}
	}

}
