package md5Magic

import (
	"errors"
	"strings"
)

func Magic(s string) (r1 uint64, r2 uint64, err error) {
	if check(s) == false {
		err = errors.New("Parameter must be between 0 and F (Case-insensitive),and length must be 32")
		return
	}

	for i := 0; i < 16; i += 2 {
		u1 := ascii2Int(s[i])
		u2 := ascii2Int(s[i+1])
		r1 = (r1 << 8) + (u1 << 4) + u2
	}

	for j := 16; j < 32; j += 2 {
		u1 := ascii2Int(s[j])
		u2 := ascii2Int(s[j+1])
		r2 = (r2 << 8) + (u1 << 4) + u2
	}

	return
}

func check(s string) bool {
	if len(s) != 32 {
		return false
	}
	for _, i := range s {
		if i < '0'|| (i > '9' && i < 'A') || (i > 'F' && i < 'a') || i > 'f' {
			return false
		}
	}
	return true
}

func ascii2Int(ut uint8) (u uint64) {
	switch ut >> 4 {
	case 3:
		u = uint64(ut - '0')
	case 4:
		u = uint64(ut - '7')
	case 6:
		u = uint64(ut - 'W')
	}
	return
}

func int2Ascii(n uint64, caps string) (s string) {
	switch n {
	case 0:
		s = "0"
	case 1:
		s = "1"
	case 2:
		s = "2"
	case 3:
		s = "3"
	case 4:
		s = "4"
	case 5:
		s = "5"
	case 6:
		s = "6"
	case 7:
		s = "7"
	case 8:
		s = "8"
	case 9:
		s = "9"
	case 10:
		if caps == "UPPER" {
			s = "A"
		} else {
			s = "a"
		}
	case 11:
		if caps == "UPPER" {
			s = "B"
		} else {
			s = "b"
		}
	case 12:
		if caps == "UPPER" {
			s = "C"
		} else {
			s = "c"
		}
	case 13:
		if caps == "UPPER" {
			s = "D"
		} else {
			s = "d"
		}
	case 14:
		if caps == "UPPER" {
			s = "E"
		} else {
			s = "e"
		}
	case 15:
		if caps == "UPPER" {
			s = "F"
		} else {
			s = "f"
		}
	}
	return
}

func Restore(n1 uint64, n2 uint64, args ... string) (s string) {
	caps := "UPPER"
	if len(args) > 0 {
		if (args[0] == "LOWER") {
			caps = "LOWER"
		}else {
			caps = "UPPER"
		}
	}

	s2 := ""
	for n2 > 0 {
		s2 = strings.Join([]string{int2Ascii(n2 & 0xf , caps), s2}, "")
		n2 = n2 >> 4
	}

	for len(s2) < 16 {
		s2 = strings.Join([]string{"0", s2}, "")
	}

	s1 := ""
	for n1 > 0 {
		s1 = strings.Join([]string{int2Ascii(n1 & 0xf, caps), s1}, "")
		n1 = n1 >> 4
	}

	for len(s1) < 16 {
		s1 = strings.Join([]string{"0", s1}, "")
	}

	s = s1 + s2
	return
}
