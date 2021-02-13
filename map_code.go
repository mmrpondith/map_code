package main
import "fmt"
func main(){
//key=value
//x:=make (map [string]string)
//x["name"]="rahul"
//x["age"]="30"
//x["hight"]="5.5"
//delete(x,"age")
//fmt.Println(x)
var myMap map [string]int
myMap=map[string]int{}
myMap["a"]=12
myMap["b"]=120
myMap["c"]=1200
myMap["d"]=12000
myMap["e"]=120000
fmt.Println(myMap)
delete(myMap,"e")
}