package logger

import (
    "fmt"
    "log"
    "os"
)

var (
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

func InitLog(
    infoHandle *os.File,
    warningHandle *os.File,
    errorHandle *os.File) {
    
    Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogMessage(message string, level string) {
    switch level {
    case "INFO":
        Info.Println(message)
    case "WARNING":
        Warning.Println(message)
    case "ERROR":
        Error.Println(message)
    default:
        fmt.Println("Unknown log level")
    }
}
