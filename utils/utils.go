package utils

func MemUnitConversion(byte_value int) (int, int, int) {
	b_kb := byte_value / 1024
	kb_mb := b_kb / 1024
	mb_gb := kb_mb / 1024

	return b_kb, kb_mb, mb_gb

}
