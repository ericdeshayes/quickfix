package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/config"
	"log"
	"os"
	"path"
	"strings"
)

type fileLog struct {
	eventLogger   *log.Logger
	messageLogger *log.Logger
}

func (l fileLog) OnIncoming(msg string) {
	l.messageLogger.Print(msg)
}

func (l fileLog) OnOutgoing(msg string) {
	l.messageLogger.Print(msg)
}

func (l fileLog) OnEvent(msg string) {
	l.eventLogger.Print(msg)
}

func (l fileLog) OnEventf(format string, v ...interface{}) {
	l.eventLogger.Printf(format, v...)
}

type fileLogFactory struct {
	globalLogPath   string
	sessionLogPaths map[SessionID]string
}

//NewFileLogFactory creates an instance of LogFactory that writes messages and events to file.
//The location of global and session log files is configured via FileLogPath.
func NewFileLogFactory(settings *Settings) (LogFactory, error) {
	logFactory := fileLogFactory{}

	var err error
	if logFactory.globalLogPath, err = settings.GlobalSettings().Setting(config.FileLogPath); err != nil {
		return logFactory, requiredConfigurationMissing(config.FileLogPath)
	}

	logFactory.sessionLogPaths = make(map[SessionID]string)

	for sid, sessionSettings := range settings.SessionSettings() {
		logPath, err := sessionSettings.Setting(config.FileLogPath)
		if err != nil {
			return logFactory, requiredConfigurationMissing(config.FileLogPath)
		}
		logFactory.sessionLogPaths[sid] = logPath
	}

	return logFactory, nil
}

func newFileLog(prefix string, logPath string) (fileLog, error) {
	l := fileLog{}

	eventLogName := path.Join(logPath, prefix+".event.current.log")
	messageLogName := path.Join(logPath, prefix+".messages.current.log")

	if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
		return l, err
	}

	fileFlags := os.O_RDWR | os.O_CREATE | os.O_APPEND
	eventFile, err := os.OpenFile(eventLogName, fileFlags, os.ModePerm)
	if err != nil {
		return l, err
	}

	messageFile, err := os.OpenFile(messageLogName, fileFlags, os.ModePerm)
	if err != nil {
		return l, err
	}

	l.eventLogger = log.New(eventFile, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	l.messageLogger = log.New(messageFile, "", 0)

	return l, nil
}

func (f fileLogFactory) Create() (Log, error) {
	return newFileLog("GLOBAL", f.globalLogPath)
}

func (f fileLogFactory) CreateSessionLog(sessionID SessionID) (Log, error) {
	logPath, ok := f.sessionLogPaths[sessionID]

	if !ok {
		return nil, fmt.Errorf("logger not defined for %v", sessionID)
	}

	prefixParts := []string{sessionID.BeginString, sessionID.SenderCompID, sessionID.TargetCompID}
	if len(sessionID.Qualifier) > 0 {
		prefixParts = append(prefixParts, sessionID.Qualifier)
	}

	return newFileLog(strings.Join(prefixParts, "-"), logPath)
}
