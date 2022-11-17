package main



import (
	
	"testing"
	"awesomeProject/Ninja_Level_13/Hands_On_exercise_3/mymath"
	
)

func ExampleQuote(){
	mymath.CenteredAvg([]int{1, 4, 6, 8, 100})
	//output : 6
}



func BenchmarkYear(b *testing.B){
	for i := 0; i < b.N; i++{
		mymath.CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}



func TestYears(t *testing.T){
	v := mymath.CenteredAvg([]int{1, 4, 6, 8, 100})
	if v != 6{
		t.Error("Expected 6 but got ",v)
	}
}

