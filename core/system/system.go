package system

func ExecShutDown() error {
	// Use the "shutdown" command to shut down the computer immediately
	os := GetOS()
	var cmds []string
	if os == "windows" {
		cmds = []string{"psshutdown", "-s", " -t", "0"}
	} else {
		cmds = []string{"sudo", "shutdown", "now"}
	}

	err := ExecShell(cmds)
	if err != nil {
		return err
	}
	return nil
}

func ExecReboot() error {
	os := GetOS()
	var cmds []string
	if os == "windows" {
		cmds = []string{"psshutdown", "-r", " -t", "0"}
	} else {
		cmds = []string{"sudo", "reboot", "now"}
	}

	err := ExecShell(cmds)
	if err != nil {
		return err
	}
	return nil
}

func ExecSleep() error {
	os := GetOS()
	var cmds []string
	if os == "windows" {
		cmds = []string{"psshutdown", "-d", " -t", "0"}
	} else {
		cmds = []string{"sudo", "sleep", "now"}
	}

	err := ExecShell(cmds)
	if err != nil {
		return err
	}

	return nil
}
