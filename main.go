package main

import (
	"fmt"
)

var name string= "qwerty"
var age int = 20
const pi float64 = 3.14

 

func main() {
	name = "qwerty123"
	age = 30
	for i := 0; i < age; i++ {
		if i == 10 {
			
			result:= add(age, i)
			slice, ok:= pop(slice)
			sliceUnShift := unShift(slice, 6, 7, 8)
			fmt.Println("Hello, World!", name, age, i, result, len(slice), append(slice, 8), slice, ok, sliceUnShift)
			break
		}else{
			fmt.Println(i)
		}	
	}
	
}

func add(a int,b int) int {
	return a + b
}
func pop(slice []int) ([]int, bool) {
	if len(slice) == 0 {
		return slice, false
	}
	return slice[:len(slice)-1], true

}

func unShift(slice []int, nums ...int) []int {
	for i:= len(slice)-1; i>=0; i-- {
		slice[i+ len(nums)] = slice[i]
	}
	for i:= 0; i< len(nums); i++ {
		slice[i] = nums[i]
	}
	return slice
}