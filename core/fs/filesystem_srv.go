package fs

import (
	"context"
	"github.com/rs/zerolog/log"
)

type FileSrv struct {
	UnimplementedFilesystemServer
}

func (s *FileSrv) ListFiles(_ context.Context, path *Path) (*Folder, error) {
	entries := ListDir(path.GetPath())
	dir := ParseDir(path.GetPath(), entries)
	return dir, nil
}

func (s *FileSrv) LinkFolder(_ context.Context, linkFolders *InputFolders) (*LinkResult, error) {
	log.Info().Msgf("Linking %s to %s", linkFolders.GetSrcPath(), linkFolders.GetDestPath())

	err := LinkFolder(linkFolders.GetSrcPath(), linkFolders.GetDestPath())
	if err != nil {
		return nil, err
	}
	result := &LinkResult{}
	return result, nil
}

func (s *FileSrv) CreateFolder(_ context.Context, path *NewPath) (*LinkResult, error) {
	log.Info().Msgf("Creating folder %s", path.GetPath())

	err := CreateFolder(path.GetPath(), path.GetAnchorPath())
	if err != nil {
		return nil, err
	}

	result := &LinkResult{}
	return result, nil
}
