package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	debugLog = log.New(os.Stdout, "\033[36m[zorm debug]\033[0m ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	infoLog  = log.New(os.Stdout, "\033[34m[zorm  info]\033[0m ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	warnLog  = log.New(os.Stdout, "\033[32m[zorm  warn]\033[0m ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	errorLog = log.New(os.Stdout, "\033[31m[zorm error]\033[0m ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// log methods
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
	Warn   = warnLog.Printf
	Warnf  = warnLog.Printf
	Debug  = debugLog.Printf
	Debugf = debugLog.Printf
)

// log levels
const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelDisabled
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if LevelDebug < level {
		debugLog.SetOutput(ioutil.Discard)
	}
	if LevelError < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if LevelWarn < level {
		warnLog.SetOutput(ioutil.Discard)
	}
	if LevelInfo < level {
		infoLog.SetOutput(ioutil.Discard)
	}
	if level > LevelDisabled {

	}
}
