package moscore

// (low byte, high byte) -> 16-bit address
func addrFromBytes(low, high uint8) uint16 {
	return (uint16(high) << 8) | uint16(low)
}

// 16-bit address -> (low byte, high byte)
func bytesFromAddr(addr uint16) (uint8, uint8) {
	low := uint8(addr)
	high := uint8(addr >> 8)
	return low, high
}

// add two uint8's with carry
// ripped off from standard library code for doing with with 32 bit ints
func add8(x, y, carryIn uint8) (sum, carryOut uint8) {
	sum16 := uint16(x) + uint16(y) + uint16(carryIn)
	sum = uint8(sum16)
	carryOut = uint8(sum16 >> 8)
	return
}
