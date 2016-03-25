package geonames

import (
	"log"
	"os"
)

var debugMode bool
var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[geonames]", log.LstdFlags)
}

// EnableDebugMode enable debug mode
func EnableDebugMode() {
	debugMode = true
}

// DisableDebugMode disable debug mode
func DisableDebugMode() {
	debugMode = false
}

func logWarn(f string, vars ...interface{}) {
	if debugMode {
		logger.Printf("[WRN] "+f, vars...)
	}
}
