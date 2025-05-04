package utils

func IndexOf(target string, list []string) int {
	for i, v := range list {
		if v == target {
			return i
		}
	}
	return -1 // si no se encuentra
}

func XorBlocks(a, b []byte) []byte {
	result := make([]byte, len(a))
	for i := range a {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func IncrementCounter(counter []byte) {
	for i := len(counter) - 1; i >= 0; i-- {
		counter[i]++
		if counter[i] != 0 {
			break
		}
	}
}
