package stargopher

import (
	cr "crypto/rand"
	"fmt"
)

func newUUID() string {
	b := make([]byte, 16)
	cr.Read(b)
	b[7] = (b[7] | 0x40) & 0xBF
	b[9] = (b[9] | 0x80) & 0x3F
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
