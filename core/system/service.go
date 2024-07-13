package system

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func RegisterService() {
	if runtime.GOOS == "linux" {
		registerLinuxService()
	} else if runtime.GOOS == "windows" {
		registerWindowsService()
	} else {
		fmt.Println("Unsupported os, no action will be taken")
	}
}

func registerWindowsService() {
	fmt.Println("Warning: registering a service on windows is not yet support, no action will be taken")
}

func registerLinuxService() {
	if checkServiceLocation() {
		fmt.Printf("Service file exists, no need to register")
		return
	}

	execLoc, cwd := getCWDAndExecLocations()
	if execLoc == "" || cwd == "" {
		log.Print("Error failed to find executable location or current working dir")
		return
	}

	serviceFile := createServiceFile(execLoc, cwd)
	if serviceFile == "" {
		log.Print("Error create service file")
		return
	}

	mvServiceFile(serviceFile)
}

func checkServiceLocation() bool {
	_, err := os.Stat("/etc/systemd/system/ctrl.service\n")
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		fmt.Println("No service file found, registering...")
		return false
	} else {
		fmt.Printf("Error checking file: %v\n", err)
		return false
	}
}

func getCWDAndExecLocations() (string, string) {
	execLoc, err := os.Executable()
	if err != nil {
		fmt.Print(fmt.Sprintf("Could not find executable location: %s", err))
		return "", ""
	}

	cwd := filepath.Dir(execLoc)

	_, err = os.Stat(execLoc)
	if err == nil {
		fmt.Printf("The file exists\n")
		return execLoc, cwd
	} else if os.IsNotExist(err) {
		return "", ""
	} else {
		fmt.Printf("Error checking file: %v\n", err)
		return "", ""
	}
}

func mvServiceFile(servicePath string) {
	destination := "/etc/systemd/system/ctrl.service"

	err := os.Rename(servicePath, destination)
	if err != nil {
		fmt.Printf("Error moving file: %v\n", err)
		return
	}

	fmt.Printf("File moved successfully from %s to %s\n", servicePath, destination)
}

func createServiceFile(execLoc string, cwd string) string {
	serviceFile := "ctrl.service"

	contents := []byte(fmt.Sprintf(`[Unit]
Description=ctrl
After=network.target

[Service]
User=root
WorkingDirectory=%s
ExecStart=%s
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
`, cwd, execLoc))

	err := os.WriteFile(serviceFile, contents, 0644)
	if err != nil {
		log.Print("Error creating service file")
		return ""
	}

	log.Print("Created service file")
	return cwd + "/" + serviceFile
}
