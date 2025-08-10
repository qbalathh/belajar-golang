package helper

import "fmt"

func sayHello(name string) string {
	fmt.Println("Hello", name)
	return "Hello " + name
}
