package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"divo/internal/config"
	"divo/internal/proxy"
	"divo/internal/util"

	"github.com/charmbracelet/log"
)

func main() {
	util.ClearScreen()

	cfg := config.LoadFromEnv()

	if len(os.Args) == 3 && os.Args[1] == "http" {
		localPort := os.Args[2]
		if _, err := strconv.Atoi(localPort); err != nil {
			log.Error("Port must be a number")
			os.Exit(1)
		}

		cfg.TargetURL = "http://localhost:" + localPort
		cfg.ListenPort = ":8081"
	} else if cfg.TargetURL == config.DefaultTarget && cfg.ListenPort == config.DefaultPort {
		log.Warn("Usage: divo http <local-port> or set TARGET_URL & LISTEN_PORT env vars")
		os.Exit(1)
	}

	lanIP := util.GetLocalIP()

	util.PrettyPrintf(util.Green, "[+] Exposing http://%s%s → %s\n\n", lanIP, cfg.ListenPort, cfg.TargetURL)
	util.PrettyInfo("[ ... Ctrl + C to quit ... ]")
	
	util.PrintQRCode("http://" + lanIP + cfg.ListenPort)

	go func() {
		if err := proxy.StartReverseProxy(cfg.TargetURL, cfg.ListenPort); err != nil {
			log.Error("Proxy error:", err)
			os.Exit(1)
		}
	}()

	// wait for interrupt signal (Ctrl+C)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	util.ClearScreen()
}
