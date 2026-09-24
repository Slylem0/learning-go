package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		checkAPI(api)
	}

	alpased := time.Since(start)
	fmt.Printf("\n Listo tomo %v segundos \n", alpased.Seconds())
}

func checkAPI(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Printf("ERROR %s the api is fallen \n ", api)
	}

	fmt.Printf("SUCESS la API esta funcionando %s \n ", api)
}
