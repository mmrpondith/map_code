package main
import "fmt"
func main(){
//key=value
x:=make (map [string]string)
x["name"]="rahul"
x["age"]="30"
x["hight"]="5.5"
delete(x,"age")
fmt.Println(x)
}