package main

import (
	"fmt"
	"os"
)

var tasks []string

func main() {
	for {
		fmt.Println("CLI To-Do App")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Task")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		switch choice {
		case 1:
			addTask()
		case 2:
			listTasks()
		case 3:
			deleteTask()
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addTask() {
	var task string
	fmt.Print("Enter the task: ")
	fmt.Scanln(&task)
	tasks = append(tasks, task)
	fmt.Println("Task added!")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show.")
	} else {
		fmt.Println("Tasks:")
		for i, task := range tasks {

			fmt.Printf("%d. %s\n", i+1, task)
		}
	}
}

func deleteTask() {
	var taskNumber int
	listTasks()
	fmt.Print("Enter the task number to delete: ")
	_, err := fmt.Scanln(&taskNumber)
	if err != nil || taskNumber < 1 || taskNumber > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)
	fmt.Println("Task deleted!")
}
