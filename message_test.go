package mapache

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestTimestamp(t *testing.T) {
	now := time.Now().UnixMicro()
	fmt.Println(now)
	fmt.Printf("Size of timestamp: %d bytes\n", unsafe.Sizeof(now))
}
