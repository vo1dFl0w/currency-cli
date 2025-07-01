package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[36m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
)

type Logger interface{
	Info(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
	Close()
}

type defaultLogger struct {
	infoLog     *log.Logger
	infoConsole *log.Logger
	warnLog     *log.Logger
	warnConsole *log.Logger
	errLog      *log.Logger
	errConsole  *log.Logger
	file        *os.File
	verbose     bool
}

var loggerInstance *defaultLogger

func InitLogger(verbose bool) Logger {
	file, err := os.OpenFile("./currency.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("cannot open log file: %v\n", err)
	}

	var output io.Writer = file
	if verbose {
		output = io.MultiWriter(os.Stdout, file)
	}

	loggerInstance = &defaultLogger{
		infoLog: log.New(output, "", log.Ldate|log.Ltime),
		infoConsole: log.New(os.Stdout, "", 0),
		warnLog: log.New(output, "", log.Ldate|log.Ltime),
		warnConsole: log.New(os.Stdout, "", 0),
		errLog: log.New(output, "", log.Ldate|log.Ltime),
		errConsole: log.New(os.Stdout, "", 0),
		file: file,
	}
	return loggerInstance
}

func (l *defaultLogger) Close() {
	if l.file != nil {
		_ = l.file.Close()
	}
}

func (l *defaultLogger) Info(format string, v ...any) {
	msg := fmt.Sprintf("[INFO] "+format, v...)
	l.infoLog.Output(2, addCaller(msg))
	if l.verbose {
		l.infoConsole.Output(2, colorBlue+msg+colorReset)
	}
}

func (l *defaultLogger) Warn(format string, v ...any) {
	msg := fmt.Sprintf("[WARN] "+format, v...)
	l.warnLog.Output(2, addCaller(msg))
	if l.verbose {
		l.warnConsole.Output(2, addCaller(colorYellow+msg+colorReset))
	}
}

func (l *defaultLogger) Error(format string, v ...any) {
	msg := fmt.Sprintf("[ERROR] "+format, v...)
	l.errLog.Output(2, addCaller(msg))
	l.errConsole.Output(2, addCaller(colorRed+msg+colorReset))
}

func addCaller(msg string) string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return msg
	}
	return fmt.Sprintf("%s:%d: %s", filepath.Base(file), line, msg)
}
