package moscore

func (core *Core) addImmediate() {
	byte := core.getImmediate()
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}

func (core *Core) addZeroPage(index Register) {
	byte := core.getZeroPage(index)
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}

func (core *Core) addAbsolute(index Register) {
	byte := core.getAbsolute(index)
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}

func (core *Core) addIndexedIndirect() {
	byte := core.getIndexedIndirect()
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}

func (core *Core) addIndirectIndexed() {
	byte := core.getIndirectIndexed()
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}
