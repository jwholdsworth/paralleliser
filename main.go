package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	command := flag.String("c", "", "The command to run in each subdirectory")
	flag.Parse()

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
		go runCommandInFolder(*command, entry.Name())
	}
	wg.Wait()
}

func runCommandInFolder(command string, folder string) {
	var out bytes.Buffer

	cmd := exec.Command(os.Getenv("SHELL"), "-c", command)
	cmd.Stdout = &out
	cmd.Dir = fmt.Sprintf("./%s", folder)
	err := cmd.Run()
	if err != nil {
		log.Panic(err)
	}

	var output = fmt.Sprintf("Output from %s\n", folder)
	output += fmt.Sprintf(out.String())

	fmt.Println(output)
	wg.Done()
}
