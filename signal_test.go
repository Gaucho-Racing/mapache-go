package mapache

// import (
// 	"reflect"
// 	"testing"
// )

// func TestSignal_Decode(t *testing.T) {
// 	t.Run("Test Big Endian Unsigned", func(t *testing.T) {
// 		f := Signal{
// 			Bytes:  []byte{0x00, 0x01},
// 			Sign:   Unsigned,
// 			Endian: BigEndian,
// 		}
// 		if f.Decode().RawValue != 1 {
// 			t.Error("Expected 1, got", f.Decode().RawValue)
// 		}
// 	})
// 	t.Run("Test Big Endian Signed", func(t *testing.T) {
// 		f := Signal{
// 			Bytes:  []byte{0x00, 0x01},
// 			Sign:   Signed,
// 			Endian: BigEndian,
// 		}
// 		if f.Decode().RawValue != 1 {
// 			t.Error("Expected 1, got", f.Decode().RawValue)
// 		}
// 	})
// 	t.Run("Test Little Endian Unsigned", func(t *testing.T) {
// 		f := Signal{
// 			Bytes:  []byte{0x00, 0x01},
// 			Sign:   Unsigned,
// 			Endian: LittleEndian,
// 		}
// 		if f.Decode().RawValue != 256 {
// 			t.Error("Expected 256, got", f.Decode().RawValue)
// 		}
// 	})
// 	t.Run("Test Little Endian Signed", func(t *testing.T) {
// 		f := Signal{
// 			Bytes:  []byte{0x00, 0x01},
// 			Sign:   Signed,
// 			Endian: LittleEndian,
// 		}
// 		if f.Decode().RawValue != 256 {
// 			t.Error("Expected 256, got", f.Decode().RawValue)
// 		}
// 	})
// 	t.Run("Test Other", func(t *testing.T) {
// 		f := Signal{
// 			Bytes:  []byte{0x00, 0x01},
// 			Sign:   -1,
// 			Endian: -1,
// 		}
// 		if f.Decode().RawValue != 0 {
// 			t.Error("Expected 0, got", f.Decode().RawValue)
// 		}
// 	})
// }

// func TestSignal_Encode(t *testing.T) {
// 	t.Run("Test Big Endian Unsigned", func(t *testing.T) {
// 		f := Signal{
// 			Sign:     Unsigned,
// 			Endian:   BigEndian,
// 			RawValue: 1,
// 			Size:     2,
// 		}
// 		b, err := f.Encode()
// 		if err != nil {
// 			t.Error("Expected nil, got", err)
// 		}
// 		expected := []byte{0x00, 0x01}
// 		if !reflect.DeepEqual(b.Bytes, expected) {
// 			t.Error("Expected", expected, "got", b.Bytes)
// 		}
// 	})
// 	t.Run("Test Big Endian Signed", func(t *testing.T) {
// 		f := Signal{
// 			Sign:     Signed,
// 			Endian:   BigEndian,
// 			RawValue: 1,
// 			Size:     2,
// 		}
// 		b, err := f.Encode()
// 		if err != nil {
// 			t.Error("Expected nil, got", err)
// 		}
// 		expected := []byte{0x00, 0x01}
// 		if !reflect.DeepEqual(b.Bytes, expected) {
// 			t.Error("Expected", expected, "got", b.Bytes)
// 		}
// 	})
// 	t.Run("Test Little Endian Unsigned", func(t *testing.T) {
// 		f := Signal{
// 			Sign:     Unsigned,
// 			Endian:   LittleEndian,
// 			RawValue: 1,
// 			Size:     2,
// 		}
// 		b, err := f.Encode()
// 		if err != nil {
// 			t.Error("Expected nil, got", err)
// 		}
// 		expected := []byte{0x01, 0x00}
// 		if !reflect.DeepEqual(b.Bytes, expected) {
// 			t.Error("Expected", expected, "got", b.Bytes)
// 		}
// 	})
// 	t.Run("Test Little Endian Signed", func(t *testing.T) {
// 		f := Signal{
// 			Sign:     Signed,
// 			Endian:   LittleEndian,
// 			RawValue: 1,
// 			Size:     2,
// 		}
// 		b, err := f.Encode()
// 		if err != nil {
// 			t.Error("Expected nil, got", err)
// 		}
// 		expected := []byte{0x01, 0x00}
// 		if !reflect.DeepEqual(b.Bytes, expected) {
// 			t.Error("Expected", expected, "got", b.Bytes)
// 		}
// 	})
// 	t.Run("Test Other", func(t *testing.T) {
// 		f := Signal{
// 			Sign:     -1,
// 			Endian:   -1,
// 			RawValue: 1,
// 			Size:     2,
// 		}
// 		_, err := f.Encode()
// 		if err == nil {
// 			t.Error("Expected err, got nil")
// 		}
// 	})
// }
