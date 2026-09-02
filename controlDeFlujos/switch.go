package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS

	switch os {
	case "windows":
		fmt.Println("go run in -----> %s", os)
	case "linux":
		fmt.Println("go run on linux ")

	case "darwin":
		fmt.Println("go run in mac")

	default:
		fmt.Println("go run in other os")

	}
}
