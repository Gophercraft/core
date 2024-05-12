package srp

import (
	"bytes"
	"crypto/rand"
	"io"
	"math/big"
	"slices"
)

type Int struct {
	x big.Int
}

func (x *Int) ModExp(y, m *Int) *Int {
	i := new(Int)
	i.x.Exp(&x.x, &y.x, &m.x)
	return i
}

func (x *Int) Add(y *Int) *Int {
	i := new(Int)
	i.x.Add(&x.x, &y.x)
	return i
}

func (x *Int) Subtract(y *Int) *Int {
	i := new(Int)
	i.x.Sub(&x.x, &y.x)
	return i
}

func (x *Int) Multiply(y *Int) *Int {
	i := new(Int)
	i.x.Mul(&x.x, &y.x)
	return i
}

func (x *Int) Divide(y *Int) *Int {
	i := new(Int)
	i.x.Div(&x.x, &y.x)
	return i
}

func (x *Int) Equals(y *Int) bool {
	return bytes.Equal(x.x.Bytes(), y.x.Bytes())
}

func (x *Int) Mod(y *Int) *Int {
	i := new(Int)
	i.x.Mod(&x.x, &y.x)
	return i
}

// Serialization

func (x *Int) Len() int {
	bit_len := x.x.BitLen()
	return (bit_len + 7) / 8
}

func (x *Int) GetBytes(b []byte, little_endian bool) {
	bytes := x.x.Bytes()
	// convert to little endian
	slices.Reverse(bytes)

	copy(b, bytes)

	if !little_endian {
		// convert back to big endian
		slices.Reverse(b)
	}
}

func (x *Int) Bytes() (little_endian_bytes []byte) {
	little_endian_bytes = make([]byte, x.Len())
	x.GetBytes(little_endian_bytes, true)
	return
}

func (x *Int) BytesBigEndian() (big_endian_bytes []byte) {
	big_endian_bytes = make([]byte, x.Len())
	x.GetBytes(big_endian_bytes, false)
	return
}

// Returns a constant-size slice representing the little-endian value of x
// if x = 1 and size = 8 the result would be
// 1 0 0 0 0 0 0
func (x *Int) Array(size int) (little_endian_byte_array []byte) {
	little_endian_byte_array = make([]byte, size)
	x.GetBytes(little_endian_byte_array, true)
	return
}

// Returns a constant-size slice representing the big-endian value of x
// if x = 1 and size = 8 the result would be
// 0 0 0 0 0 0 1
func (x *Int) ArrayBigEndian(size int) (big_endian_byte_array []byte) {
	big_endian_byte_array = make([]byte, size)
	x.GetBytes(big_endian_byte_array, false)
	return
}

func (x *Int) Slice(at_least int) []byte {
	if x.Len() >= at_least {
		return x.Bytes()
	} else {
		return x.Array(at_least)
	}
}

func (x *Int) SliceBigEndian(at_least int) []byte {
	if x.Len() >= at_least {
		return x.BytesBigEndian()
	} else {
		return x.ArrayBigEndian(at_least)
	}
}

func (x *Int) String() string {
	return x.x.String()
}

func (x *Int) Copy() *Int {
	i := new(Int)
	i.x.Set(&x.x)
	return i
}

func NewInt(i64 int64) *Int {
	i := new(Int)
	i.x.SetInt64(i64)
	return i
}

func NewIntFromBytesBigEndian(big_endian_bytes []byte) *Int {
	i := new(Int)
	i.x.SetBytes(big_endian_bytes)
	return i
}

func NewIntFromBytes(little_endian_bytes []byte) *Int {
	big_endian_bytes := reverse_bytes(little_endian_bytes)
	return NewIntFromBytesBigEndian(big_endian_bytes)
}

func NewIntFromHexString(hex_str string) *Int {
	i := new(Int)
	i.x.SetString(hex_str, 16)
	return i
}

// Big-Endian.
func NewIntFromString(base10 string) *Int {
	i := new(Int)
	_, ok := i.x.SetString(base10, 10)
	if !ok {
		panic("failed to set string")
	}
	return i
}

func NewRandomInt(size int) *Int {
	bytes := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		panic(err)
	}
	return NewIntFromBytes(bytes)
}

func (i *Int) HexString() string {
	return i.x.Text(16)
}

func (i *Int) Bits() int {
	return i.x.BitLen()
}
