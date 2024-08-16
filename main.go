package main

import (
	"ctrl/core/config"
	"ctrl/core/fs"
	qbit "ctrl/core/qbit"
	"ctrl/core/system"
	"ctrl/core/updater"
	"ctrl/core/utils"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
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
	//qbit.InitBasePath()
	//system.RegisterService()

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
	// system power controls
	http.HandleFunc("/shutdown", system.ExecShutDown)
	http.HandleFunc("/reboot", system.ExecReboot)
	http.HandleFunc("/sleep", system.ExecSleep)
	// misc stuff
	http.HandleFunc("/status", status)
	http.HandleFunc("/test", test)
	//http.HandleFunc("/device", deviceCheck)

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

func runPeriodicTasks(cli *client.Client) {
	ticker := time.NewTicker(time.Hour * 1)
	defer ticker.Stop()

	for range ticker.C {
		log.Info().Msg("Running stalled torrent search")
		go qbit.RunQbitChecks([]qbit.Check{qbit.ClientCheck, qbit.StalledCheck}, cli)
	}
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
	qbit.Status()
	utils.WebhookStatus()
}
