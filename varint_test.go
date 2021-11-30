package varint

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"
)

func TestVarInt(t *testing.T) {
	var buf [16]byte
	for bits := 0; bits <= 64; bits++ {
		var value uint64 = 0xAAAAAAAAAAAAAAAA >> (64 - bits)
		t.Run(fmt.Sprintf("bits_%d-", bits), func(t *testing.T) {
			n1 := Encode(buf[:], value)
			out, n2 := Decode(buf[:])
			if value != out {
				t.Fatalf("bits %d, expect value %X, got value %X", bits, value, out)
			}
			if n1 != n2 {
				t.Fatalf("bits %d, expect length %d, got length %d", bits, n1, n2)
			}
		})
	}
}

var out int
var buf0 = make([]byte, 1024)
var buf = buf0[:rand.Int()%1010]

func BenchmarkEncode(b *testing.B) {
	for nBits := 7; nBits < 64; nBits += 7 {
		var value uint64 = 0xAAAAAAAAAAAAAAAA >> (64 - nBits)
		b.Run(fmt.Sprintf("bits=%d", nBits), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				out = Encode(buf, value)
			}
		})
	}
}

func BenchmarkDecode(b *testing.B) {
	for nBits := 7; nBits < 64; nBits += 7 {
		var value uint64 = 0xAAAAAAAAAAAAAAAA >> (64 - nBits)
		Encode(buf, value)
		b.Run(fmt.Sprintf("bits=%d", nBits), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, out = Decode(buf)
			}
		})
	}
}

func BenchmarkStdEncode(b *testing.B) {
	for nBits := 7; nBits < 64; nBits += 7 {
		var value uint64 = 0xAAAAAAAAAAAAAAAA >> (64 - nBits)
		b.Run(fmt.Sprintf("bits=%d", nBits), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				out = binary.PutUvarint(buf, value)
			}
		})
	}
}

func BenchmarkStdDecode(b *testing.B) {
	for nBits := 7; nBits < 64; nBits += 7 {
		var value uint64 = 0xAAAAAAAAAAAAAAAA >> (64 - nBits)
		binary.PutUvarint(buf, value)
		b.Run(fmt.Sprintf("bits=%d", nBits), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, out = binary.Uvarint(buf)
			}
		})
	}
}
