package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    if result != 2 {
        t.Errorf("Subtract(5, 3) = %d; want 2", result)
    }
}

func TestMultiply(t *testing.T) {
    result := Multiply(2, 3)
    if result != 6 {
        t.Errorf("Multiply(2, 3) = %d; want 6", result)
    }
}

func TestDivide(t *testing.T) {
    result, err := Divide(6, 3)
    if err != nil {
        t.Errorf("Divide(6, 3) returned error: %v", err)
    }
    if result != 2 {
        t.Errorf("Divide(6, 3) = %d; want 2", result)
    }

    _, err = Divide(1, 0)
    if err == nil {
        t.Error("Divide(1, 0) did not return an error")
    }
}
