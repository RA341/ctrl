package backup

// Mothballed for now

//
//import (
//	"fmt"
//	"github.com/rs/zerolog/log"
//	"os/exec"
//	"sync"
//	"time"
//)
//
//var wg sync.WaitGroup
//var fileList = []string{"movies", "shows", "nmov", "bmov", "music"}
//
//func periodicRcloneTask() {
//	for {
//		log.Info().Msg("Running rclone tasks")
//
//		syncMediaFiles()
//
//		time.Sleep(1 * time.Hour)
//	}
//}
//
//func syncMediaFiles() {
//	// Start the goroutines for each task
//	for _, value := range fileList {
//		go syncDir(&wg, value)
//		wg.Wait()
//	}
//}
//
//func syncDir(wg *sync.WaitGroup, folder string) {
//	// notify wg once command is completed
//	file := "/mnt/pool/" + folder
//	remote := "oneMedia:media/" + folder
//
//	log.Info("Syncing " + file + " " + remote)
//
//	runGoRoutineRclone(wg, file, remote)
//
//	fmt.Println("Completed " + folder + " sync command")
//}
//
//func runGoRoutineRclone(wg *sync.WaitGroup, file string, remote string) {
//	// notify wg once command is completed
//	defer wg.Done()
//
//	cmd := executeRcloneSyncCmd(file, remote)
//	err := cmd.Run()
//	if err != nil {
//		log.Fatal().Err(err).Msgf("Failed to execute command rclone sync")
//		return
//	}
//}
//
//func executeRcloneSyncCmd(filePath string, remote string) *exec.Cmd {
//	return exec.Command("rclone", "sync", filePath, remote)
//}
