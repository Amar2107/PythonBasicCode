package main

import (
"fmt"
"strings"
)

func main(){
var boardingCall string="Good Evening, this is the final call to AI passengers for the flight AI 466 which is planned to take off at 8.40A.M."
if(strings.HasPrefix(boardingCall,"Good Evening")){
	fmt.Println(strings.Replace(boardingCall,"Good Evening","Good Morning",-1))
    }
if(strings.Index(boardingCall,"AI"))>=0{
  fmt.Println("Welcome to Air India.")
  }
if(strings.HasSuffix(boardingCall,"A.M.")){
fmt.Println("Passengers are requested to have their breakfast.")
}
a:=strings.Split(boardingCall," ")
fmt.Println("Words in the boarding call message are:")
for i:=0;i< len(a);i++ {
  fmt.Println(a[i])
}
fmt.Println("Total number of times flight service name is specified in boarding call:",strings.Count(boardingCall,"AI"))
var message string="Thank you all..Have a nice journey!"
fmt.Println(strings.ToUpper(message))
fmt.Println(strings.ToLower(message))
}