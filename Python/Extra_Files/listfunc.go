package main

import (
"fmt"
"sort"
)

func main(){

var markList=[] int {78,90,100,96,83,95}
fmt.Println("The marks are:",markList)

markList=append(markList,55)
fmt.Println("After adding new marks:",markList)

var i int=2
markList = append(markList[:i], markList[i+1:]...)
fmt.Println("After removing the marks at the 2nd position:",markList)

sort.Ints(markList)
fmt.Println("After sorting the marks in the given list:",markList)
}