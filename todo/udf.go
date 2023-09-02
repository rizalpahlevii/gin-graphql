package todo

var todos = []Todo{}

var count = 1

func GetTodos() []Todo {
	return todos
}

func GetTodo(id int) Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return todo
		}
	}

	return Todo{}
}

func UpdateTodo(id int, title string, completed bool) bool {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = title
			todos[i].Completed = completed
			return true
		}
	}
	return false
}
func AddTodo(title string) *Todo {
	data := &Todo{
		ID:        count,
		Title:     title,
		Completed: false,
	}

	todos = append(todos, *data)
	count++
	return data
}

func DeleteTodo(id int) bool {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}
