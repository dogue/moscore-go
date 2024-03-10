package moscore

type Flag = uint8

const (
	Carry    Flag = 1
	Zero          = 2
	IRQD          = 4
	Decimal       = 8
	Break         = 16
	Unused        = 32
	Overflow      = 64
	Negative      = 128
)

func (core *Core) setFlag(bit Flag, cond bool) {
	if cond {
		core.p |= bit
	} else {
		core.p &= ^bit
	}
}

func (core *Core) getFlag(bit Flag) bool {
	return core.p&bit != 0
}

func (core *Core) getFlagUint8(bit Flag) uint8 {
	return core.p & bit
}

// setting both the negative and zero flags is common enough to warrant this helper
func (core *Core) setNZ(byte uint8) {
	core.setFlag(Negative, (byte&0x80) != 0)
	core.setFlag(Zero, byte == 0)
}
