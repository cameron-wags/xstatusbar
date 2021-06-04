package component

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Wifi shows connection strength & ssid
type Wifi struct {
	title  string
	device string
}

// NewWifi creates a new Wifi stat.
func NewWifi(title, device string) *Wifi {
	return &Wifi{
		title:  title,
		device: device,
	}
}

// Title implements statiface.
func (w *Wifi) Title() string {
	return w.title
}

// Check implements statiface.
func (w *Wifi) Check() string {
	sigb, err := os.ReadFile("/proc/net/wireless")
	if err != nil {
		return "procerr"
	}
	sig := strings.Split(string(sigb), "\n")
	if len(sig) < 3 {
		return "disconnected"
	}

	var iface string
	var status, link int
	n, err := fmt.Sscanf(sig[2], "%s %d %d", &iface, &status, &link)
	if n < 2 {
		return "disconnected"
	}
	if err != nil {
		return "fmterr"
	}

	strength := (float64(link) / float64(70) * 100)
	strengthStr := strconv.FormatFloat(strength, 'f', 0, 64)

	return strengthStr + "%"
}
