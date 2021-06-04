package main



//code from here was used as a template https://tutorialedge.net/golang/golang-mysql-tutorial/
//
//currently most things are hard coded because this was the last version to fully compile
//currently implementing  this with a webapp that responds with html pages containing the requested data
//
//webapp portion is being adapted with the result of this tutorial as the base https://golang.org/doc/articles/wiki/ 
//
//
import (
    "fmt"
)

   func main(){
   teststring1:=[]string{}
    teststring:=append(teststring1,"test")
    fmt.Println(teststring)
    teststring=append(teststring,nil)
   }