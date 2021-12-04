package tape

func Max(a int64, b int64) int64 {
	if a < b {
		return b
	} else {
		return a
	}
}

func Abs(n int64) int64 {
	if n < 0 {
		return -1 * n
	}

	return n
}
