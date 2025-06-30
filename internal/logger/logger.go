package logger

import (
	"log"
	"os"
)

func InitLogger() {
	file, err := os.OpenFile("./currency.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("cannot open log file: %v\n", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

}

func Info(format string, v ...any) {
	log.Printf("[INFO] " + format, v...)
}

func Warn(format string, v ...any) {
	log.Printf("[WARN] "+format, v...)
}

func Error(format string, v ...any) {
	log.Printf("[ERROR] "+format, v...)
}
