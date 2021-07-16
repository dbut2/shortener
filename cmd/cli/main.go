package main

func main() {
	cmd := Command()
	err := cmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
