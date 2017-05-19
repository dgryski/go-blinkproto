package blinkproto

func getDataSize(data uint64) int {
	mask := ^uint64(0) >> 8
	var p1 int
	for p1 = 8; p1 > 1; p1-- {
		if data > mask {
			break
		}
		mask >>= 8
	}
	return p1
}
func EncodeU64(buf []byte, data uint64) int {
	if data < 0x4000 {
		if data < 0x80 {
			buf[0] = byte(data)
			return 1
		}
		buf[0] = byte(128 | (data & 0x3f))
		buf[1] = byte(data >> 6)
		return 2
	}

	size := getDataSize(data)
	buf[0] = byte(192 | size)
	for p1 := 1; p1 <= size; p1++ {
		buf[p1] = byte(data)
		data >>= 8
	}
	return size + 1
}

func DecodeU64(buf []byte, data *uint64) int {
	if buf[0] < 192 {
		if buf[0] < 128 {
			*data = uint64(buf[0])
			return 1
		}
		*data = uint64(buf[0]&63) | uint64(buf[1]<<6)
		return 2

	}

	var size = int(buf[0] & 0x3f)
	var temp uint64
	for p1 := 0; p1 < size; p1++ {
		temp |= uint64(buf[p1+1]) << uint(8*p1)
	}
	*data = temp
	return size + 1
}
