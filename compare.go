package moscore

func (core *Core) bitZeroPage() {
	byte := core.getZeroPageByte(nil)
	result := core.acc & byte

	core.setFlag(Zero, result == 0)
	core.setFlag(Negative, (byte&(1<<Negative) != 0))
	core.setFlag(Overflow, (byte&(1<<Overflow) != 0))
}

func (core *Core) bitAbsolute() {
	byte := core.getAbsoluteByte(nil)
	result := core.acc & byte

	core.setNZ(result)
	core.setFlag(Overflow, (result&(1<<6) != 0))
}
