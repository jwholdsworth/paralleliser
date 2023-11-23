package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	command := os.Getenv("COMMAND")
	if command == "" {
		log.Fatal("COMMAND environment variable not specified")
	}

	entries, err := os.ReadDir("./")

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		log.Println("Running in folder", entry.Name())
		wg.Add(1)
		go runCommandInFolder(command, entry.Name())
	}
	wg.Wait()
}

func runCommandInFolder(command string, folder string) {
	var out bytes.Buffer

	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = &out
	cmd.Dir = fmt.Sprintf("./%s", folder)
	err := cmd.Run()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Output from", folder)
	fmt.Println(out.String())
	wg.Done()
}
