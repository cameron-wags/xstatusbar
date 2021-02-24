package stat

import "fmt"

// Statiface is a single-point statistic.
type Statiface interface {
	Title() string
	Check() string
}

// Format returns a formatted representation of a Statiface.
func Format(s Statiface) string {
	if s.Title() == "" {
		return s.Check()
	}
	return fmt.Sprintf("%s: %s", s.Title(), s.Check())
}
