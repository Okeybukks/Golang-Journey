package todo_package

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Edit   string
	Delete int
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "This flag is used to add your task to the Todo. \n\tUsage: go run . --add \"Task Title\"")
	flag.StringVar(&cf.Edit, "edit", "", "This flag is used to edit an uncompleted task in your Todo by specifying the task id and new title. \n\tUsage: go run . --edit \"task_id:New Task Title\"")
	flag.IntVar(&cf.Delete, "delete", -1, "This flag is used to delete a task. Specify the task number to delete. \n\tUsage: go run . --delete task_id")
	flag.IntVar(&cf.Toggle, "toggle", -1, "This flag is used to toggle a task as either completed or not. Specify the task number to toggle. \n\tUsage: go run . --toggle task_id")
	flag.BoolVar(&cf.List, "list", false, "This flag is used to print out the tasks in the Todo.")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		todo := strings.SplitN(cf.Edit, ":", 2)
		if len(todo) != 2 {
			fmt.Println("Invalid edit parameter")
			os.Exit(1)
		}
		index, err := strconv.Atoi(todo[0])
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		title := todo[1]
		todos.Edit(index, title)
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.List:
		todos.Print()
	default:
		fmt.Println("Invalid command specified")
	}
}
