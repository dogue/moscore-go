package moscore

func (core *Core) branchCarryClear() {
	offset := core.getImmediateByte()

	if core.getFlag(Carry) {
		return
	}

	core.pc += uint16(offset)
}

func (core *Core) branchCarrySet() {
	offset := core.getImmediateByte()

	if !core.getFlag(Carry) {
		return
	}

	core.pc += uint16(offset)
}

func (core *Core) branchIfEqual() {
	offset := core.getImmediateByte()

	if !core.getFlag(Zero) {
		return
	}

	core.pc += uint16(offset)
}

func (core *Core) branchIfMinus() {
	offset := core.getImmediateByte()

	if !core.getFlag(Negative) {
		return
	}

	core.pc += uint16(offset)
}

func (core *Core) branchNotEqual() {
	offset := core.getImmediateByte()

	if core.getFlag(Zero) {
		return
	}

	core.pc += uint16(offset)
}

func (core *Core) branchIfPositive() {
	offset := core.getImmediateByte()

	if core.getFlag(Negative) {
		return
	}

	core.pc += uint16(offset)
}

func (core *Core) branchNotOverflow() {
	offset := core.getImmediateByte()

	if core.getFlag(Overflow) {
		return
	}

	core.pc += uint16(offset)
}

func (core *Core) branchIfOverflow() {
	offset := core.getImmediateByte()

	if !core.getFlag(Overflow) {
		return
	}

	core.pc += uint16(offset)
}
