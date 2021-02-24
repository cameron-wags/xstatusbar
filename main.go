package main

import (
	"fmt"
	"os/exec"

	"github.com/cameron-wags/xstatusbar/component"
	"github.com/cameron-wags/xstatusbar/stat"
	"github.com/cameron-wags/xstatusbar/stat/cmd"
)

var clock = component.NewClock("2006-01-02 15:04")
var battery = cmd.New("Batt", `\d+%`, "acpi")
var backlight = component.NewBrightness()
var volume = cmd.New("Vol", `Playback.*\[(\d+%)\]`, "amixer", "get", "Master")
var mute = cmd.New("Vol", `\[(on|off)\]`, "amixer", "get", "Master")

// Example
// var statList = []stat.Statiface{
// 	stat.NewCmd("Batt", `\d+%`, "acpi"),
// }

func main() {
	//TODO unhack volume with a composite stat
	m := mute.Check()
	var v string
	if m == "off" {
		v = "MUTE"
	} else {
		v = volume.Check()
	}

	Update(fmt.Sprintf("%s    %s    %s    Vol: %s",
		stat.Format(clock),
		stat.Format(battery),
		stat.Format(backlight),
		v))
}

// Update sends a string to xsetroot -name.
func Update(name string) {
	//TODO maybe care about the error
	_ = exec.Command("xsetroot", "-name", name).Run()
}
