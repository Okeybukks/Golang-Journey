package main

func main() {
	logs := Logs{}
	commands := NewFlags()
	commands.Execute(&logs)
}
