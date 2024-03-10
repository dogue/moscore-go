package moscore

func (core *Core) adc(data uint8) {
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, data, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}
