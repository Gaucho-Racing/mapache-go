package mapache

import "fmt"

type Message struct {
	ID     int
	Data   []byte
	Fields []Field
}

type Field struct {
	// Name of the field. Will be mapped to a signal name unless otherwise specified by ExportSignalFunc.
	Name string
	// Bytes, Size, Sign, and Endian are used to properly decode and encode the signal.
	Bytes  []byte
	Size   int
	Sign   SignMode
	Endian Endian
	// Value is the integer value of the field.
	Value int
	// ExportSignalFunc is the function that is used to export the field as an array of signals.
	ExportSignalFunc ExportSignalFunc
}

// ExportSignalFunc is a function that indicates how a field should be exported as an array of signals.
// Any required scaling will be applied here. If ExportSignalFunc is not set, the field will be directly
// exported as a single signal without scaling.
type ExportSignalFunc func(Field) []Signal

// NewField creates a new Field object with the given name, bytes, size, sign, endian, and export function.
// If no export function is provided, DefaultSignalExportFunc will be used.
func NewField(name string, bytes []byte, size int, sign SignMode, endian Endian, exportSignalFunc ExportSignalFunc) Field {
	return Field{
		Name:             name,
		Bytes:            bytes,
		Size:             size,
		Sign:             sign,
		Endian:           endian,
		ExportSignalFunc: exportSignalFunc,
	}
}

// Decode takes a Field object, decodes the bytes into an integer value, and returns the decoded Field object.
func (f Field) Decode() Field {
	if f.Sign == Signed && f.Endian == BigEndian {
		f.Value = BigEndianBytesToSignedInt(f.Bytes)
	} else if f.Sign == Signed && f.Endian == LittleEndian {
		f.Value = LittleEndianBytesToSignedInt(f.Bytes)
	} else if f.Sign == Unsigned && f.Endian == BigEndian {
		f.Value = BigEndianBytesToUnsignedInt(f.Bytes)
	} else if f.Sign == Unsigned && f.Endian == LittleEndian {
		f.Value = LittleEndianBytesToUnsignedInt(f.Bytes)
	}
	return f
}

// Encode takes a Field object, encodes the integer value into bytes, and returns the encoded Field object.
func (f Field) Encode() (Field, error) {
	var err error
	if f.Sign == Signed && f.Endian == BigEndian {
		f.Bytes, err = BigEndianSignedIntToBinary(f.Value, f.Size)
	} else if f.Sign == Signed && f.Endian == LittleEndian {
		f.Bytes, err = LittleEndianSignedIntToBinary(f.Value, f.Size)
	} else if f.Sign == Unsigned && f.Endian == BigEndian {
		f.Bytes, err = BigEndianUnsignedIntToBinary(f.Value, f.Size)
	} else if f.Sign == Unsigned && f.Endian == LittleEndian {
		f.Bytes, err = LittleEndianUnsignedIntToBinary(f.Value, f.Size)
	} else {
		return f, fmt.Errorf("invalid sign or endian")
	}
	return f, err
}

// CheckBit takes a Field object and a bit position, and returns true if the bit at the given position is set to 1, otherwise false.
func (f Field) CheckBit(bit int) bool {
	return (f.Bytes[0] & (1 << bit)) != 0
}

// ExportSignals takes a Field object and exports it as an array of signals.
func (f Field) ExportSignals() []Signal {
	if f.ExportSignalFunc == nil {
		return DefaultSignalExportFunc(f)
	}
	return f.ExportSignalFunc(f)
}

// DefaultSignalExportFunc is the default export function for a field. It exports the field as a single signal with no scaling.
func DefaultSignalExportFunc(f Field) []Signal {
	return []Signal{Signal{
		Name:     f.Name,
		Value:    float64(f.Value),
		RawValue: f.Value,
	}}
}
