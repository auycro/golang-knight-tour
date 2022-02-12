package utility

func GetFirstKeyByValue(m map[string]int, value int) string {
	for k, v := range m {
		if value == v {
			return k
		}
	}
	return ""
}
