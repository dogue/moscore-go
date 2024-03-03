package moscore

func (core *Core) branchCarryClear() {
	offset := core.getImmediate()

	if core.getFlag(Carry) {
		return
	}

	core.pc += uint16(offset)
}

func (core *Core) branchCarrySet() {
	offset := core.getImmediate()

	if !core.getFlag(Carry) {
		return
	}

	core.pc += uint16(offset)
}

func (core *Core) branchIfEqual() {
	offset := core.getImmediate()

	if !core.getFlag(Zero) {
		return
	}

	core.pc += uint16(offset)
}