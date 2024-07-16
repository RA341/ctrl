package system

import (
	"github.com/rs/zerolog/log"
	"os/exec"
	"runtime"
)

// ExecShell commands should be formated like [command, arg1, arg2,...]
func ExecShell(cmds []string) {
	cmd := exec.Command(cmds[0], cmds[1:]...)
	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Error().Err(err).Msgf("Error running: %s", cmd)
	}
}

func GetOS() string {
	return runtime.GOOS
}
