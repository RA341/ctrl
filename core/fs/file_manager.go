package fs

import (
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

const uid = 1000
const gid = 1000

func ChangeUserPermission(path string) {
	err := os.Chown(path, uid, gid)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to change user permission for path %s", path)
		return
	}
}

func ChangeDirectory(path string) string {
	err := os.Chdir(path)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to change directory for path %s", path)
		return ""
	}
	return path
}

func LinkFolder(src string, dest string) {
	entries := ListDir(src)
	if entries == nil {
		log.Error().Msgf("Failed to list folder %s", src)
		return
	}

	err := os.MkdirAll(dest, 0750)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to make directory for path %s", dest)
		return
	}
	ChangeUserPermission(dest)

	for _, entry := range entries {
		if entry.IsDir() {
			LinkFolder(src+"/"+entry.Name(), dest+"/"+entry.Name())
			continue
		}
		LinkFile(src+"/"+entry.Name(), dest+"/"+entry.Name())
	}
}

func ListDir(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to list directory %s", dir)
		return nil
	}

	return entries
}

func LinkFile(src string, dest string) {
	err := os.Link(src, dest)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to link file %s to %s", src, dest)
		return
	}
}

func ParseDir(filePath string, entries []os.DirEntry) *Folder {
	fullPath, err := filepath.Abs(filePath)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get full path %s", filePath)
		return nil
	}

	directory := Folder{}
	directory.FullPath = &fullPath

	for _, entry := range entries {
		if entry.IsDir() {
			fold := Folder{}
			tmp := fullPath + "/" + entry.Name()
			fold.FullPath = &tmp
			directory.Folders = append(directory.Folders, &fold)
		} else {
			tmp := entry.Name()
			file := File{Name: &tmp}

			directory.Files = append(directory.Files, &file)
		}
	}

	return &directory
}
