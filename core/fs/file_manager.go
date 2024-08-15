package fs

import (
	"errors"
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

func CreateFolder(path string) error {
	err := os.MkdirAll(path, 0750)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to create folder for path %s", path)
		return err
	}
	ChangeUserPermission(path)
	return nil
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

func LinkFolder(src string, dest string) error {
	log.Info().Msgf("Linking folder")

	err := os.MkdirAll(dest, 0750)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to make directory for path %s", dest)
		return nil
	}
	ChangeUserPermission(dest)

	isDir, err := isDirectory(src)
	if err != nil {
		return err
	}
	if !isDir {
		log.Info().Msg("Linking src file to dest dir")
		err := os.Link(src, dest+"/"+filepath.Base(src))
		if err != nil {
			return err
		}
		return nil
	}

	entries := ListDir(src)
	if entries == nil {
		log.Error().Msgf("Failed to list folder %s", src)
		return errors.New("failed to list folder")
	}

	for _, entry := range entries {
		if entry.IsDir() {
			err := LinkFolder(src+"/"+entry.Name(), dest+"/"+entry.Name())
			if err != nil {
				return err
			}
			continue
		}
		err := LinkFile(src+"/"+entry.Name(), dest+"/"+entry.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func ListDir(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to list directory %s", dir)
		return nil
	}

	return entries
}

func LinkFile(src string, dest string) error {
	err := os.Link(src, dest)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to link file %s to %s", src, dest)
		return err
	}
	return nil
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
