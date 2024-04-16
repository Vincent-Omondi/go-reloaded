package transformWords

func ToLower(s string) string {
	var result string
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			result += string(v + ('a' - 'A'))
		} else if v >= 'a' && v <= 'z' {
			result += string(v)
		} else {
			result += string(v)
		}
	}
	return result
}
