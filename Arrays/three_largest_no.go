
// Write a function that takes in an array of at least three integers and, 
// without sorting the input array, returns a sorted array of the three largest 
// integers in the input array.

package main

import (
	"fmt"
)

func three_max_noo(arr []int) []int{
	a,b,c:=int(-1<<63),int(-1<<63),int(-1<<63)
	// iterate through the array
	for _,v:=range arr{
		if v>a{
			a,b,c=v,a,b
		}else if v>b{
			b,c=v,b
		}else if v>c{
			c=v
		}
	}
	// return a array of three elements
	return []int{a,b,c}
}


func main() {
	arr:=[5]int{1,2,3,4,5}
	// pass the arr to the function
	msg:=three_max_noo(arr[:])
	fmt.Println(msg)
}
