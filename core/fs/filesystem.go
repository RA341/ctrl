package fs

import (
	"context"
	"github.com/rs/zerolog/log"
)

type FileSrv struct {
	UnimplementedFilesystemServer
}

func (s *FileSrv) ListFiles(context context.Context, path *Path) (*Folder, error) {
	entries := ListDir(path.GetPath())
	dir := ParseDir(path.GetPath(), entries)
	return dir, nil
}

func (s *FileSrv) LinkFolder(context context.Context, linkFolders *InputFolders) (*LinkResult, error) {
	log.Info().Msgf("Linking %s to %s", linkFolders.GetSrcPath(), linkFolders.GetDestPath())

	result := &LinkResult{}
	return result, nil
}

func (s *FileSrv) CreateFolder(context context.Context, path *Path) (*LinkResult, error) {
	log.Info().Msgf("Creating folder %s", path.GetPath())

	result := &LinkResult{}
	return result, nil
}
