package main

import (
	"fmt"
	"time"
	"os/exec"

	"github.com/cameron-wags/xstatusbar/stat"
)

var battery = stat.New(`\d+%`, "acpi")
var volume = stat.New(`Volume:.*\s(\d+%)`, "pactl", "list", "sinks")
var mute = stat.New(`Mute:\s(yes|no)`, "pactl", "list", "sinks")

func main() {
	b := battery.Check()
	m := mute.Check()
	var v string
	if m == "yes" {
		v = "MUTE"
	} else {
		v = volume.Check()
	}

	Update(fmt.Sprintf("%s    Batt: %s    Vol: %s", timeFmt(), b, v))
}

// Update sends a string to xsetroot -name.
func Update(name string) {
	//TODO maybe care about the error
	_ = exec.Command("xsetroot", "-name" , name).Run()
}

func timeFmt() string {
	return time.Now().Format("2006-01-02 15:04")
}
