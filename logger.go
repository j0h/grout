package grout

import log15 "github.com/inconshreveable/log15"

// Log is a log15 instance which is the base for all loggers in grout
var Log = log15.New("module", "grout")

func init() {
	Log.SetHandler(log15.DiscardHandler())
}

// EnableProductionLog enables logging starting from INFO level. Thus information about requests will be logged. If the logpath is empty, no fileHandler will be created.
func EnableProductionLog(logfilePath string) {
	level := log15.LvlInfo
	var fileHandler log15.Handler
	if logfilePath != "" {
		var err error
		fileHandler, err = log15.FileHandler(logfilePath, log15.JsonFormatEx(false, true))
		if err != nil {
			panic(err)
		}
	} else {
		fileHandler = log15.DiscardHandler()
	}
	Log.SetHandler(log15.MultiHandler(
		log15.LvlFilterHandler(level, fileHandler),
		log15.LvlFilterHandler(level, log15.StdoutHandler),
	))
}

// EnableDebugLog enables logging starting from Debug level.
func EnableDebugLog() {
	Log.SetHandler(log15.LvlFilterHandler(log15.LvlDebug, log15.StdoutHandler))
}
