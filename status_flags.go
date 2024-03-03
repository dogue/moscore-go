package moscore

type Flag int

const (
	Carry Flag = iota
	Zero
	IRQD
	Decimal
	Break
	Overflow
	Negative
)

func (core *Core) setFlag(bit Flag, cond bool) {
	if cond {
		core.p |= 1 << bit
	} else {
		core.p &= ^uint8(1) << bit
	}
}

func (core *Core) getFlag(bit Flag) bool {
	return core.p&(1<<bit) != 0
}

func (core *Core) getFlagUint8(bit Flag) uint8 {
	return core.p & (1 << bit)
}

// setting both the negative and zero flags is common enough to warrant this helper
func (core *Core) setNZ(byte uint8) {
	core.setFlag(Negative, (byte&0x80) != 0)
	core.setFlag(Zero, byte == 0)
}
