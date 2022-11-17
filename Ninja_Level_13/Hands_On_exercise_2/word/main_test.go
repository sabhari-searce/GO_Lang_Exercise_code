package main



import (
	
	"testing"
	"awesomeProject/Ninja_Level_13/Hands_On_exercise_2/quote"
	"awesomeProject/Ninja_Level_13/Hands_On_exercise_2/word"
	
)

func ExampleQuote(){
	word.Count(quote.SunAlso)
	//output : 1349
}



func BenchmarkYear(b *testing.B){
	for i := 0; i < b.N; i++{
		word.Count(quote.SunAlso)
	}
}



func TestYears(t *testing.T){
	v := word.Count(quote.SunAlso)
	if v != 1349{
		t.Error("Expected 1349 but got ",v)
	}
}

