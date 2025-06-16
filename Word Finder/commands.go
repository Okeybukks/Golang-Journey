package main

import (
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	Path     string
	Keyword  string
	FromTime string
	ToTime   string
	Print    bool
	Head     int
	Tail     int
}

func NewFlags() *Flags {
	fl := Flags{}

	flag.StringVar(&fl.Path, "filepath", "", "filepath(*Required*) => used to specify the filepath of the logs file. \n\tUsage: go run . --filepath \"logs.txt\"")
	flag.StringVar(&fl.Keyword, "keyword", "", "keyword(*Required*) => used to specify the keyword to be searched in the log file. \n\tUsage: go run . --filepath \"logs.txt\" --keyword \"Error\"")
	flag.BoolVar(&fl.Print, "print", false, "print(*Required*) => used to print out the result of the search. \n\tUsage: go run . --filepath \"logs.txt\" --keyword \"Error\" --print")
	flag.StringVar(&fl.FromTime, "from", "", "from => used to specify the start time of the searchntimestamp in the log file. Time should be in this format \"2025-06-14 16:00:14\"  \n\tUsage: --from \"2025-06-14 14:00:14\"")
	flag.StringVar(&fl.ToTime, "to", "", "to => used to specify the end time of the search timestamp in the log file. Time should be in this format \"2025-06-14 16:00:14\"  \n\tUsage: --to \"2025-06-14 16:00:14\"")
	flag.IntVar(&fl.Head, "head", -1, "head => display first n lines of filtered logs.  \n\tUsage: --head 5")
	flag.IntVar(&fl.Tail, "tail", -1, "tail => display last n lines of a filtered logs.  \n\tUsage: --tail 5")
	flag.Parse()

	return &fl
}

func (fl *Flags) Execute(logs *Logs) {
	if fl.Path == "" {
		fmt.Println("Error: --filepath is required")
		os.Exit(1)
	}
	if fl.Keyword == "" {
		fmt.Println("Error: --keyword is required")
		os.Exit(1)
	}
	if !fl.Print {
		fmt.Println("Error: --print is required")
		os.Exit(1)
	}
	logs.fileScanner(fl.Path, fl.Keyword)
	if fl.Print {
		logs.Print(fl.FromTime, fl.ToTime, fl.Head, fl.Tail)
	}

}
