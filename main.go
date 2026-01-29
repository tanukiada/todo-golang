package main

import (
		"fmt"
		"github.com/google/uuid"
	)

type Todo struct {
	ID string
	Text string
	Completed bool
}

func makeTodo(text string, completed bool) Todo {
	return Todo {
		ID: uuid.New().String(),
		Text: text,
		Completed: completed,
	}
}

func main() {
	var i string
	fmt.Println("What would you like to do today?")
	fmt.Scanln(&i)
	t1 := makeTodo(i, false)
	fmt.Println("You input " + t1.Text)
}
