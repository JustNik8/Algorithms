package main

func main() {

}

func hammingWeight(num uint32) int {
	mask := uint32(1)
	ans := 0

	for num > 0 {
		ans += int(num & mask)
		num = num >> mask
	}

	return ans
}
