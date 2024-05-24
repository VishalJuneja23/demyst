package todo

import (
	"reflect"
	"sync"
	"testing"
)

func TestFetchEvenTodos_GivenNoErrors_ShouldAddTodosToSlice(t *testing.T) {
	originalCollectTodo := collectTodo
	collectTodo = func(id int, todos *[]Todo, wg *sync.WaitGroup, mutex *sync.Mutex) {
		defer wg.Done()
		todo := Todo{
			Id:        2,
			Title:     "Test Todo",
			Completed: false,
			UserId:    43,
		}
		mutex.Lock()
		defer mutex.Unlock()
		*todos = append(*todos, todo)
	}
	defer func() { collectTodo = originalCollectTodo }()

	expected := []Todo{
		{Id: 2, Title: "Test Todo", Completed: false, UserId: 43},
	}

	actual := FetchEvenTodos(1)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected todos %v, got %v", expected, actual)
	}

}

func TestFetchEvenTodos_GivenErrors_ShouldNotAddTodosToSlice(t *testing.T) {
	originalCollectTodo := collectTodo
	collectTodo = func(id int, todos *[]Todo, wg *sync.WaitGroup, mutex *sync.Mutex) {
		defer wg.Done()
		return
	}
	defer func() { collectTodo = originalCollectTodo }()

	var expected []Todo

	actual := FetchEvenTodos(1)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected todos %v, got %v", expected, actual)
	}
}
