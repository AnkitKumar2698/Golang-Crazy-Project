package main

import "fmt"

func main(){
	var exit int
	exit=3
	var m map[string] int
	m=make(map[string]int)
	for exit>0 {
		fmt.Println("Enter the Name ")
		var s string 
        fmt.Scanln(&s)
        var num int
		fmt.Println("Enter the Contact number of ",s)
		fmt.Scanln(&num)
		m[s]=num
		exit--;
        
	}
	for key :=range m{
		fmt.Println(key," -> ",m[key])
	}

}