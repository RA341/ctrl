package system

import (
	"fmt"
	"os/exec"
	"runtime"
)

// ExecShell commands should be formated like [command, arg1, arg2,...]
func ExecShell(cmds []string) {
	cmd := exec.Command(cmds[0], cmds[1:]...)
	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running", cmd, err)
	}
}

func GetOS() string {
	return runtime.GOOS
}
