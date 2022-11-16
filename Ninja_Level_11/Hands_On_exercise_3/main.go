package main


import "fmt"

type customErr struct{
	msg string
}

func (ce customErr)Error() string{
	return fmt.Sprintln(ce.msg)
}

func foo(e error){
	fmt.Println("Error found!!",e)
}

func main(){
ce := customErr{
	msg : "Invalid argument",
}
foo(ce)
}



