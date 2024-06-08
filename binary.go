package mapache

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// BigEndianUnsignedIntToBinaryString converts an unsigned integer to a binary
// string in big endian format. The input num will be packed into a number of
// bytes specified by length. If num is too large to fit in length bytes, an
// error will be returned.
//
// The function returns a string of 0s and 1s representing the binary.
func BigEndianUnsignedIntToBinaryString(num int, numBytes int) (string, error) {
	bytes, err := BigEndianUnsignedIntToBinary(num, numBytes)
	if err != nil {
		return "", err
	}
	var bs = ""
	for i := 0; i < numBytes; i++ {
		bs += fmt.Sprintf("%08b", bytes[i])
	}
	return bs, nil
}

// BigEndianUnsignedIntToBinary converts an unsigned integer to bytes in big endian
// format. The input num will be packed into a number of bytes specified by length.
// If num is too large to fit in length bytes, an error will be returned.
//
// The function returns a slice of bytes representing the binary.
func BigEndianUnsignedIntToBinary(num int, numBytes int) ([]byte, error) {
	if num < 0 {
		return nil, fmt.Errorf("cannot convert negative number to binary")
	} else if numBytes < 1 {
		return nil, fmt.Errorf("cannot convert to binary with less than 1 byte")
	} else if num >= 1<<(numBytes*8) && numBytes != 8 {
		return nil, fmt.Errorf("number is too large to fit in %d bytes", numBytes)
	}

	var result []byte
	if numBytes == 1 {
		return []byte{byte(num)}, nil
	} else if numBytes == 2 {
		result = make([]byte, 2)
		binary.BigEndian.PutUint16(result, uint16(num))
		return result, nil
	} else if numBytes == 4 {
		result = make([]byte, 4)
		binary.BigEndian.PutUint32(result, uint32(num))
		return result, nil
	} else if numBytes == 8 {
		result = make([]byte, 8)
		binary.BigEndian.PutUint64(result, uint64(num))
		return result, nil
	}

	// fallback for arbitrary number of bytes
	for i := 0; i < numBytes; i++ {
		result = append(result, byte(num>>uint((numBytes-i-1)*8)))
	}
	return result, nil
}

// BigEndianSignedIntToBinaryString converts a signed integer to a binary string
// in big endian format. The input num will be packed into a number of bytes specified
// by length. If num is too large to fit in length bytes, an error will be returned.
//
// The function returns a string of 0s and 1s representing the binary.
func BigEndianSignedIntToBinaryString(num int, numBytes int) (string, error) {
	bytes, err := BigEndianSignedIntToBinary(num, numBytes)
	if err != nil {
		return "", err
	}
	var bs = ""
	for i := 0; i < numBytes; i++ {
		bs += fmt.Sprintf("%08b", bytes[i])
		fmt.Printf("%02x ", bytes[i])
	}
	return bs, nil
}

// BigEndianSignedIntToBinary converts a signed integer to bytes in big endian format.
// The input num will be packed into a number of bytes specified by length. If num is
// too large to fit in length bytes, an error will be returned.
//
// The function returns a slice of bytes representing the binary.
func BigEndianSignedIntToBinary(num int, numBytes int) ([]byte, error) {
	if numBytes < 1 {
		return nil, fmt.Errorf("cannot convert to binary with less than 1 byte")
	}
	minValue := -1 << ((numBytes * 8) - 1)
	maxValue := (1 << ((numBytes * 8) - 1)) - 1
	if num < minValue || num > maxValue {
		return nil, fmt.Errorf("number is too large to fit in %d bytes", numBytes)
	}
	var result []byte
	var buf bytes.Buffer
	if numBytes == 1 {
		val := int8(num)
		err := binary.Write(&buf, binary.BigEndian, val)
		return buf.Bytes(), err
	} else if numBytes == 2 {
		val := int16(num)
		err := binary.Write(&buf, binary.BigEndian, val)
		return buf.Bytes(), err
	} else if numBytes == 4 {
		val := int32(num)
		err := binary.Write(&buf, binary.BigEndian, val)
		return buf.Bytes(), err
	} else if numBytes == 8 {
		val := int64(num)
		err := binary.Write(&buf, binary.BigEndian, val)
		return buf.Bytes(), err
	}

	// fallback for arbitrary number of bytes
	for i := 0; i < numBytes; i++ {
		shift := uint((numBytes - i - 1) * 8)
		result = append(result, byte(num>>shift))
	}
	return result, nil
}

// BigEndianBytesToUnsignedInt converts bytes to an unsigned integer in big endian format.
//
// The function returns the integer value of the bytes.
func BigEndianBytesToUnsignedInt(bytes []byte) int {
	var result int
	for i := 0; i < len(bytes); i++ {
		result += int(bytes[i]) << uint((len(bytes)-i-1)*8)
	}
	return result
}

