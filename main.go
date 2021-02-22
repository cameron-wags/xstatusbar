package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Printf("%s\n%s\n", Volume(), Battery())
}

// Battery returns the current battery percentage.
func Battery() string {
	batt := cmdRun("acpi")
	//TODO regex filter
	return batt
}

// Volume returns the current volume.
// There might be an issue with discerning which output is active.
func Volume() string {
	volStr := cmdRun("pactl", "list", "sinks")
	//TODO regex filter

	return volStr
}

func cmdRun(bin string, arg ...string) string {
	out, err := exec.Command(bin, arg...).Output()
	if err != nil {
		fmt.Println(err)
		// TODO maybe handle
	}

	return string(out)
}
