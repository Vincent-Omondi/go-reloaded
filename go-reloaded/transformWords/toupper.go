package transformWords

func ToUpper(s string) string {
	var result string
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			result += string(v - ('a' - 'A'))
		} else if v >= 'A' && v <= 'Z' {
			result += string(v)
		} else {
			result += string(v)
		}
	}
	return result
}
