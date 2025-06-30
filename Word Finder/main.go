package main

import "sync"

func main() {
	logs := Logs{}
	commands := NewFlags()
	commands.Execute(&logs)

	var wg sync.WaitGroup
	wg.Done()

	

}
