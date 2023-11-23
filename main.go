package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	entries, err := os.ReadDir("./")

	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		fmt.Println(e)
	}
}