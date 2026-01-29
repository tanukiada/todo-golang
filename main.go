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

func newTodo(text string, completed bool) Todo {
	return Todo {
		ID: uuid.New().String(),
		Text: text,
		Completed: completed,
	}
}

func main() {
	var i String
	fmt.Println("What would you like to do today?")
	fmt.Scanln(&i)
	t1 := Todo
	fmt.Println("You input " + todo.Text)
}
