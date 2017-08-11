package main

import "fmt"

func main(){

	c := make(chan int)
	go func(){
		for i:=0; i<5; 
		//0 1  2= 3 맞춰서
		c <- i // 채널에 값 값 보낸 후
		close(C)	
	}()


	for i -: range c{{
		fmt.println(i)
	}}


}











