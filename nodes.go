package mapache

import "fmt"

// Node is a type to represent a vehicle node. It is a slice of Signals,
// where each Signal represents a specific item in the vehicle node.
type Node []Signal

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

// FillFromBytes fills the fields of a Node with the bytes provided in the data slice.
// It decodes the bytes into integer values and stores them in the RawValue field of each Signal.
// It returns an error if the length of the data slice does not match the total size of the Node.
func (n Node) FillFromBytes(data []byte) error {
	if len(data) != n.Size() {
		return fmt.Errorf("invalid data length, expected %d bytes", n.Size())
	}
	counter := 0
	for i, signal := range n {
		signal.Bytes = data[counter : counter+signal.Size]
		counter += signal.Size
		n[i] = signal.Decode()
	}
	return nil
}

// FillFromInts fills the fields of a Node with the integer values provided in the ints slice.
// It encodes the integer values into bytes and stores them in the Bytes field of each Signal.
// It returns an error if the length of the ints slice does not match the number of fields in the Node.
func (n Node) FillFromInts(ints []int) error {
	if len(ints) != n.Length() {
		return fmt.Errorf("invalid ints length, expected %d ints", n.Length())
	}
	for i, signal := range n {
		signal.RawValue = ints[i]
		signal, err := signal.Encode()
		if err != nil {
			return err
		}
		n[i] = signal
	}
	return nil
}
