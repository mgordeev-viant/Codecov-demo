package logger

import "fmt"

func Log(message string) string {
    return fmt.Sprintf("Log: %s", message)
}
