package iseven

// IsEven returns true if the number is even, false otherwise
func IsEven(n int) bool {
	return n&1 == 0
}
