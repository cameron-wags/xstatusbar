package component

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const backlightPath = "/sys/class/backlight"
const backlightDevice = "intel_backlight"

var maxBrightness = fmt.Sprintf("%s/%s/max_brightness", backlightPath, backlightDevice)
var actualBrightness = fmt.Sprintf("%s/%s/actual_brightness", backlightPath, backlightDevice)

// Brightness tracks backlight brightness
type Brightness struct{}

// NewBrightness returns a Brightness stat.
func NewBrightness() *Brightness {
	return &Brightness{}
}

// Title implements statiface
func (b *Brightness) Title() string {
	return "Backlight"
}

// Check implements statiface
func (b *Brightness) Check() string {
	return fmt.Sprintf("%s%%", percentage())
}

func percentage() string {
	current := readFileInt(actualBrightness)
	max := readFileInt(maxBrightness)
	v := float64(current) * 100 / float64(max)
	return strconv.FormatFloat(v, 'f', 0, 64)
}

func readFileInt(path string) int {
	//TODO go 1.16 os.ReadFile
	bytes, _ := ioutil.ReadFile(path)
	value, _ := strconv.Atoi(strings.TrimSpace(string(bytes)))
	return value
}
