package main
import "awesomeProject/Ninja_Level_12/Hands_On_exercise_1/dog"
import "fmt"

func main(){
var y int
fmt.Println("Enter the year")
fmt.Scan(&y)

fmt.Println("The corresponding Dog age is ", dog.Years(y))
}