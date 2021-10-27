package golang

import "strings"

func isNumber(s string) bool {
	if len(s) <= 0 {
		return false
	}
	var ok bool
	str, numberic := scanInteger(strings.TrimSpace(s))
	if len(str) >0 && str[0] == '.' {
		str = str[1:]
		str, ok = scanUnsignInteger(str)
		numberic = numberic || ok
	}
	if len(str) >0 && (str[0] == 'e' || str[0] == 'E') {
		str = str[1:]
		str, ok = scanInteger(str)
		numberic = numberic && ok
	}
	return numberic && str == ""
}

func scanUnsignInteger(s string) (string, bool) {
	i := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		i++
	}
	str := s[i:]
	return str, i > 0
}

func scanInteger(s string) (string, bool) {
	 str := s
	if str[0] == '+' || str[0] == '-' {
		str = s[1:]
	}
	return scanUnsignInteger(str)
}
