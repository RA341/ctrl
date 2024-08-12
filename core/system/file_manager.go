package system

import (
	"github.com/rs/zerolog/log"
	"os"
)

const uid = 1000
const gid = 1000

func changeUserPermission(path string) {
	err := os.Chown(path, uid, gid)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to change user permission for path %s", path)
		return
	}
}

func changeDirectory(path string) {
	err := os.Chdir(path)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to change directory for path %s", path)
		return
	}
}

func linkFolder(src string, dest string) {
	entries := listDir(src)
	if entries == nil {
		log.Error().Msgf("Failed to list folder %s", src)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			linkFolder(src+"/"+entry.Name(), dest+"/"+entry.Name())
			continue
		}
		linkFile(src+"/"+entry.Name(), dest+"/"+entry.Name())
	}
}

func listDir(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)

	if err != nil {
		log.Error().Err(err).Msgf("Failed to list directory %s", dir)
		return nil
	}

	return entries
}

func linkFile(src string, dest string) {
	err := os.Link(src, dest)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to link file %s to %s", src, dest)
		return
	}
}
