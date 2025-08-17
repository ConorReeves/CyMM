package core

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

var (
	DATA_DIR string
	TEMP_DIR string
	CFG_DIR  string
)

func SuffixFind(dir string, suffix string) []string {
	log.Info().Msgf(`Reading %s for files ending in %s`, dir, suffix)
	dirRead, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal().
			Err(err).
			Str(`suffix`, suffix)
	}

	var foundFiles []string

	for _, file := range dirRead {
		if !file.IsDir() && strings.HasSuffix(file.Name(), suffix) {
			foundFiles = append(foundFiles, file.Name())
		}
	}
	return foundFiles
}

func SetDatDirs() (string, string, string) {
	var d string
	var dErr error

	d, dErr = os.UserConfigDir()
	if dErr != nil {
		log.Fatal().
			Err(dErr).
			Msg(`Could not find user AppData location.`)
	}
	log.Debug().Msgf(`User AppData location found @ %s`, d)
	DATA_DIR = filepath.Join(d, `CyMM`)

	d, dErr = os.UserCacheDir()
	if dErr != nil {
		log.Fatal().
			Err(dErr).
			Msg(`Could not find user Temp location.`)
	}
	log.Debug().Msgf(`User Temp location found @ %s`, d)
	TEMP_DIR = filepath.Join(d, `CyMM`)

	CFG_DIR = filepath.Join(DATA_DIR, "configs")

	log.Info().Msgf(`Using %s for our data.`, DATA_DIR)
	log.Info().Msgf(`Using %s for our temp data.`, TEMP_DIR)
	log.Info().Msgf(`Using %s as our configs dir.`, CFG_DIR)

	return DATA_DIR, TEMP_DIR, CFG_DIR
}
