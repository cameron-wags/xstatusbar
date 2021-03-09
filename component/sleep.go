package component

import (
	"github.com/cameron-wags/xstatusbar/stat/cmd"
)

// Sleep shows if the system will lock and blank screen automatically..
type Sleep struct {
	sleepStat     cmd.Cmd
	sleepDisabled string
}

// NewSleep creates a new auto sleep indicator.
func NewSleep(sleepDisabledCaption string) *Sleep {
	return &Sleep{
		sleepStat:     *cmd.New("", `(\w+)`, "xssstate", "-s"),
		sleepDisabled: sleepDisabledCaption,
	}
}

// Title implements statiface.
func (s *Sleep) Title() string {
	return ""
}

// Check implements statiface.
func (s *Sleep) Check() (value string) {
	status := s.sleepStat.Check()
	if status == "disabled" {
		value = s.sleepDisabled
	}
	return
}
