package main

import (
	cmd "Todo_Project/todo_package"
)

func main() {
	todos := cmd.Todos{}
	commands := cmd.NewCmdFlags()
	storage := cmd.NewStorage[cmd.Todos]("tasks.json")
	storage.Load(&todos)
	commands.Execute(&todos)
	storage.Save(&todos)
}
