package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mzaid-dev/task-tracker-cli.git/services"
)

func Execute() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	service := services.TaskService{
		FileName: "tasks.json",
	}

	command := os.Args[1]

	switch command {

	case "add":
		handleAdd(service)

	case "list":
		handleList(service)

	case "update":
		handleUpdate(service)

	case "delete":
		handleDelete(service)

	default:
		fmt.Println("Unknown command:", command)
		printUsage()

	}
}

func handleAdd(service services.TaskService) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task add \"description\"")
		return
	}

	description := os.Args[2]

	err := service.AddTask(description)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task added successfully")
}

func handleList(service services.TaskService) {
	status := ""
	if len(os.Args) > 2 {
		status = os.Args[2]
	}

	tasks, err := service.ListTasks(status)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, task := range tasks {
		fmt.Printf(
			"[%d] %s | %s | %s\n",
			task.ID,
			task.Status,
			task.Description,
			task.CreatedAt.Format("2006-01-02"),
		)
	}
}

func handleUpdate(service services.TaskService) {
	if len(os.Args) < 4 {
		fmt.Println("Usage: task update <id> <status>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	status := os.Args[3]

	err = service.UpdateTaskStatus(id, status)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task updated successfully")
}

func handleDelete(service services.TaskService) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task delete <id>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	err = service.DeleteTask(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task deleted successfully")
}

func printUsage() {
	fmt.Println("Task Tracker CLI")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  add \"description\"     Add a new task")
	fmt.Println("  list [status]          List tasks")
	fmt.Println("  update <id> <status>   Update task status")
	fmt.Println("  delete <id>            Delete a task")
}
