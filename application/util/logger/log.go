package logger

import (
	"github.com/chentongming/yszcc/application/util/config"
	"log"
	"os"
)

const DEBUG = 100
const INFO = 200
const NOTICE = 250
const WARNING = 300
const ERROR = 400
const CRITICAL = 500
const ALERT = 550
const EMERGENCY = 600

var goLogger *log.Logger

func init() {
	goLogger = new(log.Logger)
}

func Debug(content string, logName string) {
	go writeLog("DEBUG\t", content, logName, DEBUG)
}

func Info(content string, logName string) {
	go writeLog("INFO\t", content, logName, INFO)
}

func Notice(content string, logName string) {
	go writeLog("NOTICE\t", content, logName, NOTICE)
}
func Warning(content string, logName string) {
	go writeLog("WARNING\t", content, logName, WARNING)
}

func Error(content string, logName string) {
	go writeLog("ERROR\t", content, logName, ERROR)
}

func Crital(content string, logName string) {
	go writeLog("CRITICAL\t", content, logName, CRITICAL)
}
func Alert(content string, logName string) {
	go writeLog("ALERT\t", content, logName, ALERT)
}
func Emergency(content string, logName string) {
	go writeLog("EMERGENCY\t", content, logName, EMERGENCY)
}

func writeLog(prefix string, content string, logName string, level int) {
	filePath := config.Get("logPath") + logName + ".log"
	logfile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	goLogger := log.New(logfile, prefix, log.LstdFlags|log.Lshortfile)
	goLogger.Output(3, "\t"+content+"\t")

}
