package conversion

import (
	"fmt"
	"testing"
)

func TestUInt64ToBytes(t *testing.T) {
	i := int64(-123456789)
	fmt.Printf("%x\n", UInt64ToBytes(uint64(i)))

}
