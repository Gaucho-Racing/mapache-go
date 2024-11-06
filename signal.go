package mapache

import (
	"time"
)

type Signal struct {
	// ID is a unique identifier for the signal.
	ID string `json:"id" gorm:"primaryKey"`
	// VehicleID is the unique identifier for the vehicle that the signal is associated with.
	VehicleID string `json:"vehicle_id"`
	// Name represents the type of signal.
	Name string `json:"name"`
	// Description is a brief description of the signal.
	Description string `json:"description"`
	// Value is the value of the signal post-scaling.
	Value float64 `json:"value"`
	// RawValue is the raw value of the signal before scaling.
	RawValue int `json:"raw_value"`
	// ProducedAt is the time at which the signal was produced.
	ProducedAt time.Time `json:"produced_at"`
	// CreatedAt is the time at which the signal was created.
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;precision:6"`

	// Bytes, Size, Sign, and Endian are used to properly decode and encode the signal.
	Bytes  []byte   `json:"-" gorm:"-"`
	Size   int      `json:"-" gorm:"-"`
	Sign   SignMode `json:"-" gorm:"-"`
	Endian Endian   `json:"-" gorm:"-"`
	// ScalingFunc is the function that is used to scale the signal.
	ScalingFunc ScalingFunc `json:"-" gorm:"-"`
}

func (Signal) TableName() string {
	return "signal"
}

// ScalingFunc is a function that indicated how a value should be scaled.
type ScalingFunc func(float64) float64

// Scale scales the value of the signal by the given function.
func (s Signal) Scale() Signal {
	if s.ScalingFunc == nil {
		return s
	}
	s.Value = s.ScalingFunc(float64(s.RawValue))
	return s
}

// Decode decodes the bytes stored in a Signal object and returns a signal object with the decoded raw value.
func (s Signal) Decode() Signal {
	if s.Sign == Signed && s.Endian == BigEndian {
		s.RawValue = BigEndianBytesToSignedInt(s.Bytes)
	} else if s.Sign == Signed && s.Endian == LittleEndian {
		s.RawValue = LittleEndianBytesToSignedInt(s.Bytes)
	} else if s.Sign == Unsigned && s.Endian == BigEndian {
		s.RawValue = BigEndianBytesToUnsignedInt(s.Bytes)
	} else if s.Sign == Unsigned && s.Endian == LittleEndian {
		s.RawValue = LittleEndianBytesToUnsignedInt(s.Bytes)
	}
	return s
}

// Encode encodes the integer value stored in a Field object and returns a signal object with the encoded bytes.
func (s Signal) Encode() ([]byte, error) {
	var err error
	if s.Sign == Signed && s.Endian == BigEndian {
		s.Bytes, err = BigEndianSignedIntToBinary(s.RawValue, s.Size)
	} else if s.Sign == Signed && s.Endian == LittleEndian {
		s.Bytes, err = LittleEndianSignedIntToBinary(s.RawValue, s.Size)
	} else if s.Sign == Unsigned && s.Endian == BigEndian {
		s.Bytes, err = BigEndianUnsignedIntToBinary(s.RawValue, s.Size)
	} else if s.Sign == Unsigned && s.Endian == LittleEndian {
		s.Bytes, err = LittleEndianUnsignedIntToBinary(s.RawValue, s.Size)
	}
	return s.Bytes, err
}

// NewSignal creates a new Signal object with the provided vehicle ID, name, size, sign, endian, and scaling function.
// If the scaling function is nil, the signal will not be scaled (default to just x1).
func NewSignal(vehicleID string, name string, size int, sign SignMode, endian Endian, scalingFunc ScalingFunc) Signal {
	return Signal{
		VehicleID:   vehicleID,
		Name:        name,
		Size:        size,
		Sign:        sign,
		Endian:      endian,
		ScalingFunc: scalingFunc,
	}
}
