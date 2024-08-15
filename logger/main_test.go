package logger


import (
    "os"
    "testing"
)

func TestLogMessage(t *testing.T) {
    file, err := os.Create("test.log")
    if err != nil {
        t.Fatal(err)
    }
    defer file.Close()
    InitLog(file, file, file)
    
    LogMessage("This is an info message:", "INFO")
    LogMessage("This is a warning message:", "WARNING")
    LogMessage("This is an error message:", "ERROR")
}
