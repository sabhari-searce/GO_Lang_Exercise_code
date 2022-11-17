package main

import "testing"
import "awesomeProject/Ninja_Level_13/Hands_On_exercise_1/dog"
import "fmt"

func ExampleYear(){
	fmt.Println(dog.Years(5))
	//output : 35
}

func ExampleYearTwo(){
	fmt.Println(dog.YearsTwo(5))
	//output : 35
}

func BenchmarkYear(b *testing.B){
	for i := 0; i < b.N; i++{
		dog.Years(5)
	}
}

func BenchmarkYearTwo(b *testing.B){
	for i := 0; i < b.N; i++{
		dog.YearsTwo(5)
	}
}

func TestYears(t *testing.T){
	v := dog.Years(5)
	if v != 35{
		t.Error("Expected 35 but got ",v)
	}
}

func TestYearsTwo(t *testing.T){
	v := dog.YearsTwo(5)
	if v!= 35{
		t.Error("Expected 35 but got ",v)
	}

}