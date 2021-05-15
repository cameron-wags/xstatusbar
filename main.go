package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/cameron-wags/xstatusbar/component"
	"github.com/cameron-wags/xstatusbar/stat"
	"github.com/cameron-wags/xstatusbar/stat/cmd"
)

var padding = "    "

// Statistics in the status bar in display order.
var components = []stat.Statiface{
	component.NewSleep("CAFFEINE"),
	component.NewClock("2006-01-02 15:04"),
	cmd.New("Bat", `Battery.*\s(\d+%)`, "acpi"),
	component.NewBrightness("Backlight"),
	component.NewVolume("Vol"),
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

	// -t flag prints remaining seconds before the clock time goes stale.
	// Sets up for a handy refresh script in ~/.xinitrc
	// Bash example:
	// while true: do
	//     sleepTime=$(xstatusbar -t)
	//     sleep $sleepTime
	// done &
	if len(os.Args) > 1 && os.Args[1] == "-t" {
		fmt.Fprintln(os.Stdout, 60-time.Now().Second())
	}
}

// Update sends a string to xsetroot -name.
func Update(name string) {
	//TODO maybe care about the error
	_ = exec.Command("xsetroot", "-name", name).Run()
}
