package conversion

import (
	"encoding/binary"
)

func UInt64ToBytes(i uint64) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, i)
	return bs
}
