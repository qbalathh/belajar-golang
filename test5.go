package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args) // menampilkan argumen yang diberikan saat menjalankan program
	fmt.Println("Jumlah argumen:", len(args))
	fmt.Println("Argumen pertama:", args[0])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname :", name)
	} else {
		fmt.Println("error", err.Error())

		username := os.Getenv("APP_USER")     // mendapatkan nilai dari environment variable APP_USER
		password := os.Getenv("APP_PASSWORD") // mendapatkan nilai dari environment variable APP_PASSWORD
		fmt.Println("Username:", username)
		fmt.Println("Password:", password)

		var (
			host *string = flag.String("host", "localhost", "host server")      // mendeklarasikan flag host dengan default localhost
			user *string = flag.String("user", "root", "user database")         // mendeklarasikan flag user dengan default root
			pass *string = flag.String("pass", "password", "password database") // mendeklarasikan flag pass dengan default password
		)

		flag.Parse()
		// parsing flag yang diberikan saat menjalankan program
		fmt.Println("host:", *host) // menampilkan nilai dari flag host, user, dan pass
		fmt.Println("user:", *user)
		fmt.Println("pass:", *pass)

	}
}
