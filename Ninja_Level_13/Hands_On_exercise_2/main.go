package main

import (
	"fmt"
	"awesomeProject/Ninja_Level_13/Hands_On_exercise_2/quote"
	"awesomeProject/Ninja_Level_13/Hands_On_exercise_2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}