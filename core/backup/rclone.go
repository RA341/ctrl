package backup

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"
)

var wg sync.WaitGroup
var fileList = []string{"movies", "shows", "nmov", "bmov", "music"}

func periodicRcloneTask() {
	for {
		fmt.Println("Running rclone tasks")

		syncMediaFiles()

		time.Sleep(1 * time.Hour)
	}
}

func syncMediaFiles() {
	// Start the goroutines for each task
	for _, value := range fileList {
		go syncDir(&wg, value)
		wg.Wait()
	}
}

func syncDir(wg *sync.WaitGroup, folder string) {
	// notify wg once command is completed
	file := "/mnt/pool/" + folder
	remote := "oneMedia:media/" + folder

	fmt.Println("Syncing " + file + " " + remote)

	runGoRoutineRclone(wg, file, remote)

	fmt.Println("Completed " + folder + " sync command")
}

func runGoRoutineRclone(wg *sync.WaitGroup, file string, remote string) {
	// notify wg once command is completed
	defer wg.Done()

	cmd := executeRcloneSyncCmd(file, remote)
	err := cmd.Run()
	if err != nil {
		log.Fatal("Failed to execute command " + err.Error())
		return
	}
}

func executeRcloneSyncCmd(filePath string, remote string) *exec.Cmd {
	return exec.Command("rclone", "sync", filePath, remote)
}
