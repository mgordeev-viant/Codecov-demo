package logger

import "testing"

func TestLog(t *testing.T) {
    result := Log("Hello")
    if result != "Log: Hello" {
        t.Errorf("Log(\"Hello\") = %s; want Log: Hello", result)
    }
}
