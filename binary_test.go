package mapache

import "testing"

func TestBigEndianUnsignedIntToBinaryString(t *testing.T) {
	t.Run("Test Negative Number", func(t *testing.T) {
		_, err := BigEndianUnsignedIntToBinaryString(-1, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test 0 Bytes", func(t *testing.T) {
		_, err := BigEndianUnsignedIntToBinaryString(100, 0)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test Number Too Large", func(t *testing.T) {
		_, err := BigEndianUnsignedIntToBinaryString(31241, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

func TestBigEndianUnsignedIntToBinary(t *testing.T) {
	t.Run("Test Negative Number", func(t *testing.T) {
		_, err := BigEndianUnsignedIntToBinary(-1, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test 0 Bytes", func(t *testing.T) {
		_, err := BigEndianUnsignedIntToBinary(100, 0)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test Number Too Large", func(t *testing.T) {
		_, err := BigEndianUnsignedIntToBinary(31241, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
