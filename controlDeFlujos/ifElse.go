package main

import (
	"fmt"
	"time"
)

func main() {
	if t := time.Now(); t.Hour() < 12 {
		fmt.Printf("mañana")
	} else if t.Hour() > 17 {
		fmt.Printf("media tarde")
	} else {
		fmt.Printf("ya es de noche")
	}
}
