package moscore

type Bus interface {
	Read(uint16) uint8
	Write(uint16, uint8)
}

type Core struct {
	acc    uint8
	idx    uint8
	idy    uint8
	sp     uint8
	pc     uint16
	p      uint8
	halted bool
	bus    Bus
}

func New(bus Bus) *Core {
	return &Core{bus: bus}
}

func (core *Core) Reset() {
	core.acc = 0
	core.idx = 0
	core.idy = 0
	core.sp = 0xff
	core.p = 0

	// 6502 reset vector
	pcl := core.bus.Read(0xfffc)
	pch := core.bus.Read(0xfffd)
	core.pc = addrFromBytes(pcl, pch)
	core.halted = false
}

func (core *Core) Step() {
	opcode := core.fetch()
	core.decode(opcode)
}

func (core *Core) fetch() uint8 {
	byte := core.bus.Read(core.pc)
	core.pc++
	return byte
}

func (core *Core) push(byte uint8) {
	core.sp--
	addr := addrFromBytes(core.sp, 0x01)
	core.bus.Write(addr, byte)
}

func (core *Core) pull() uint8 {
	addr := addrFromBytes(core.sp, 0x01)
	core.sp++
	return core.bus.Read(addr)
}

func (core *Core) decode(opcode uint8) {
	switch opcode {

	// ASL, zp
	case 0x06:
		core.aslZeroPage(nil)

	// PHP
	case 0x08:
		core.push(core.p)

	// ASL, A
	case 0x0A:
		core.aslAccumulator()

	// ASL, abs
	case 0x0E:
		core.aslAbsolute(nil)

	// BPL
	case 0x10:
		core.branchIfPositive()

	// ASL, zp X
	case 0x16:
		core.aslZeroPage(&core.idx)

	// CLC
	case 0x18:
		core.setFlag(Carry, false)

	// ASL, abs X
	case 0x1E:
		core.aslAbsolute(&core.idx)

	// AND, ind X
	case 0x21:
		core.andIndexedIndirect()

	// BIT, zp
	case 0x24:
		core.bitZeroPage()

	// AND, zp
	case 0x25:
		core.andZeroPage(nil)

	// PLP
	case 0x28:
		core.p = core.pull()

	// AND, byte
	case 0x29:
		core.andImmediate()

	// BIT, abs
	case 0x2C:
		core.bitAbsolute()

	// AND, abs
	case 0x2D:
		core.andAbsolute(nil)

	// BMI
	case 0x30:
		core.branchIfMinus()

	// AND, ind Y
	case 0x31:
		core.andIndirectIndexed()

	// AND, zp X
	case 0x35:
		core.andZeroPage(&core.idx)

	// SEC
	case 0x38:
		core.setFlag(Carry, true)

	// AND, abs Y
	case 0x39:
		core.andAbsolute(&core.idy)

	// AND, abs X
	case 0x3D:
		core.andAbsolute(&core.idx)

	// PHA
	case 0x48:
		core.push(core.acc)

	// BVC
	case 0x50:
		core.branchNotOverflow()

	// CLI
	case 0x58:
		core.setFlag(IRQD, false)

	// ADC, zp
	case 0x65:
		core.addZeroPage(nil)

	// PLA
	case 0x68:
		core.acc = core.pull()

	// ADC, byte
	case 0x69:
		core.addImmediate()

	// ADC, abs
	case 0x6D:
		core.addAbsolute(nil)

	// BVS
	case 0x70:
		core.branchIfOverflow()

	// ADC, zp X
	case 0x75:
		core.addZeroPage(&core.idx)

	// SEI
	case 0x78:
		core.setFlag(IRQD, true)

	// ADC, abs Y
	case 0x79:
		core.addAbsolute(&core.idy)

	// ADC, abs X
	case 0x7D:
		core.addAbsolute(&core.idx)

	// STA, ind X
	case 0x81:
		addr := core.getIndexedIndirectAddr()
		core.bus.Write(addr, core.acc)

	// STY, zp
	case 0x84:
		addr := core.getZeroPageAddr(nil)
		core.bus.Write(addr, core.idy)

	// STA, zp
	case 0x85:
		addr := core.getZeroPageAddr(nil)
		core.bus.Write(addr, core.acc)

	// STX, zp
	case 0x86:
		addr := core.getZeroPageAddr(nil)
		core.bus.Write(addr, core.idx)

	// TXA
	case 0x8A:
		core.acc = core.idx
		core.setNZ(core.acc)

	// STY, abs
	case 0x8C:
		addr := core.getAbsoluteAddr(nil)
		core.bus.Write(addr, core.idy)

	// STA, abs
	case 0x8D:
		addr := core.getAbsoluteAddr(nil)
		core.bus.Write(addr, core.acc)

	// STX, abs
	case 0x8E:
		addr := core.getAbsoluteAddr(nil)
		core.bus.Write(addr, core.idx)

	// BCC
	case 0x90:
		core.branchCarryClear()

	// STA, ind Y
	case 0x91:
		addr := core.getIndirectIndexedAddr()
		core.bus.Write(addr, core.acc)

	// STY, zp X
	case 0x94:
		addr := core.getZeroPageAddr(&core.idx)
		core.bus.Write(addr, core.idy)

	// STA, zp X
	case 0x95:
		addr := core.getZeroPageAddr(&core.idx)
		core.bus.Write(addr, core.acc)

	// STX, zp Y
	case 0x96:
		addr := core.getZeroPageAddr(&core.idy)
		core.bus.Write(addr, core.idx)

	// TYA
	case 0x98:
		core.acc = core.idy
		core.setNZ(core.acc)

	// STA, abs Y
	case 0x99:
		addr := core.getAbsoluteAddr(&core.idy)
		core.bus.Write(addr, core.acc)

	// TXS
	case 0x9A:
		core.sp = core.idx

	// STA, abs X
	case 0x9D:
		addr := core.getAbsoluteAddr(&core.idx)
		core.bus.Write(addr, core.acc)

	// LDY, byte
	case 0xA0:
		core.loadImmediate(&core.idy)

	// LDA, ind X
	case 0xA1:
		core.loadIndexedIndirect(&core.acc)

	// LDX, byte
	case 0xA2:
		core.loadImmediate(&core.idx)

	// LDY, zp
	case 0xA4:
		core.loadZeroPage(&core.idy, nil)

	// LDA, zp
	case 0xA5:
		core.loadZeroPage(&core.acc, nil)

	// LDX, zp
	case 0xA6:
		core.loadZeroPage(&core.idx, nil)

	// TAY
	case 0xA8:
		core.idy = core.acc
		core.setNZ(core.idy)

	// LDA, byte
	case 0xA9:
		core.loadImmediate(&core.acc)

	// TAX
	case 0xAA:
		core.idx = core.acc
		core.setNZ(core.idx)

	// LDY, abs
	case 0xAC:
		core.loadAbsolute(&core.idy, nil)

	// LDA, abs
	case 0xAD:
		core.loadAbsolute(&core.acc, nil)

	// LDX, abs
	case 0xAE:
		core.loadAbsolute(&core.idx, nil)

	// BCS
	case 0xB0:
		core.branchCarrySet()

	// LDA, ind Y
	case 0xB1:
		core.loadIndirectIndexed(&core.acc)

	// LDY, zp X
	case 0xB4:
		core.loadZeroPage(&core.idy, &core.idx)

	// LDA, zp X
	case 0xB5:
		core.loadZeroPage(&core.acc, &core.idx)

	// LDX, zp Y
	case 0xB6:
		core.loadZeroPage(&core.idx, &core.idy)

	// CLV
	case 0xB8:
		core.setFlag(Overflow, false)

	// LDA, abs Y
	case 0xB9:
		core.loadAbsolute(&core.acc, &core.idy)

	// TSX
	case 0xBA:
		core.idx = core.sp
		core.setNZ(core.idx)

	// LDY, abs X
	case 0xBC:
		core.loadAbsolute(&core.idy, &core.idx)

	// LDA, abs X
	case 0xBD:
		core.loadAbsolute(&core.acc, &core.idx)

	// LDX, abs Y
	case 0xBE:
		core.loadAbsolute(&core.idx, &core.idy)

	// BNE
	case 0xD0:
		core.branchNotEqual()

	// CLD
	case 0xD8:
		core.setFlag(Decimal, false)

	// BEQ
	case 0xF0:
		core.branchIfEqual()

	// SED
	case 0xF8:
		core.setFlag(Decimal, true)
	}
}
