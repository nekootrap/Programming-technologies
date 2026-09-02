package main

import "fmt"

// greet возвращает приветствие для указанного имени.
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	var userName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&userName)

	message := greet(userName)
	fmt.Println(message)
}
