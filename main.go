package main

import (
	"CyMM/core"
	"flag"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	/* Log setup */
	lFileName := time.Now().Format("2006-01-02_15-04-05") + ".log"
	runLogFile, _ := os.OpenFile(
		lFileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	defer runLogFile.Close()
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	DEBUG := flag.Bool("debug", true, `Sets debug mode.`)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *DEBUG {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	flag.Parse()

	/* Let's start working */
	core.SetDatDirs()
	log.Info().Msg("Application Started...")

	c := core.Config{}
	c.Prep()

}
