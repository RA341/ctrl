package system

import (
	"context"
	"github.com/rs/zerolog/log"
)

type SysSrv struct {
	UnimplementedSystemServer
}

func (s *SysSrv) Shutdown(context context.Context, empty *Empty) (*Empty, error) {
	log.Info().Msg("Setting Shutdown")
	ExecShutDown()

	result := &Empty{}
	return result, nil
}

func (s *SysSrv) Restart(context context.Context, empty *Empty) (*Empty, error) {
	log.Info().Msg("Setting Restart")
	ExecReboot()

	result := &Empty{}
	return result, nil
}

func (s *SysSrv) Sleep(context context.Context, empty *Empty) (*Empty, error) {
	log.Info().Msg("Setting Sleep")
	ExecSleep()

	result := &Empty{}
	return result, nil
}