// BigEndianBytesToSignedInt converts bytes to a signed integer in big endian format.
//
// The function returns the integer value of the bytes.
func BigEndianBytesToSignedInt(bytes []byte) int {
	var result int
	if bytes[0] >= 128 {
		result = -1 << uint((len(bytes)-1)*8)
	}
	for i := 0; i < len(bytes); i++ {
		result += int(bytes[i]) << uint((len(bytes)-i-1)*8)
	}
	return result
}

// LittleEndianUnsignedIntToBinaryString converts an unsigned integer to a binary string
// in little endian format. The input num will be packed into a number of bytes specified
// by length. If num is too large to fit in length bytes, an error will be returned.
//
// The function returns a string of 0s and 1s representing the binary.
func LittleEndianUnsignedIntToBinaryString(num int, num_bytes int) (string, error) {
	if num < 0 {
		return "", fmt.Errorf("cannot convert negative number to binary")
	} else if num_bytes < 1 {
		return "", fmt.Errorf("cannot convert to binary with less than 1 byte")
	} else if num >= 1<<(num_bytes*8) {
		return "", fmt.Errorf("number is too large to fit in %d bytes", num_bytes)
	}
	bytes, _ := LittleEndianUnsignedIntToBinary(num, num_bytes)
	var bs = ""
	for i := 0; i < num_bytes; i++ {
		bs += fmt.Sprintf("%08b", bytes[i])
	}
	return bs, nil
}

// LittleEndianUnsignedIntToBinary converts an unsigned integer to bytes in little endian
// format. The input num will be packed into a number of bytes specified by length. If num
// is too large to fit in length bytes, an error will be returned.
//
// The function returns a slice of bytes representing the binary.
func LittleEndianUnsignedIntToBinary(num int, num_bytes int) ([]byte, error) {
	if num < 0 {
		return nil, fmt.Errorf("cannot convert negative number to binary")
	} else if num_bytes < 1 {
		return nil, fmt.Errorf("cannot convert to binary with less than 1 byte")
	} else if num >= 1<<(num_bytes*8) {
		return nil, fmt.Errorf("number is too large to fit in %d bytes", num_bytes)
	}
	var result []byte
	for i := 0; i < num_bytes; i++ {
		result = append(result, byte(num>>uint(i*8)))
	}
	return result, nil
}

// LittleEndianSignedIntToBinaryString converts a signed integer to a binary string in little
// endian format. The input num will be packed into a number of bytes specified by length. If
// num is too large to fit in length bytes, an error will be returned.
//
// The function returns a string of 0s and 1s representing the binary.
func LittleEndianSignedIntToBinaryString(num int, num_bytes int) (string, error) {
	if num_bytes < 1 {
		return "", fmt.Errorf("cannot convert to binary with less than 1 byte")
	}
	minValue := -1 << ((num_bytes * 8) - 1)
	maxValue := (1 << ((num_bytes * 8) - 1)) - 1
	if num < minValue || num > maxValue {
		return "", fmt.Errorf("number is too large to fit in %d bytes", num_bytes)
	}
	bytes, _ := LittleEndianSignedIntToBinary(num, num_bytes)
	var bs = ""
	println()
	for i := 0; i < num_bytes; i++ {
		bs += fmt.Sprintf("%08b", bytes[i])
		fmt.Printf("%02x ", bytes[i])
	}
	println()
	return bs, nil
}

// LittleEndianSignedIntToBinary converts a signed integer to bytes in little endian format.
// The input num will be packed into a number of bytes specified by length. If num is too large
// to fit in length bytes, an error will be returned.
//
// The function returns a slice of bytes representing the binary.
func LittleEndianSignedIntToBinary(num int, num_bytes int) ([]byte, error) {
	if num_bytes < 1 {
		return nil, fmt.Errorf("cannot convert to binary with less than 1 byte")
	}
	minValue := -1 << ((num_bytes * 8) - 1)
	maxValue := (1 << ((num_bytes * 8) - 1)) - 1
	if num < minValue || num > maxValue {
		return nil, fmt.Errorf("number is too large to fit in %d bytes", num_bytes)
	}
	var result []byte
	for i := 0; i < num_bytes; i++ {
		shift := uint(i * 8)
		result = append(result, byte(num>>shift))
	}
	return result, nil
}

// LittleEndianBytesToUnsignedInt converts bytes to an unsigned integer in little endian format.
//
// The function returns the integer value of the bytes.
func LittleEndianBytesToUnsignedInt(bytes []byte) int {
	var result int
	for i := 0; i < len(bytes); i++ {
		result += int(bytes[i]) << uint(i*8)
	}
	return result
}

// LittleEndianBytesToSignedInt converts bytes to a signed integer in little endian format.
//
// The function returns the integer value of the bytes.
func LittleEndianBytesToSignedInt(bytes []byte) int {
	var result int
	if bytes[len(bytes)-1] >= 128 {
		result = -1 << uint((len(bytes)-1)*8)
	}
	for i := 0; i < len(bytes); i++ {
		result += int(bytes[i]) << uint(i*8)
	}
	return result
}
