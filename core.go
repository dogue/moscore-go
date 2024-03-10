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
		core.aslZeroPage(0)

	// PHP
	case 0x08:
		core.push(core.p)

	// ASL, A
	case 0x0A:
		core.aslAccumulator()

	// ASL, abs
	case 0x0E:
		core.aslAbsolute(0)

	// BPL
	case 0x10:
		core.branchIfPositive()

	// ASL, zp X
	case 0x16:
		core.aslZeroPage(core.idx)

	// CLC
	case 0x18:
		core.setFlag(Carry, false)

	// ASL, abs X
	case 0x1E:
		core.aslAbsolute(core.idx)

	// AND, ind X
	case 0x21:
		data := core.getIndirectXByte()
		core.acc &= data
		core.setNZ(core.acc)

	// BIT, zp
	case 0x24:
		data := core.getZeroPageByte(0)
		core.bit(data)

	// AND, zp
	case 0x25:
		data := core.getZeroPageByte(0)
		core.acc &= data
		core.setNZ(core.acc)

	// PLP
	case 0x28:
		core.p = core.pull()

	// AND, byte
	case 0x29:
		data := core.fetch()
		core.acc &= data
		core.setNZ(core.acc)

	// BIT, abs
	case 0x2C:
		data := core.getAbsoluteByte(0)
		core.bit(data)

	// AND, abs
	case 0x2D:
		data := core.getAbsoluteByte(0)
		core.acc &= data
		core.setNZ(core.acc)

	// BMI
	case 0x30:
		core.branchIfMinus()

	// AND, ind Y
	case 0x31:
		data := core.getIndirectYByte()
		core.acc &= data
		core.setNZ(core.acc)

	// AND, zp X
	case 0x35:
		data := core.getZeroPageByte(core.idx)
		core.acc &= data
		core.setNZ(core.acc)

	// SEC
	case 0x38:
		core.setFlag(Carry, true)

	// AND, abs Y
	case 0x39:
		data := core.getAbsoluteByte(core.idy)
		core.acc &= data
		core.setNZ(core.acc)

	// AND, abs X
	case 0x3D:
		data := core.getAbsoluteByte(core.idx)
		core.acc &= data
		core.setNZ(core.acc)

	// PHA
	case 0x48:
		core.push(core.acc)

	// JMP, abs
	case 0x4C:
		offset := core.getAbsoluteByte(0)
		core.pc += uint16(offset)

	// BVC
	case 0x50:
		core.branchNotOverflow()

	// CLI
	case 0x58:
		core.setFlag(IRQD, false)

	// ADC, zp
	case 0x65:
		data := core.getZeroPageByte(0)
		core.adc(data)

	// PLA
	case 0x68:
		core.acc = core.pull()

	// ADC, byte
	case 0x69:
		data := core.fetch()
		core.adc(data)

	// JMP, ind
	case 0x6C:
		addr := core.getAbsoluteAddr(0)
		targetLow := core.bus.Read(addr)
		targetHigh := core.bus.Read(addr + 1)
		targetAddr := addrFromBytes(targetLow, targetHigh)
		core.pc = targetAddr

	// ADC, abs
	case 0x6D:
		data := core.getAbsoluteByte(0)
		core.adc(data)

	// BVS
	case 0x70:
		core.branchIfOverflow()

	// ADC, zp X
	case 0x75:
		data := core.getZeroPageByte(core.idx)
		core.adc(data)

	// SEI
	case 0x78:
		core.setFlag(IRQD, true)

	// ADC, abs Y
	case 0x79:
		data := core.getAbsoluteByte(core.idy)
		core.adc(data)

	// ADC, abs X
	case 0x7D:
		data := core.getAbsoluteByte(core.idx)
		core.adc(data)

	// STA, ind X
	case 0x81:
		addr := core.getIndirectXAddr()
		core.bus.Write(addr, core.acc)

	// STY, zp
	case 0x84:
		addr := core.getZeroPageAddr(0)
		core.bus.Write(addr, core.idy)

	// STA, zp
	case 0x85:
		addr := core.getZeroPageAddr(0)
		core.bus.Write(addr, core.acc)

	// STX, zp
	case 0x86:
		addr := core.getZeroPageAddr(0)
		core.bus.Write(addr, core.idx)

	// DEY
	case 0x88:
		core.idy--

	// TXA
	case 0x8A:
		core.acc = core.idx
		core.setNZ(core.acc)

	// STY, abs
	case 0x8C:
		addr := core.getAbsoluteAddr(0)
		core.bus.Write(addr, core.idy)

	// STA, abs
	case 0x8D:
		addr := core.getAbsoluteAddr(0)
		core.bus.Write(addr, core.acc)

	// STX, abs
	case 0x8E:
		addr := core.getAbsoluteAddr(0)
		core.bus.Write(addr, core.idx)

	// BCC
	case 0x90:
		core.branchCarryClear()

	// STA, ind Y
	case 0x91:
		addr := core.getIndirectYAddr()
		core.bus.Write(addr, core.acc)

	// STY, zp X
	case 0x94:
		addr := core.getZeroPageAddr(core.idx)
		core.bus.Write(addr, core.idy)

	// STA, zp X
	case 0x95:
		addr := core.getZeroPageAddr(core.idx)
		core.bus.Write(addr, core.acc)

	// STX, zp Y
	case 0x96:
		addr := core.getZeroPageAddr(core.idy)
		core.bus.Write(addr, core.idx)

	// TYA
	case 0x98:
		core.acc = core.idy
		core.setNZ(core.acc)

	// STA, abs Y
	case 0x99:
		addr := core.getAbsoluteAddr(core.idy)
		core.bus.Write(addr, core.acc)

	// TXS
	case 0x9A:
		core.sp = core.idx

	// STA, abs X
	case 0x9D:
		addr := core.getAbsoluteAddr(core.idx)
		core.bus.Write(addr, core.acc)

	// LDY, byte
	case 0xA0:
		core.idy = core.fetch()
		core.setNZ(core.idy)

	// LDA, ind X
	case 0xA1:
		core.acc = core.getIndirectXByte()
		core.setNZ(core.acc)

	// LDX, byte
	case 0xA2:
		core.idx = core.fetch()
		core.setNZ(core.idx)

	// LDY, zp
	case 0xA4:
		core.idy = core.getZeroPageByte(0)
		core.setNZ(core.idy)

	// LDA, zp
	case 0xA5:
		core.acc = core.getZeroPageByte(0)
		core.setNZ(core.acc)

	// LDX, zp
	case 0xA6:
		core.idx = core.getZeroPageByte(0)
		core.setNZ(core.idx)

	// TAY
	case 0xA8:
		core.idy = core.acc
		core.setNZ(core.idy)

	// LDA, byte
	case 0xA9:
		core.acc = core.fetch()
		core.setNZ(core.acc)

	// TAX
	case 0xAA:
		core.idx = core.acc
		core.setNZ(core.idx)

	// LDY, abs
	case 0xAC:
		core.idy = core.getAbsoluteByte(0)
		core.setNZ(core.idy)

	// LDA, abs
	case 0xAD:
		core.acc = core.getAbsoluteByte(0)
		core.setNZ(core.acc)

	// LDX, abs
	case 0xAE:
		core.idx = core.getAbsoluteByte(0)
		core.setNZ(core.idx)

	// BCS
	case 0xB0:
		core.branchCarrySet()

	// LDA, ind Y
	case 0xB1:
		core.acc = core.getIndirectYByte()
		core.setNZ(core.acc)

	// LDY, zp X
	case 0xB4:
		core.idy = core.getZeroPageByte(core.idx)
		core.setNZ(core.idy)

	// LDA, zp X
	case 0xB5:
		core.acc = core.getZeroPageByte(core.idx)
		core.setNZ(core.acc)

	// LDX, zp Y
	case 0xB6:
		core.idx = core.getZeroPageByte(core.idy)
		core.setNZ(core.idx)

	// CLV
	case 0xB8:
		core.setFlag(Overflow, false)

	// LDA, abs Y
	case 0xB9:
		core.loadAbsolute(&core.acc, &core.idy)
		core.setNZ(core.acc)

	// TSX
	case 0xBA:
		core.idx = core.sp
		core.setNZ(core.idx)

	// LDY, abs X
	case 0xBC:
		core.idy = core.getAbsoluteByte(core.idx)
		core.setNZ(core.idy)

	// LDA, abs X
	case 0xBD:
		core.loadAbsolute(&core.acc, &core.idx)
		core.setNZ(core.acc)

	// LDX, abs Y
	case 0xBE:
		core.idx = core.getAbsoluteByte(core.idy)
		core.setNZ(core.idx)

	// CPY, byte
	case 0xC0:
		data := core.getImmediateByte()
		core.compare(core.idy, data)

	// CMP, ind X
	case 0xC1:
		data := core.getIndirectXByte()
		core.compare(core.acc, data)

	// CPY, zp
	case 0xC4:
		data := core.getZeroPageByte(0)
		core.compare(core.idy, data)

	// CMP, zp
	case 0xC5:
		data := core.getZeroPageByte(0)
		core.compare(core.acc, data)

	// DEC, zp
	case 0xC6:
		addr := core.getZeroPageAddr(0)
		data := core.bus.Read(addr)
		data--
		core.setNZ(data)
		core.bus.Write(addr, data)

	// INY
	case 0xC8:
		core.idy++

	// CMP, byte
	case 0xC9:
		data := core.getImmediateByte()
		core.compare(core.acc, data)

	// DEX
	case 0xCA:
		core.idx--

	// CPY, abs
	case 0xCC:
		data := core.getAbsoluteByte(0)
		core.compare(core.idy, data)

	// CMP, abs
	case 0xCD:
		data := core.getAbsoluteByte(0)
		core.compare(core.acc, data)

	// DEC, abs
	case 0xCE:
		addr := core.getAbsoluteAddr(0)
		data := core.bus.Read(addr)
		data--
		core.setNZ(data)
		core.bus.Write(addr, data)

	// BNE
	case 0xD0:
		core.branchNotEqual()

	// CMP, ind y
	case 0xD1:
		data := core.getIndirectYByte()
		core.compare(core.acc, data)

	// CMP, zp X
	case 0xD5:
		data := core.getZeroPageByte(core.idx)
		core.compare(core.acc, data)

	// DEC, zp X
	case 0xD6:
		addr := core.getZeroPageAddr(core.idx)
		data := core.bus.Read(addr)
		data--
		core.setNZ(data)
		core.bus.Write(addr, data)

	// CLD
	case 0xD8:
		core.setFlag(Decimal, false)

	// CMP, abs Y
	case 0xD9:
		data := core.getAbsoluteByte(core.idy)
		core.compare(core.acc, data)

	// CMP, abs X
	case 0xDD:
		data := core.getAbsoluteByte(core.idx)
		core.compare(core.acc, data)

	// DEC, abs X
	case 0xDE:
		addr := core.getAbsoluteAddr(core.idx)
		data := core.bus.Read(addr)
		data--
		core.setNZ(data)
		core.bus.Write(addr, data)

	// CPX, byte
	case 0xE0:
		data := core.getImmediateByte()
		core.compare(core.idx, data)

	// CPX, zp
	case 0xE4:
		data := core.getZeroPageByte(0)
		core.compare(core.idx, data)

	// INC, zp
	case 0xE6:
		addr := core.getZeroPageAddr(0)
		data := core.bus.Read(addr)
		data++
		core.setNZ(data)
		core.bus.Write(addr, data)

	// INX
	case 0xE8:
		core.idx++

	// CPX, abs
	case 0xEC:
		data := core.getAbsoluteByte(0)
		core.compare(core.idx, data)

	// INC, abs
	case 0xEE:
		addr := core.getAbsoluteAddr(0)
		data := core.bus.Read(addr)
		data++
		core.setNZ(data)
		core.bus.Write(addr, data)

	// BEQ
	case 0xF0:
		core.branchIfEqual()

	// INC, zp X
	case 0xF6:
		addr := core.getZeroPageAddr(core.idx)
		data := core.bus.Read(addr)
		data++
		core.setNZ(data)
		core.bus.Write(addr, data)

	// SED
	case 0xF8:
		core.setFlag(Decimal, true)

	// INC, abs X
	case 0xFE:
		addr := core.getAbsoluteAddr(core.idx)
		data := core.bus.Read(addr)
		data++
		core.setNZ(data)
		core.bus.Write(addr, data)

	}
}
