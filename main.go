package main

import (
	"fmt"
	"github.com/VishalJuneja23/demyst/todo"
	"strconv"
)

func main() {

	todos := todo.FetchEvenTodos(20)
	for _, todo := range todos {
		fmt.Println("Title : " + todo.Title + " " + "Status : " + strconv.FormatBool(todo.Completed))
	}

}
