package commands

import (
	"os/exec"
	"sync"
	"time"
)

// Command wraps os/exec.Cmd with a timeout, and arg parsing.
type Command struct {
	Name    string
	Cmd     *exec.Cmd
	Exec    string
	Args    []string
	Timeout time.Duration
	lock    *sync.Mutex
}
