package varint

// Encode v as a prefixed with 1 bits variable length integer into b.
// Returns the number of bytes written or 0 if b is too small.
// The encoded integer is at most 9 bytes long. This function doesn't
// panic.
func Encode(b []byte, v uint64) int {
	switch {
	case v < 0x80:
		if len(b) != 0 {
			b[0] = byte(v)
			return 1
		}
	case v < 1<<14:
		if len(b) >= 2 {
			b[0] = byte(v>>8) | 0x80
			b[1] = byte(v)
			return 2
		}
	case v < 1<<21:
		if len(b) >= 3 {
			b[0] = byte(v>>16) | 0xC0
			b[1] = byte(v >> 8)
			b[2] = byte(v)
			return 3
		}
	case v < 1<<28:
		if len(b) >= 4 {
			w := uint32(v) | 0xE0000000
			b[0] = byte(w >> 24)
			b[1] = byte(w >> 16)
			b[2] = byte(w >> 8)
			b[3] = byte(w)
			return 4
		}
	case v < 1<<35:
		if len(b) >= 5 {
			b[0] = byte(v>>32) | 0xF0
			b[1] = byte(v >> 24)
			b[2] = byte(v >> 16)
			b[3] = byte(v >> 8)
			b[4] = byte(v)
			return 5
		}
	case v < 1<<42:
		if len(b) >= 6 {
			b[0] = byte(v>>40) | 0xF8
			b[1] = byte(v >> 32)
			b[2] = byte(v >> 24)
			b[3] = byte(v >> 16)
			b[4] = byte(v >> 8)
			b[5] = byte(v)
			return 6
		}
	case v < 1<<49:
		if len(b) >= 7 {
			b[0] = byte(v>>48) | 0xFC
			b[1] = byte(v >> 40)
			b[2] = byte(v >> 32)
			b[3] = byte(v >> 24)
			b[4] = byte(v >> 16)
			b[5] = byte(v >> 8)
			b[6] = byte(v)
			return 7
		}
	case v < 1<<56:
		if len(b) >= 8 {
			b[0] = 0xFE
			b[1] = byte(v >> 48)
			b[2] = byte(v >> 40)
			b[3] = byte(v >> 32)
			b[4] = byte(v >> 24)
			b[5] = byte(v >> 16)
			b[6] = byte(v >> 8)
			b[7] = byte(v)
			return 8
		}
	default:
		if len(b) >= 9 {
			b[0] = 0xFF
			b[1] = byte(v >> 56)
			b[2] = byte(v >> 48)
			b[3] = byte(v >> 40)
			b[4] = byte(v >> 32)
			b[5] = byte(v >> 24)
			b[6] = byte(v >> 16)
			b[7] = byte(v >> 8)
			b[8] = byte(v)
			return 9
		}
	}
	return 0
}

// Decode returns the varying length integer in front of b and the
// number of bytes read or zero if b is empty or the integer is
// truncated. This function doesn't panic.
func Decode(b []byte) (uint64, int) {
	lb := len(b)
	if lb > 0 {
		hdr := b[0]
		switch {
		case hdr <= 0x7F:
			return uint64(hdr), 1
		case hdr <= 0xBF:
			if lb >= 2 {
				return uint64(hdr&0x3F)<<8 | uint64(b[1]), 2
			}
		case hdr <= 0xDF:
			if lb >= 3 {
				return uint64(hdr&0x1F)<<16 | uint64(b[1])<<8 | uint64(b[2]), 3
			}
		case hdr <= 0xEF:
			if lb >= 4 {
				return uint64(hdr&0x0F)<<24 | uint64(b[1])<<16 | uint64(b[2])<<8 | uint64(b[3]), 4
			}
		case hdr <= 0xF7:
			if lb >= 5 {
				return uint64(hdr&0x07)<<32 | uint64(b[1])<<24 | uint64(b[2])<<16 | uint64(b[3])<<8 |
					uint64(b[4]), 5
			}
		case hdr <= 0xFB:
			if lb >= 6 {
				return uint64(hdr&0x03)<<40 | uint64(b[1])<<32 | uint64(b[2])<<24 | uint64(b[3])<<16 |
					uint64(b[4])<<8 | uint64(b[5]), 6
			}
		case hdr <= 0xFD:
			if lb >= 7 {
				return uint64(hdr&0x01)<<48 | uint64(b[1])<<40 | uint64(b[2])<<32 | uint64(b[3])<<24 |
					uint64(b[4])<<16 | uint64(b[5])<<8 | uint64(b[6]), 7
			}
		case hdr <= 0xFE:
			if lb >= 8 {
				return uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 |
					uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7]), 8
			}
		default:
			if lb >= 9 {
				return uint64(b[1])<<56 | uint64(b[2])<<48 | uint64(b[3])<<40 | uint64(b[4])<<32 |
					uint64(b[5])<<24 | uint64(b[6])<<16 | uint64(b[7])<<8 | uint64(b[8]), 9
			}
		}
	}
	return 0, 0
}
