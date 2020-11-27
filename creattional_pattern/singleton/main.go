package main

import (
	"crea/singleton/a"
	"fmt"
	"time"
)

func main(){
	for i:=0;i<10;i++{
		go func(){
		fmt.Printf("%p\n",a.GetInstance())
		}()
	}
	time.Sleep(time.Second * 10)
}
