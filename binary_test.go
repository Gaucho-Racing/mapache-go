package mapache

import "testing"

func TestBigEndianUnsignedIntToBinary(t *testing.T) {
	t.Run("Test Negative Number", func(t *testing.T) {
		_, err := BigEndianUnsignedIntToBinaryString(-1, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
