package main

import "fmt"

type customErr struct {
	s string
}

func (c customErr) Error() string {
	return fmt.Sprintln(c.s)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	myError := customErr{"Erro de teste"}
	foo(myError)
}
