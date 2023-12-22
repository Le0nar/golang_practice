package main

import "fmt"

type Human struct {
	Age  int
	Name string
}

type Action struct {
	Human Human
}

func (a *Action) sendMessage(mes string) {
	fullMes := a.Human.Name + ": " + mes

	fmt.Printf("%s\n", fullMes)
}

func main () {
	user := Human{Age: 16, Name: "Naruto"}
	localAction := Action{Human: user}

	localAction.sendMessage("SASUKE")
}
