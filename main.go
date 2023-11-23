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

	if *command == "" {
		log.Fatal("You need to specify a command to run with the -c flag")
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
		go runCommandInFolder(*command, entry.Name())
	}
	wg.Wait()
}

func runCommandInFolder(command string, folder string) {
	var out bytes.Buffer
	var err bytes.Buffer

	cmd := exec.Command(os.Getenv("SHELL"), "-c", command)
	cmd.Stdout = &out
	cmd.Stderr = &err
	cmd.Dir = fmt.Sprintf("./%s", folder)
	error := cmd.Run()
	if error != nil {
		log.Fatalf("Error running in %s. Error was %s", folder, err.String())
	}

	var output string
	output += "--------------------------------------------------------------------------------\n"
	output += fmt.Sprintf("Output from %s\n", folder)
	output += out.String()
	output += err.String()
	output += "--------------------------------------------------------------------------------\n\n"

	fmt.Println(output)
	wg.Done()
}
