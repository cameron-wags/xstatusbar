package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

// Cmd defines a single metric that can be checked with a command.
type Cmd struct {
	// Name is associated with the stat when formatted.
	Name string

	// Command gets results containing the stat.
	Command string

	// Args are supplied to the command.
	Args []string

	// Pattern is a Regular Expression which matches the desired stat.
	// Submatches are allowed but only one. The last submatch hit will be returned.
	Pattern string

	regex *regexp.Regexp
}

// New makes a Cmd Stat with the settings provided.
func New(name string, pattern string, command string, args ...string) *Cmd {
	return &Cmd{
		Name:    name,
		Command: command,
		Args:    args,
		Pattern: pattern,
	}
}

// NewBuild makes a Cmd Stat like cmd.New() and compiles its regular expression.
func NewBuild(name string, pattern string, command string, args ...string) *Cmd {
	s := NewBuild(name, pattern, command, args...)
	s.regex = regexp.MustCompile(s.Pattern)
	return s
}

// Title returns the title associated with a stat.
func (s *Cmd) Title() string {
	return s.Name
}

// Check returns the string value of a stat.
func (s *Cmd) Check() string {
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
