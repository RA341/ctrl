package system

import (
	"fmt"
	"github.com/rs/zerolog/log"
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
		log.Info().Msg("Unsupported os, no action will be taken")
	}
}

func registerWindowsService() {
	log.Warn().Msg("Registering a service on windows is not yet support, no action will be taken")
}

func registerLinuxService() {
	if checkServiceLocation() {
		log.Debug().Msg("Service file exists")
		return
	}

	execLoc, cwd := getCWDAndExecLocations()
	if execLoc == "" || cwd == "" {
		log.Error().Msg("Error failed to find executable location or current working dir")
		return
	}

	serviceFile := createServiceFile(execLoc, cwd)
	if serviceFile == "" {
		log.Error().Msg("Could not create service file")
		return
	}

	mvServiceFile(serviceFile)
}

func checkServiceLocation() bool {
	_, err := os.Stat("/etc/systemd/system/ctrl.service")
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		log.Info().Msg("No service file found, registering...")
		return false
	} else {
		log.Error().Err(err).Msg("Error checking service file")
		return false
	}
}

func getCWDAndExecLocations() (string, string) {
	execLoc, err := os.Executable()
	if err != nil {
		log.Debug().Err(err).Msg("Could not find executable location")
		return "", ""
	}

	cwd := filepath.Dir(execLoc)

	_, err = os.Stat(execLoc)
	if err == nil {
		log.Debug().Msg("The file exists")
		return execLoc, cwd
	} else if os.IsNotExist(err) {
		return "", ""
	} else {
		log.Debug().Err(err).Msg("Error checking file: %v")
		return "", ""
	}
}

func mvServiceFile(servicePath string) {
	destination := "/etc/systemd/system/ctrl.service"

	err := os.Rename(servicePath, destination)
	if err != nil {
		log.Error().Err(err).Msg("Error moving service file to systemd")
		return
	}

	log.Debug().Msgf("File moved successfully from %s to %s\n", servicePath, destination)
	log.Info().Msg("Enable the service by running 'systemctl daemon-reload'")
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
		log.Error().Err(err).Msg("Error creating service file")
		return ""
	}

	log.Debug().Msg("Created service file")
	return cwd + "/" + serviceFile
}
