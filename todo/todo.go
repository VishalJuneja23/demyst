package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var collectTodo = func(id int, todos *[]Todo, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	baseURL := "https://jsonplaceholder.typicode.com/todos/%d"
	resp, err := http.Get(fmt.Sprintf(baseURL, id))
	if err != nil {
		fmt.Println(err)
		return

	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("error occured, status code is ", resp.StatusCode)
		return

	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return

	}
	mutex.Lock()
	defer mutex.Unlock()
	*todos = append(*todos, todo)

}

func FetchEvenTodos(count int) []Todo {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var todos []Todo
	for number, id := 0, 2; number < count; id += 2 {

		go collectTodo(id, &todos, &wg, &mutex)
		wg.Add(1)
		number++
	}
	wg.Wait()
	return todos
}
