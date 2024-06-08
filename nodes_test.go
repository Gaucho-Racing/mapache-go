package mapache

import (
	"reflect"
	"testing"
)

func TestField_Decode(t *testing.T) {
	t.Run("Test Big Endian Unsigned", func(t *testing.T) {
		f := Field{
			Bytes:  []byte{0x00, 0x01},
			Sign:   Unsigned,
			Endian: BigEndian,
		}
		if f.Decode() != 1 {
			t.Error("Expected 1, got", f.Decode())
		}
	})
	t.Run("Test Big Endian Signed", func(t *testing.T) {
		f := Field{
			Bytes:  []byte{0x00, 0x01},
			Sign:   Signed,
			Endian: BigEndian,
		}
		if f.Decode() != 1 {
			t.Error("Expected 1, got", f.Decode())
		}
	})
	t.Run("Test Little Endian Unsigned", func(t *testing.T) {
		f := Field{
			Bytes:  []byte{0x00, 0x01},
			Sign:   Unsigned,
			Endian: LittleEndian,
		}
		if f.Decode() != 256 {
			t.Error("Expected 256, got", f.Decode())
		}
	})
	t.Run("Test Little Endian Signed", func(t *testing.T) {
		f := Field{
			Bytes:  []byte{0x00, 0x01},
			Sign:   Signed,
			Endian: LittleEndian,
		}
		if f.Decode() != 256 {
			t.Error("Expected 256, got", f.Decode())
		}
	})
	t.Run("Test Other", func(t *testing.T) {
		f := Field{
			Bytes:  []byte{0x00, 0x01},
			Sign:   -1,
			Endian: -1,
		}
		if f.Decode() != 0 {
			t.Error("Expected 0, got", f.Decode())
		}
	})
}

func TestField_Encode(t *testing.T) {
	t.Run("Test Big Endian Unsigned", func(t *testing.T) {
		f := Field{
			Sign:   Unsigned,
			Endian: BigEndian,
			Value:  1,
			Size:   2,
		}
		b, err := f.Encode()
		if err != nil {
			t.Error("Expected nil, got", err)
		}
		expected := []byte{0x00, 0x01}
		if !reflect.DeepEqual(b, expected) {
			t.Error("Expected", expected, "got", b)
		}
	})
	t.Run("Test Big Endian Signed", func(t *testing.T) {
		f := Field{
			Sign:   Signed,
			Endian: BigEndian,
			Value:  1,
			Size:   2,
		}
		b, err := f.Encode()
		if err != nil {
			t.Error("Expected nil, got", err)
		}
		expected := []byte{0x00, 0x01}
		if !reflect.DeepEqual(b, expected) {
			t.Error("Expected", expected, "got", b)
		}
	})
	t.Run("Test Little Endian Unsigned", func(t *testing.T) {
		f := Field{
			Sign:   Unsigned,
			Endian: LittleEndian,
			Value:  1,
			Size:   2,
		}
		b, err := f.Encode()
		if err != nil {
			t.Error("Expected nil, got", err)
		}
		expected := []byte{0x01, 0x00}
		if !reflect.DeepEqual(b, expected) {
			t.Error("Expected", expected, "got", b)
		}
	})
	t.Run("Test Little Endian Signed", func(t *testing.T) {
		f := Field{
			Sign:   Signed,
			Endian: LittleEndian,
			Value:  1,
			Size:   2,
		}
		b, err := f.Encode()
		if err != nil {
			t.Error("Expected nil, got", err)
		}
		expected := []byte{0x01, 0x00}
		if !reflect.DeepEqual(b, expected) {
			t.Error("Expected", expected, "got", b)
		}
	})
	t.Run("Test Other", func(t *testing.T) {
		f := Field{
			Sign:   Signed,
			Endian: LittleEndian,
			Value:  -1,
			Size:   -2,
		}
		_, err := f.Encode()
		if err == nil {
			t.Error("Expected err, got nil")
		}
	})
}

func TestField_CheckBit(t *testing.T) {
	t.Run("Test Bit Set", func(t *testing.T) {
		f := Field{
			Bytes: []byte{0x01},
		}
		if !f.CheckBit(0) {
			t.Error("Expected true, got", f.CheckBit(0))
		}
	})
	t.Run("Test Bit Not Set", func(t *testing.T) {
		f := Field{
			Bytes: []byte{0x00},
		}
		if f.CheckBit(0) {
			t.Error("Expected false, got", f.CheckBit(0))
		}
	})
}

func TestField_String(t *testing.T) {
	f := Field{
		Name:  "test",
		Value: 1,
	}
	if f.String() != "test: 1" {
		t.Error("Expected test: 1, got", f.String())
	}
}
