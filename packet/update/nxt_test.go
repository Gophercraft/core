package update

import (
	"fmt"
	"testing"
)

func _nxtBit(chunk, bit *uint32) {
	fmt.Println("Bit", *chunk, *bit)
	nxtBit(chunk, bit)
}

func _nxtByte(chunk, bit *uint32) {
	fmt.Println("Byte", *chunk, *bit)
	nxtByte(chunk, bit)
}

func TestNxt(t *testing.T) {
	var chunk, bit uint32

	_nxtBit(&chunk, &bit)
	_nxtBit(&chunk, &bit)
	_nxtBit(&chunk, &bit)
	_nxtBit(&chunk, &bit)
	_nxtByte(&chunk, &bit)
	_nxtByte(&chunk, &bit)

}
