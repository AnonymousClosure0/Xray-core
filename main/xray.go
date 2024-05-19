package main

import "C"

import (
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"

	"github.com/xtls/xray-core/core"
	_ "github.com/xtls/xray-core/main/distro/all"
)

//export XrayRun
func XrayRun(configdir string) {
	file, err := os.Open(configdir)
	if err != nil {
		println(err.Error())
		newError("failed to load config files").AtError().WriteToLog()
		return
	}
	config, err := core.LoadConfig("json", file)
	if err != nil {
		newError("failed to load config files").Base(err).AtError().WriteToLog()
		return
	}

	server, err := core.New(config)
	if err != nil {
		newError("failed to create server").Base(err).AtError().WriteToLog()
		return
	}

	if err := server.Start(); err != nil {
		newError("Failed to start:").Base(err).AtError().WriteToLog()
		return
	}
	defer server.Close()

	/*
		conf.FileCache = nil
		conf.IPCache = nil
		conf.SiteCache = nil
	*/

	// Explicitly triggering GC to remove garbage from config loading.
	runtime.GC()
	debug.FreeOSMemory()

	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		<-osSignals
	}
	newError("exit").AtError().WriteToLog()
}

func main() {

}
