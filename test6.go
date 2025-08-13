package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	NamaSaya := "Muhammad Iqbal Athhar"

	fmt.Println(
		strings.Contains("joko widodo", "o"),
		strings.Split(NamaSaya, " "),
		strings.ToLower(NamaSaya),
		strings.ToUpper(NamaSaya),
		strings.Trim("Msaya suka makanM", "M"),
	)

	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("error", err.Error())
	}
	fmt.Println(strconv.ParseInt("205003", 10, 64))
	fmt.Println(
		strconv.FormatInt(2025, 10),
		strconv.FormatInt(2025, 2),
		strconv.FormatInt(2025, 8),
		strconv.FormatInt(2025, 16),
		strconv.Itoa(1000),
	)
}
