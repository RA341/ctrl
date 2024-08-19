package main

import (
	"ctrl/core/config"
	"ctrl/core/fs"
	"ctrl/core/system"
	"ctrl/core/updater"
	"ctrl/core/utils"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msgf("CTRL Version: %s", updater.Version)

	if runtime.GOOS == "linux" && !verifyRootStatus() {
		log.Fatal().Msg("The program needs to run with root privileges")
	} else if runtime.GOOS == "windows" && !verifyRootStatus() {
		// todo windows support
	}

	config.Load()
	system.RegisterService()
	//qbit.InitBasePath()

	//cli := docker.InitDocker()
	//defer docker.DisposeDocker(cli)

	// system checks
	//SystemStatus()
	//docker.ListDocker(cli)

	/////////////////////////////////////////////////////////////////////////////
	// GRPC setup
	go func() {
		grpcPort := "9221"
		portS := fmt.Sprintf(":%s", grpcPort)

		listen, err := net.Listen("tcp", portS)
		if err != nil {
			log.Error().Err(err).Msgf("failed to start server on %s", grpcPort)
			return
		}

		srv := grpc.NewServer()

		fsSrv := &fs.FileSrv{}
		fs.RegisterFilesystemServer(srv, fsSrv)

		sysSrv := &system.SysSrv{}
		system.RegisterSystemServer(srv, sysSrv)

		log.Info().Msgf("Grpc server started on %s", grpcPort)
		err = srv.Serve(listen)
		if err != nil {
			log.Error().Err(err).Msgf("failed to start server on %s", grpcPort)
			return
		}
	}()
	/////////////////////////////////////////////////////////////////////////////

	/////////////////////////////////////////////////////////////////////////////
	// http server setup
	// misc stuff
	http.HandleFunc("/status", status)
	http.HandleFunc("/test", test)

	// start periodic func
	//go runPeriodicTasks(cli)

	settings := config.Get()
	port := strconv.Itoa(settings.Network.Port)
	result := fmt.Sprintf("%s:%s", settings.Network.Host, port)

	log.Info().Msg("Starting http server on " + port)
	err := http.ListenAndServe(result, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	/////////////////////////////////////////////////////////////////////////////
}

func test(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}

func status(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func verifyRootStatus() bool {
	return os.Geteuid() == 0
}

func SystemStatus() {
	utils.WebhookStatus()
}
