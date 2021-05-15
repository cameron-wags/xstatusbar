package component

import (
	"github.com/cameron-wags/xstatusbar/stat/cmd"
)

// Volume tracks the system volume.
type Volume struct {
	muteStat   cmd.Cmd
	volumeStat cmd.Cmd
	title      string
}

// NewVolume creates a new volume stat.
func NewVolume(title string) *Volume {
	return &Volume{
		// I'm still too lazy to make this one expression.
		volumeStat: *cmd.New("Vol", `Volume:.*\s(\d+%)`, "pactl", "list", "sinks"),
		muteStat:   *cmd.New("Vol", `Mute:\s+(yes|no)`, "pactl", "list", "sinks"),
		title:      title,
	}
}

// Title implements statiface.
func (v *Volume) Title() string {
	return v.title
}

// Check implements statiface.
func (v *Volume) Check() string {
	mute := v.muteStat.Check()
	var value string
	if mute == "yes" {
		value = "MUTE"
	} else {
		value = v.volumeStat.Check()
	}
	return value
}
