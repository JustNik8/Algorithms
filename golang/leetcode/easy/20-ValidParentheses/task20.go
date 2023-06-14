package main

func main() {

}

func isValid(s string) bool {
	matching := map[uint8]uint8{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := make([]uint8, 0)

	for i := 0; i < len(s); i++ {
		b := s[i]

		if b == '(' || b == '{' || b == '[' {
			stack = append(stack, b)
		} else {
			sizeStack := len(stack)
			if sizeStack == 0 {
				return false
			}
			open := stack[sizeStack-1]
			closed := matching[open]

			if b != closed {
				return false
			}

			stack = stack[:sizeStack-1]
		}
	}

	return len(stack) == 0
}
