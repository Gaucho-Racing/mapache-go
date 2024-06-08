package mapache

import "fmt"

// SignMode is a type to represent whether an integer is signed or unsigned.
type SignMode int

const (
	Signed   SignMode = 1
	Unsigned SignMode = 0
)

// Endian is a type to represent whether an integer is big endian or little endian.
type Endian int

const (
	BigEndian    Endian = 1
	LittleEndian Endian = 0
)

// Field is a struct to represent a specific item in a vehicle node.
// It contains the name of the field, the bytes that represent the field,
// the integer value of the field, the size of the field in bytes, the sign
// mode of the field, and the endian mode of the field.
type Field struct {
	Name   string
	Bytes  []byte
	Value  int
	Size   int
	Sign   SignMode
	Endian Endian
}

// Decode decodes the bytes stored in a Field object and returns the integer value.
func (f Field) Decode() int {
	if f.Sign == Signed && f.Endian == BigEndian {
		return BigEndianBytesToSignedInt(f.Bytes)
	} else if f.Sign == Signed && f.Endian == LittleEndian {
		return LittleEndianBytesToSignedInt(f.Bytes)
	} else if f.Sign == Unsigned && f.Endian == BigEndian {
		return BigEndianBytesToUnsignedInt(f.Bytes)
	} else if f.Sign == Unsigned && f.Endian == LittleEndian {
		return LittleEndianBytesToUnsignedInt(f.Bytes)
	}
	return 0
}

// Encode encodes the integer value stored in a Field object and returns the bytes.
func (f Field) Encode() ([]byte, error) {
	if f.Sign == Signed && f.Endian == BigEndian {
		return BigEndianSignedIntToBinary(f.Value, f.Size)
	} else if f.Sign == Signed && f.Endian == LittleEndian {
		return LittleEndianSignedIntToBinary(f.Value, f.Size)
	} else if f.Sign == Unsigned && f.Endian == BigEndian {
		return BigEndianUnsignedIntToBinary(f.Value, f.Size)
	} else if f.Sign == Unsigned && f.Endian == LittleEndian {
		return LittleEndianUnsignedIntToBinary(f.Value, f.Size)
	}
	return nil, fmt.Errorf("invalid sign or endian")
}

// CheckBit checks if a specific bit is set in the first byte of a Field object.
func (f Field) CheckBit(bit int) bool {
	return (f.Bytes[0] & (1 << bit)) != 0
}

// String returns a string representation of a Field object.
func (f Field) String() string {
	return fmt.Sprintf("%s: %d", f.Name, f.Value)
}

// NewField creates a new Field object with the provided name, size, sign mode, and endian mode.
func NewField(name string, size int, sign SignMode, endian Endian) Field {
	return Field{
		Name:   name,
		Size:   size,
		Sign:   sign,
		Endian: endian,
	}
}

// Node is a type to represent a vehicle node. It is a slice of Fields,
// where each Field represents a specific item in the vehicle node.
type Node []Field

// Length returns the number of fields in a Node.
func (n Node) Length() int {
	return len(n)
}

// Size returns the total size of a Node in bytes.
func (n Node) Size() int {
	size := 0
	for _, field := range n {
		size += field.Size
	}
	return size
}

// String returns a string representation of a Node object.
func (n Node) String() string {
	str := ""
	for _, field := range n {
		str += field.String() + "\n"
	}
	return str
}

// FillFromBytes fills the fields of a Node with the bytes provided in the data slice.
// It decodes the bytes into integer values and stores them in the Value field of each Field.
// It returns an error if the length of the data slice does not match the total size of the Node.
func (n Node) FillFromBytes(data []byte) error {
	if len(data) != n.Size() {
		return fmt.Errorf("invalid data length, expected %d bytes", n.Size())
	}
	counter := 0
	for i, field := range n {
		field.Bytes = data[counter : counter+field.Size]
		counter += field.Size
		field.Value = field.Decode()
		n[i] = field
	}
	return nil
}

// FillFromInts fills the fields of a Node with the integer values provided in the ints slice.
// It encodes the integer values into bytes and stores them in the Bytes field of each Field.
// It returns an error if the length of the ints slice does not match the number of fields in the Node.
func (n Node) FillFromInts(ints []int) error {
	if len(ints) != n.Length() {
		return fmt.Errorf("invalid ints length, expected %d ints", n.Length())
	}
	for i, field := range n {
		field.Value = ints[i]
		bytes, err := field.Encode()
		if err != nil {
			return err
		}
		field.Bytes = bytes
		n[i] = field
	}
	return nil
}
