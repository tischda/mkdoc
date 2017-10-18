package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// Executes process and print elapsed time.
func executeProcess(command string, args ...string) {
	defer whenDone()()
	if noop {
		fmt.Printf("Execution skipped: %s with options: %s\n", command, append([]string{}, args...))
	} else {
		fmt.Printf("Running %s with options: %s\n", command, append([]string{}, args...))
		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		checkFatal(cmd.Run())
	}
}

// Callback function executed when process is done.
func whenDone() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Total time: %v\n", time.Since(start))
	}
}

// Prints error and exit if err != nil
func checkFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
