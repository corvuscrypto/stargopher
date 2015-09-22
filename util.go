package stargopher

//Varint takes a byte array representing a starbound VLQ and returns an int64
func Varint(buf []byte) int64 {
	var x int64
	for i, b := range buf {

		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				return 0
			}
			x = x<<6 | int64((b>>1)&0x3f)
			if b&1 == 0 {
				return x << 1
			}
			return -((x << 1) + 1)
		}
		x = x<<7 | int64(b&0x7f)
	}
	return 0
}
