package bidder

import "testing"

func TestMultiply(t *testing.T) {
    result := Multiply(2, 3)
    if result != 6 {
        t.Errorf("Multiply(2, 3) = %d; want 6", result)
    }
}
