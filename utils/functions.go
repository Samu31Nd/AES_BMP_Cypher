package utils

func IndexOf(target string, list []string) int {
	for i, v := range list {
		if v == target {
			return i
		}
	}
	return -1 // si no se encuentra
}
