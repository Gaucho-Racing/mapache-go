package mapache

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
)

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
	t.Run("Test 38134 2 Byte", func(t *testing.T) {
		v, _ := BigEndianUnsignedIntToBinaryString(38134, 2)
		expected := "1001010011110110"
		if v != expected {
			t.Errorf("Expected %v, got %v", expected, v)
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
	t.Run("Test Number Too Large 2", func(t *testing.T) {
		_, err := BigEndianUnsignedIntToBinary(3172123, 2)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	// [0]
	t.Run("Test 0 1 Byte", func(t *testing.T) {
		v, _ := BigEndianUnsignedIntToBinary(0, 1)
		expected := []byte{0}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [123]
	t.Run("Test 123 1 Byte", func(t *testing.T) {
		v, _ := BigEndianUnsignedIntToBinary(123, 1)
		expected := []byte{123}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [255]
	t.Run("Test 255 1 Byte", func(t *testing.T) {
		v, _ := BigEndianUnsignedIntToBinary(255, 1)
		expected := []byte{255}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 172]
	t.Run("Test 172 2 Byte", func(t *testing.T) {
		v, _ := BigEndianUnsignedIntToBinary(172, 2)
		expected := make([]byte, 2)
		binary.BigEndian.PutUint16(expected, uint16(172))
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [148 246]
	t.Run("Test 38134 2 Byte", func(t *testing.T) {
		v, _ := BigEndianUnsignedIntToBinary(38134, 2)
		expected := make([]byte, 2)
		binary.BigEndian.PutUint16(expected, uint16(38134))
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [25 153 151 231]
	t.Run("Test 429496295 4 Byte", func(t *testing.T) {
		v, _ := BigEndianUnsignedIntToBinary(429496295, 4)
		expected := make([]byte, 4)
		binary.BigEndian.PutUint32(expected, uint32(429496295))
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 65 137 55 71 243 174 255]
	t.Run("Test 18446744009551615 8 Byte", func(t *testing.T) {
		v, _ := BigEndianUnsignedIntToBinary(18446744009551615, 8)
		expected := make([]byte, 8)
		binary.BigEndian.PutUint64(expected, uint64(18446744009551615))
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
}

func TestBigEndianSignedIntToBinaryString(t *testing.T) {
	t.Run("Test 0 Bytes", func(t *testing.T) {
		_, err := BigEndianSignedIntToBinaryString(100, 0)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test Number Too Large", func(t *testing.T) {
		_, err := BigEndianSignedIntToBinaryString(31241, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

func TestBigEndianSignedIntToBinary(t *testing.T) {
	t.Run("Test 0 Bytes", func(t *testing.T) {
		_, err := BigEndianSignedIntToBinary(100, 0)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test Number Too Large", func(t *testing.T) {
		_, err := BigEndianSignedIntToBinary(31241, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test Number Too Large 2", func(t *testing.T) {
		_, err := BigEndianSignedIntToBinary(3172123, 2)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test 0 1 Byte", func(t *testing.T) {
		v, _ := BigEndianSignedIntToBinary(0, 1)
		expected := []byte{0}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	t.Run("Test 123 1 Byte", func(t *testing.T) {
		v, _ := BigEndianSignedIntToBinary(123, 1)
		expected := []byte{123}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	t.Run("Test 255 1 Byte", func(t *testing.T) {
		_, err := BigEndianSignedIntToBinary(255, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test -172 2 Byte", func(t *testing.T) {
		v, _ := BigEndianSignedIntToBinary(-172, 2)
		var buf bytes.Buffer
		val := int16(-172)
		_ = binary.Write(&buf, binary.BigEndian, val)
		expected := buf.Bytes()
		fmt.Printf("%v\n", expected)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	t.Run("Test 32767 2 Byte", func(t *testing.T) {
		v, _ := BigEndianSignedIntToBinary(32767, 2)
		var buf bytes.Buffer
		val := int16(32767)
		_ = binary.Write(&buf, binary.BigEndian, val)
		expected := buf.Bytes()
		fmt.Printf("%v\n", expected)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	t.Run("Test -2147483648 4 Byte", func(t *testing.T) {
		v, _ := BigEndianSignedIntToBinary(-2147483648, 4)
		var buf bytes.Buffer
		val := int32(-2147483648)
		_ = binary.Write(&buf, binary.BigEndian, val)
		expected := buf.Bytes()
		fmt.Printf("%v\n", expected)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
}
