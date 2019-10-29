package reverse

func Reverse(s string) string {
	reversed := ""
	if s == "" {
		return ""
	}
	for i := len(s)-1; i >= 0;  i-- {
		reversed += string(s[i])
	}
	return reversed
}