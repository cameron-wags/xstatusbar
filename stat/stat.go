package stat

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

// Stat defines
type Stat struct {
	// Command gets results containing the stat.
	Command string

	// Args are supplied to the command.
	Args []string

	// Pattern is a Regular Expression which matches the desired stat.
	// Submatches are allowed but only one. The last submatch hit will be returned.
	Pattern string

	// regex is set automatically.
	regex *regexp.Regexp
}

// New makes a Stat with the settings provided.
func New(pattern string, command string, args ...string) Stat {
	return Stat{
		Command: command,
		Args:    args,
		Pattern: pattern,
	}
}

// NewBuild makes a Stat like stat.New() and compiles its regular expression.
func NewBuild(pattern string, command string, args ...string) Stat {
	s := New(pattern, command, args...)
	s.regex = regexp.MustCompile(s.Pattern)
	return s
}

// Check returns the string value of a stat.
func (s *Stat) Check() string {
	if s.regex == nil {
		s.regex = regexp.MustCompile(s.Pattern)
	}
	out, err := exec.Command(s.Command, s.Args...).Output()
	if err != nil {
		fmt.Println(err)
		// TODO maybe handle
	}
	// kind of a hack.
	result := s.regex.FindSubmatch(out)
	if len(os.Args) > 1 && os.Args[1] == "-d" {
		fmt.Printf("Matches for: %s\n", s.Pattern)
		for _, r := range result {
			fmt.Println(string(r))
		}
	}
	return string(result[len(result)-1])
}
