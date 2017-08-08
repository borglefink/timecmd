package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("timecmd: Displays start time, end time and duration for a given command.")
		fmt.Println("Usage  : timecmd [command] [optional parameters to command]")
		os.Exit(0)
	}

	var cmd = exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var startTime = time.Now()
	var err = cmd.Run()
	var endTime = time.Now()

	if err != nil {
		fmt.Printf("===================================================\n")
		fmt.Printf("Error running command %s, %v\n", cmd.Path, err)
	}

	fmt.Printf("===================================================\n")
	fmt.Printf("Start time: %v\n", startTime)
	fmt.Printf("End time  : %v\n", endTime)
	fmt.Printf("Duration  : %v\n", endTime.Sub(startTime))
}
