package core

import (
	"log"
	"os"
)

var Debug bool
var CyLog *log.Logger

func StartLogger() {
	f, err := os.OpenFile("current.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	if Debug {
		CyLog = log.New(f, "CyMM: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		CyLog = log.New(f, "CyMM: ", log.Ldate|log.Ltime)
	}
}
