package main

import (
	"os/exec"
	"strings"

	"github.com/cameron-wags/xstatusbar/component"
	"github.com/cameron-wags/xstatusbar/stat"
	"github.com/cameron-wags/xstatusbar/stat/cmd"
)

var padding = "    "

// Statistics in the status bar in display order.
var components = []stat.Statiface{
	component.NewClock("2006-01-02 15:04"),
	cmd.New("Batt", `Battery.*\s(\d+%)`, "acpi"),
	component.NewBrightness(),
	component.NewVolume(),
	cmd.New("WiFi", `(connected|connecting|disconnected)`, "nmcli", "g"),
}

func main() {
	statusBar := strings.Builder{}
	stopPadding := len(components) - 1
	for index, part := range components {
		// Ignore errors because we're useless..
		statusBar.WriteString(stat.Format(part))
		if index != stopPadding {
			statusBar.WriteString(padding)
		}
	}
	Update(statusBar.String())
}

// Update sends a string to xsetroot -name.
func Update(name string) {
	//TODO maybe care about the error
	_ = exec.Command("xsetroot", "-name", name).Run()
}
