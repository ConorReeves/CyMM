package main

import (
	"CyMM/core"
	"fmt"
	"log"
	"strings"
)

var DATA_LOC, CACHE_LOC string
var CyLog *log.Logger
var Debug bool = true

type Config struct {
	isCurrent   bool
	isDefault   bool
	version     string
	gamePath    string
	gameVersion string
	configName  string
}

func (c *Config) load() string {
	fmt.Println(`Found following configs:`)
	/* This is not reading directory as needed. "The system cannot find the file specified." */
	CyLog.Println("DATA_LOC is ", DATA_LOC)
	var fileList []string = core.SuffixFind(DATA_LOC, ".conf")

	CyLog.Println(fileList)

	for _, file := range fileList {
		file = strings.TrimSpace(file)
		if file == "" {
			continue
		}
		fmt.Println(`Found Config: `, file)
	}

	fmt.Println(`Select your config file, or type 'NEW' to load defaults.`)
	var selectedConf, uInput string
	fmt.Scan(uInput)

	switch uInput {
	case "":
		fmt.Println(`No input given, loading defaults.`)
		return "default"
	case "NEW":
		fmt.Println(`Loading default config.`)
		return "default"
	default:
		fmt.Println(`Loading selected config.`)
	}

	return selectedConf
}

func (c *Config) create(cName string) string {
	fmt.Println(`Please select a name for your config file.`)
	var userInput string
	fmt.Scan(userInput)

	return ""
}

func main() {
	core.StartLogger()
	core.SetDatDirs()
	CyLog = core.CyLog
	CyLog.Println("Application Started...")

	c := Config{}
	c.load()
}
