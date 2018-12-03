package main
import "fmt"

func main() {
var x []int
var l,n int
  fmt.Println("Plerase enter the size of array: ")
  fmt.Scanln(&l)
  fmt.Println("Enter the values: ")
  for i:=0; i < l; i++ {
      var temp int
      fmt.Scanln(&temp)
      x = append(x, temp)
  }
  var largest int
  for i:=0;i < l; i++{
    if x[i] > largest {
      n = x[i]
      x[i] = largest
      largest = n
    }
  }
fmt.Println("The largest number is:", largest)
}
