package mapache

// import (
// 	"reflect"
// 	"testing"
// )

// func TestField_CheckBit(t *testing.T) {
// 	t.Run("Test Bit Set", func(t *testing.T) {
// 		f := Field{
// 			Bytes: []byte{0x01},
// 		}
// 		if !f.CheckBit(0) {
// 			t.Error("Expected true, got", f.CheckBit(0))
// 		}
// 	})
// 	t.Run("Test Bit Not Set", func(t *testing.T) {
// 		f := Field{
// 			Bytes: []byte{0x00},
// 		}
// 		if f.CheckBit(0) {
// 			t.Error("Expected false, got", f.CheckBit(0))
// 		}
// 	})
// }

// func TestField_String(t *testing.T) {
// 	f := Field{
// 		Name:  "test",
// 		Value: 1,
// 	}
// 	if f.String() != "test: 1" {
// 		t.Error("Expected test: 1, got", f.String())
// 	}
// }

// func TestNewField(t *testing.T) {
// 	f := NewField("test", 1, Signed, BigEndian)
// 	if f.Name != "test" || f.Size != 1 || f.Sign != Signed || f.Endian != BigEndian {
// 		t.Error("Expected test, 1, Signed, BigEndian, got", f.Name, f.Size, f.Sign, f.Endian)
// 	}
// }

// func TestNode_FillFromBytes(t *testing.T) {
// 	t.Run("Test FillFromBytes", func(t *testing.T) {
// 		n := Node{
// 			{Name: "test", Size: 2, Sign: Unsigned, Endian: BigEndian},
// 		}
// 		err := n.FillFromBytes([]byte{0x00, 0x01})
// 		if err != nil {
// 			t.Error("Expected nil, got", err)
// 		}
// 		if n[0].Value != 1 {
// 			t.Error("Expected 1, got", n[0].Value)
// 		}
// 	})
// 	t.Run("Test FillFromBytes Error", func(t *testing.T) {
// 		n := Node{
// 			{Name: "test", Size: 2, Sign: Unsigned, Endian: BigEndian},
// 		}
// 		err := n.FillFromBytes([]byte{0x00})
// 		if err == nil {
// 			t.Error("Expected err, got nil")
// 		}
// 	})
// }

// func TestNode_FillFromInts(t *testing.T) {
// 	t.Run("Test FillFromInts", func(t *testing.T) {
// 		n := Node{
// 			{Name: "test", Size: 2, Sign: Unsigned, Endian: BigEndian},
// 		}
// 		err := n.FillFromInts([]int{1})
// 		if err != nil {
// 			t.Error("Expected nil, got", err)
// 		}
// 		if n[0].Value != 1 {
// 			t.Error("Expected 1, got", n[0].Value)
// 		}
// 	})
// 	t.Run("Test FillFromInts Error", func(t *testing.T) {
// 		n := Node{
// 			{Name: "test", Size: 2, Sign: Unsigned, Endian: BigEndian},
// 		}
// 		err := n.FillFromInts([]int{})
// 		if err == nil {
// 			t.Error("Expected err, got nil")
// 		}
// 	})
// 	t.Run("Test FillFromInts Error 2", func(t *testing.T) {
// 		n := Node{
// 			{Name: "test", Size: 2, Sign: -1, Endian: -1},
// 		}
// 		err := n.FillFromInts([]int{1})
// 		if err == nil {
// 			t.Error("Expected err, got nil")
// 		}
// 	})
// }

// func TestNode_Length(t *testing.T) {
// 	n := Node{
// 		{Name: "test", Size: 1},
// 	}
// 	if n.Length() != 1 {
// 		t.Error("Expected 1, got", n.Length())
// 	}
// }

// func TestNode_Size(t *testing.T) {
// 	n := Node{
// 		{Name: "test", Size: 4},
// 	}
// 	if n.Size() != 4 {
// 		t.Error("Expected 4, got", n.Size())
// 	}
// }

// func TestNode_String(t *testing.T) {
// 	n := Node{
// 		{Name: "test", Value: 1},
// 	}
// 	if n.String() != "test: 1\n" {
// 		t.Error("Expected test: 1, got", n.String())
// 	}
// }
