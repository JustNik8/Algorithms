package main

func reverseBits(num uint32) uint32 {
	ans := uint32(0)
	mask := uint32(1 << 31)
	ansMask := uint32(1)

	for i := 0; i < 32; i++ {
		bit := num & mask

		if bit > 0 {
			ans = ans | ansMask
		}

		mask = mask >> 1
		ansMask = ansMask << 1
	}

	return ans
}
