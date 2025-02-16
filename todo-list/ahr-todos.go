package main

import (
	"fmt"
)

// Define a struct for the todo
type Todo struct {
	ID          int
	Title       string
	Description string
}

// Define a slice of todos
var todos []Todo

// Define a function to print the menu
func menu() {
	fmt.Println("############### Welcome to the ***AHR*** Todos App ###############")
	fmt.Println("1. Create a new todo")
	fmt.Println("2. List all todos")
	fmt.Println("3. Update a todo")
	fmt.Println("4. Delete a todo")
	fmt.Println("0. Exit")

}

// Define a function to take the menu option from the user
func takeMenuOption() int {
	var option int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&option)
	return option
}

// Define a function to create a new todo
func createTodo() {
	fmt.Println("-------------------------------------")
	fmt.Println("Create a new todo")
	fmt.Println("-------------------------------------")
	var todo Todo
	fmt.Print("Enter the title of the todo: ")
	fmt.Scanln(&todo.Title)
	fmt.Print("Enter the description of the todo: ")
	fmt.Scanln(&todo.Description)
	todo.ID = getUniqueID()
	todos = append(todos, todo)
	fmt.Print("\033[H\033[2J")
	fmt.Println("Todo created successfully")
	listTodos()
}

// Define a function to get a unique ID for the todo
func getUniqueID() int {
	if len(todos) == 0 {
		return 1
	}
	return todos[len(todos)-1].ID + 1
}

// Define a function to list all todos
func listTodos() {

	if len(todos) == 0 {
		fmt.Println("No todos to display")
		return
	}

	fmt.Print("\033[H\033[2J")
	fmt.Println("-------------------------------------")
	fmt.Println("List all todos")
	fmt.Println("-------------------------------------")
	for _, todo := range todos {
		fmt.Printf("ID: %d, Title: %s, Description: %s\n", todo.ID, todo.Title, todo.Description)
	}
	fmt.Println("-------------------------------------")
}

// Define a function to update a todo
func updateTodo() {
	fmt.Println("-------------------------------------")
	fmt.Println("Update a todo")
	fmt.Println("-------------------------------------")

	var id int
	fmt.Print("Enter the ID of the todo to update: ")
	fmt.Scanln(&id)

	for i, todo := range todos {
		if todo.ID == id {
			fmt.Printf("Enter the new title of the todo (old title: %s): ", todo.Title)
			fmt.Scanln(&todos[i].Title)
			fmt.Printf("Enter the new description of the todo (old description: %s): ", todo.Description)
			fmt.Scanln(&todos[i].Description)
			fmt.Print("\033[H\033[2J")
			fmt.Println("Todo updated successfully")
			listTodos()
			return
		}
	}

	fmt.Println("Todo not found")
}

// Define a function to delete a todo
func deleteTodo() {
	fmt.Println("-------------------------------------")
	fmt.Println("Delete a todo")
	fmt.Println("-------------------------------------")

	var id int
	fmt.Print("Enter the ID of the todo to delete: ")
	fmt.Scanln(&id)

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Print("\033[H\033[2J")
			fmt.Println("Todo deleted successfully")
			listTodos()
			return
		}
	}

	fmt.Println("Todo not found")
}

// Define a function to exit the application
func exit() {
	fmt.Println("Exiting the application")
}

func app() {
	for {
		menu()
		option := takeMenuOption()
		switch option {
		case 1:
			createTodo()
			// fmt.Print("\033[H\033[2J")
			// menu()
		case 2:
			listTodos()
		case 3:
			updateTodo()
		case 4:
			deleteTodo()
		case 0:
			exit()
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func main() {
	app()
}
