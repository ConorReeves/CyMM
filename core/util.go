package core

import (
	"os"
	"path/filepath"
	"strings"
)

var DATA_LOC string
var CACHE_LOC string

func SuffixFind(dir string, suffix string) []string {
	dirRead, err := os.ReadDir(dir)
	if err != nil {
		CyLog.Fatal(`Could not read directory for suffix search - `, suffix, `: `, err)
	}

	var foundFiles []string

	for _, file := range dirRead {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".conf") {
			foundFiles = append(foundFiles, file.Name())
		}
	}
	return foundFiles
}

func SetDatDirs() {
	var d string
	var dErr error

	d, dErr = os.UserConfigDir()
	if dErr != nil {
		CyLog.Fatal(`Could not find user AppData location: `, dErr)
	}
	CyLog.Println(`User AppData location found @ `, d)
	DATA_LOC = filepath.Join(d, `CyMM`)

	d, dErr = os.UserCacheDir()
	if dErr != nil {
		CyLog.Fatal(`Could not find user Temp location: `, dErr)
	}
	CyLog.Println(`User Temp location found @ `, d)
	CACHE_LOC = filepath.Join(d, `CyMM`)

	CyLog.Println(`Using `, DATA_LOC, ` for our data.`)
	CyLog.Println(`Using `, CACHE_LOC, ` for our cached data.`)

}
