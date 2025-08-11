package helper

import "fmt"

func SayHello(name string) string {
	fmt.Println("Hello", name)
	return "Hello " + name
}
