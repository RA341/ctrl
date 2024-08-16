package system

func ExecShutDown() {
	// Use the "shutdown" command to shut down the computer immediately
	os := GetOS()
	var cmds []string
	if os == "windows" {
		cmds = []string{"psshutdown", "-s", " -t", "0"}
	} else {
		cmds = []string{"sudo", "shutdown", "now"}
	}
	ExecShell(cmds)
}

func ExecReboot() {
	os := GetOS()
	var cmds []string
	if os == "windows" {
		cmds = []string{"psshutdown", "-r", " -t", "0"}
	} else {
		cmds = []string{"sudo", "reboot", "now"}
	}
	ExecShell(cmds)
}

func ExecSleep() {
	os := GetOS()
	var cmds []string
	if os == "windows" {
		cmds = []string{"psshutdown", "-d", " -t", "0"}
	} else {
		cmds = []string{"sudo", "sleep", "now"}
	}
	ExecShell(cmds)

}
