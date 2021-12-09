package main

import (
	"admin/config"
	"admin/core/log"
	"admin/server"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var configFile *string = flag.String("config", "./etc/config.yaml", "kingshard config file")
var version *bool = flag.Bool("v", false, "the version ")

var (
	BuildDate    string
	BuildVersion string
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	if *version {
		fmt.Printf("Git commit:%s\n", BuildVersion)
		fmt.Printf("Build time:%s\n", BuildDate)
		return
	}

	cfg, err := config.ParseConfigFile(*configFile)
	if err != nil {
		fmt.Println("parse config file error: ", err.Error())
		return
	}

	err = log.Setup(cfg.Log)
	if err != nil {
		fmt.Println("setup log error: ", err.Error())
		return
	}

	srv := new(server.Server)
	err = srv.Setup(cfg)
	if err != nil {
		fmt.Println("setup server error: ", err.Error())
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGPIPE,
		// syscall.SIGUSR1,
	)

	go func() {
		for {
			sig := <-sc
			if sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == syscall.SIGQUIT {
				_ = srv.Close()
			} else if sig == syscall.SIGPIPE {
				// do nothing
			}
		}
	}()

	_ = srv.Run()
}
