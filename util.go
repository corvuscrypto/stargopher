package stargopher

//ReadVarint takes a byte array representing a starbound VLQ and returns an int64
func ReadVarint(buf []byte) (int64, int) {
	var value int64
	var count int
	for i := 0; i < len(buf); i++ {
		value = (value << 7) | (int64(buf[i]) & 0x7f)
		if (buf[i] & 0x80) == 0 {
			break
		}
		count++
	}
	if (value & 1) == 0x00 {
		return value >> 1, count
	}
	return -((value >> 1) + 1), count
}

//WriteVarint takes a byte array representing a starbound VLQ and returns an int64
func WriteVarint(num int64) []byte {
	if num == 0 {
		return []byte{0}
	}
	var b []byte
	for num > 0 {
		var tmp = num & 0x7F
		num >>= 7
		if num != 0 {
			tmp |= 0x80
		}
		b = append([]byte{byte(tmp)}, b...)
	}
	if len(b) > 1 {
		b[0] |= 0x80
		b[len(b)-1] ^= 0x80
	}
	return b
}
