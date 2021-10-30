package main

import "fmt"

func main()  {
   fmt.Println("enter the size")
   var n int
   fmt.Scanln(&n)
   slice:=make([]int,n)
   fmt.Println("Enter the elements of the array")
   for i:=0;i<n;i++ {
      fmt.Scanln(&slice[i])
   }
   doubleTheNumbers(&slice,n);
 length:=  display(slice)
 fmt.Println(length)
 fmt.Println("calling recursion in the next step")
 rec(n);
 fmt.Println("Implementing 2d array in go lang")
 var arr[2][2] int
 fmt.Println("enter the data of the 2d matrix 2 *2")
 for i:=0;i<2;i++{
    for j:=0;j<2;j++{
       fmt.Scanln(&arr[i][j])
    }
 }
 fmt.Println("output of the data entered in the matrix")
 for i:=0;i<2;i++{
    fmt.Println("\nrow ",i,)
   for j:=0;j<2;j++{
      fmt.Print(arr[i][j],"  ")
   }

}
}
func doubleTheNumbers(slice *[] int,n int){
   for i:=0;i<len(*slice);i++ {
      (*slice)[i]=(*slice)[i]*2
   }
}
func display(slice [] int ) int {
   for i:=0;i<len(slice);i++ {
      fmt.Println(slice[i])
   }
   fmt.Println(".")
return len(slice)
}
func rec(n int){
   if n==0{
      return 
   }
   fmt.Println(n)
   rec(n-1)
   fmt.Println(n)
}