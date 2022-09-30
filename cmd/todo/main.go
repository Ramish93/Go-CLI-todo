package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/Ramish93/Go-CLI-todo"
)

const todoFile = ".todos.json"


func main() {
	add := flag.Bool("add",false,"add a new todo")
	complete := flag.Int("complete", 0, "mark a  todo as completed")
	del := flag.Int("del", 0, "Delete a todo")
	list:= flag.Bool("list", false, "list all todos")

	flag.Parse()

	todos := &todo.Todos{}
	
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add(task: "sample todo")
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Println(os.stderr, err.Error())
			os.Exit(1)
		}
	case complete > 0:
		err:= todos.Complete(*complete)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *del > 0 :
		err := todos.Delete(*del)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
		err:= todos.Store(todoFile)
		if err != nil {
			fmt.Println(os.stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.Print()
	default:
		fmt.Println(os.Stdout, "invalid command")
		os.Exit(0)
	}
}

func(t *Todos) Print(){

	for i,item := range *t {
		fmt.Printf(format: "%d: %s\n", i, item.Task)
	}
}