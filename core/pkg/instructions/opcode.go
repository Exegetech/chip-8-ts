package instructions

func CalculateOpcode(firstByte, secondByte uint8) uint16 {
	firstByteUint16 := uint16(firstByte)
	secondByteUint16 := uint16(secondByte)
	shiftedFirstByte := firstByteUint16 << 8

	opcode := shiftedFirstByte | secondByteUint16

	return opcode
}
