package mapache

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
)

func TestBigEndianUnsignedIntToBinaryString(t *testing.T) {
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
	// [0 10 63 44 118 255]
	t.Run("Test 44009551615 6 Byte", func(t *testing.T) {
		testValue := 44009551615
		v, _ := BigEndianUnsignedIntToBinary(testValue, 6)
		expected := []byte{0, 10, 63, 44, 118, 255}
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
	t.Run("Test 32767 2 Byte", func(t *testing.T) {
		v, _ := BigEndianSignedIntToBinaryString(32767, 2)
		expected := "0111111111111111"
		if v != expected {
			t.Errorf("Expected %v, got %v", expected, v)
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
	// [0]
	t.Run("Test 0 1 Byte", func(t *testing.T) {
		v, _ := BigEndianSignedIntToBinary(0, 1)
		expected := []byte{0}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [123]
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
	// [0 172]
	t.Run("Test 172 2 Byte", func(t *testing.T) {
		testValue := 172
		v, _ := BigEndianSignedIntToBinary(testValue, 2)
		fmt.Printf("%v\n", v)
		var buf bytes.Buffer
		val := int16(testValue)
		_ = binary.Write(&buf, binary.BigEndian, val)
		expected := buf.Bytes()
		fmt.Printf("%v\n", expected)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [127 255]
	t.Run("Test 32767 2 Byte", func(t *testing.T) {
		testValue := 32767
		v, _ := BigEndianSignedIntToBinary(testValue, 2)
		fmt.Printf("%v\n", v)
		var buf bytes.Buffer
		val := int16(testValue)
		_ = binary.Write(&buf, binary.BigEndian, val)
		expected := buf.Bytes()
		fmt.Printf("%v\n", expected)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [128 1]
	t.Run("Test -32767 2 Byte", func(t *testing.T) {
		testValue := -32767
		v, _ := BigEndianSignedIntToBinary(testValue, 2)
		fmt.Printf("%v\n", v)
		var buf bytes.Buffer
		val := int16(testValue)
		_ = binary.Write(&buf, binary.BigEndian, val)
		expected := buf.Bytes()
		fmt.Printf("%v\n", expected)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [25 153 151 231]
	t.Run("Test 429496295 4 Byte", func(t *testing.T) {
		testValue := 429496295
		v, _ := BigEndianSignedIntToBinary(testValue, 4)
		fmt.Printf("%v\n", v)
		var buf bytes.Buffer
		val := int32(testValue)
		_ = binary.Write(&buf, binary.BigEndian, val)
		expected := buf.Bytes()
		fmt.Printf("%v\n", expected)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [230 102 104 25]
	t.Run("Test -429496295 4 Byte", func(t *testing.T) {
		testValue := -429496295
		v, _ := BigEndianSignedIntToBinary(testValue, 4)
		fmt.Printf("%v\n", v)
		var buf bytes.Buffer
		val := int32(testValue)
		_ = binary.Write(&buf, binary.BigEndian, val)
		expected := buf.Bytes()
		fmt.Printf("%v\n", expected)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 10 63 44 118 255]
	t.Run("Test 44009551615 6 Byte", func(t *testing.T) {
		testValue := 44009551615
		v, _ := BigEndianSignedIntToBinary(testValue, 6)
		fmt.Printf("%v\n", v)
		expected := []byte{0, 10, 63, 44, 118, 255}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 65 137 55 71 243 174 255]
	t.Run("Test 18446744009551615 8 Byte", func(t *testing.T) {
		testValue := 18446744009551615
		v, _ := BigEndianSignedIntToBinary(testValue, 8)
		fmt.Printf("%v\n", v)
		var buf bytes.Buffer
		val := int64(testValue)
		_ = binary.Write(&buf, binary.BigEndian, val)
		expected := buf.Bytes()
		fmt.Printf("%v\n", expected)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [255 190 118 200 184 12 81 1]
	t.Run("Test -18446744009551615 8 Byte", func(t *testing.T) {
		testValue := -18446744009551615
		v, _ := BigEndianSignedIntToBinary(testValue, 8)
		fmt.Printf("%v\n", v)
		var buf bytes.Buffer
		val := int64(testValue)
		_ = binary.Write(&buf, binary.BigEndian, val)
		expected := buf.Bytes()
		fmt.Printf("%v\n", expected)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
}

func TestBigEndianBytesToUnsignedInt(t *testing.T) {
	// [0]
	t.Run("Test 0 1 Byte", func(t *testing.T) {
		input := []byte{0}
		expected := 0
		v := BigEndianBytesToUnsignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [123]
	t.Run("Test 123 1 Byte", func(t *testing.T) {
		input := []byte{123}
		expected := 123
		v := BigEndianBytesToUnsignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [255]
	t.Run("Test 255 1 Byte", func(t *testing.T) {
		input := []byte{255}
		expected := 255
		v := BigEndianBytesToUnsignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 172]
	t.Run("Test 172 2 Byte", func(t *testing.T) {
		input := []byte{0, 172}
		expected := 172
		v := BigEndianBytesToUnsignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [148 246]
	t.Run("Test 38134 2 Byte", func(t *testing.T) {
		input := []byte{148, 246}
		expected := 38134
		v := BigEndianBytesToUnsignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [25 153 151 231]
	t.Run("Test 429496295 4 Byte", func(t *testing.T) {
		input := []byte{25, 153, 151, 231}
		expected := 429496295
		v := BigEndianBytesToUnsignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 10 63 44 118 255]
	t.Run("Test 44009551615 6 Byte", func(t *testing.T) {
		input := []byte{0, 10, 63, 44, 118, 255}
		expected := 44009551615
		v := BigEndianBytesToUnsignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 65 137 55 71 243 174 255]
	t.Run("Test 18446744009551615 8 Byte", func(t *testing.T) {
		input := []byte{0, 65, 137, 55, 71, 243, 174, 255}
		expected := 18446744009551615
		v := BigEndianBytesToUnsignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
}

func TestBigEndianBytesToSignedInt(t *testing.T) {
	// [0]
	t.Run("Test 0 1 Byte", func(t *testing.T) {
		input := []byte{0}
		expected := 0
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [123]
	t.Run("Test 123 1 Byte", func(t *testing.T) {
		input := []byte{123}
		expected := 123
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	t.Run("Test 255 1 Byte", func(t *testing.T) {
		input := []byte{255}
		expected := 255
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 172]
	t.Run("Test 172 2 Byte", func(t *testing.T) {
		input := []byte{0, 172}
		expected := 172
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [127 255]
	t.Run("Test 32767 2 Byte", func(t *testing.T) {
		input := []byte{127, 255}
		expected := 32767
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [128 1]
	t.Run("Test -32767 2 Byte", func(t *testing.T) {
		input := []byte{128, 1}
		expected := -32767
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [25 153 151 231]
	t.Run("Test 429496295 4 Byte", func(t *testing.T) {
		input := []byte{25, 153, 151, 231}
		expected := 429496295
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [230 102 104 25]
	t.Run("Test -429496295 4 Byte", func(t *testing.T) {
		input := []byte{230, 102, 104, 25}
		expected := -429496295
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 10 63 44 118 255]
	t.Run("Test 44009551615 6 Byte", func(t *testing.T) {
		input := []byte{0, 10, 63, 44, 118, 255}
		expected := 44009551615
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 65 137 55 71 243 174 255]
	t.Run("Test 18446744009551615 8 Byte", func(t *testing.T) {
		input := []byte{0, 65, 137, 55, 71, 243, 174, 255}
		expected := 18446744009551615
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [255 190 118 200 184 12 81 1]
	t.Run("Test -18446744009551615 8 Byte", func(t *testing.T) {
		input := []byte{255, 190, 118, 200, 184, 12, 81, 1}
		expected := -18446744009551615
		v := BigEndianBytesToSignedInt(input)
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
}

func TestLittleEndianUnsignedIntToBinaryString(t *testing.T) {
	t.Run("Test 38134 2 Byte", func(t *testing.T) {
		v, _ := LittleEndianUnsignedIntToBinaryString(38134, 2)
		expected := "1111011010010100"
		if v != expected {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
}

func TestLittleEndianUnsignedIntToBinary(t *testing.T) {
	t.Run("Test Negative Number", func(t *testing.T) {
		_, err := LittleEndianUnsignedIntToBinary(-1, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test 0 Bytes", func(t *testing.T) {
		_, err := LittleEndianUnsignedIntToBinary(100, 0)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test Number Too Large", func(t *testing.T) {
		_, err := LittleEndianUnsignedIntToBinary(31241, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	t.Run("Test Number Too Large 2", func(t *testing.T) {
		_, err := LittleEndianUnsignedIntToBinary(3172123, 2)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
	// [0]
	t.Run("Test 0 1 Byte", func(t *testing.T) {
		v, _ := LittleEndianUnsignedIntToBinary(0, 1)
		expected := []byte{0}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [123]
	t.Run("Test 123 1 Byte", func(t *testing.T) {
		v, _ := LittleEndianUnsignedIntToBinary(123, 1)
		expected := []byte{123}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [255]
	t.Run("Test 255 1 Byte", func(t *testing.T) {
		v, _ := LittleEndianUnsignedIntToBinary(255, 1)
		expected := []byte{255}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 172]
	t.Run("Test 172 2 Byte", func(t *testing.T) {
		v, _ := LittleEndianUnsignedIntToBinary(172, 2)
		expected := make([]byte, 2)
		binary.LittleEndian.PutUint16(expected, uint16(172))
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [148 246]
	t.Run("Test 38134 2 Byte", func(t *testing.T) {
		v, _ := LittleEndianUnsignedIntToBinary(38134, 2)
		expected := make([]byte, 2)
		binary.LittleEndian.PutUint16(expected, uint16(38134))
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [25 153 151 231]
	t.Run("Test 429496295 4 Byte", func(t *testing.T) {
		v, _ := LittleEndianUnsignedIntToBinary(429496295, 4)
		expected := make([]byte, 4)
		binary.LittleEndian.PutUint32(expected, uint32(429496295))
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 10 63 44 118 255]
	t.Run("Test 44009551615 6 Byte", func(t *testing.T) {
		testValue := 44009551615
		v, _ := LittleEndianUnsignedIntToBinary(testValue, 6)
		expected := []byte{255, 118, 44, 63, 10, 0}
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
	// [0 65 137 55 71 243 174 255]
	t.Run("Test 18446744009551615 8 Byte", func(t *testing.T) {
		v, _ := LittleEndianUnsignedIntToBinary(18446744009551615, 8)
		expected := make([]byte, 8)
		binary.LittleEndian.PutUint64(expected, uint64(18446744009551615))
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	})
}
