package moscore

func (core *Core) bit(data uint8) {
	result := core.acc & data

	core.setFlag(Zero, result == 0)
	core.setFlag(Negative, (data&Negative) != 0)
	core.setFlag(Overflow, (data&Overflow) != 0)
}

func (core *Core) compare(register uint8, data uint8) {
	result := register - data

	core.setFlag(Negative, (result&Negative) != 0)
	core.setFlag(Carry, register >= data)
	core.setFlag(Zero, register == data)
}
