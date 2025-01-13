package mock

const (
	// RAX, RBX, RCX, RDI, RSI, R8, R9, R10, R11.
	intArgRegs = 9

	// X0 -> X14.
	floatArgRegs = 15

	// We use SSE2 registers which support 64-bit float operations.
	floatRegSize = 8

	bigEndian = false
)
