package md5Magic

import (
	"fmt"
	"errors"
)

func ascii2int(ut uint8) (u uint64) {
	switch ut >> 4 {
	case 3:
		u = uint64(ut - uint8(48))
	case 4:
		u = uint64(ut - uint8(55))
	case 5:
		u = uint64(ut - uint8(87))
	default:
		u = u
	}
	return
}

func check(s string) bool {
	for _, i := range s {
		if i < '0'|| (i > '9' && i < 'A') || (i > 'F' && i < 'a') || i > 'f' {
			return false
		}
	}
	return true
}

func Magic(s string) (r1 uint64, r2 uint64, err error) {
	if check(s) == false {
		err = errors.New("Characters must be between 0 and F (Case-insensitive)")
		return
	}

	for i := 0; i < 16; i += 2 {
		u1 := ascii2int(s[i])
		u2 := ascii2int(s[i+1])
		r1 = (r1 << 8) + (u1 << 4) + u2
	}

	for j := 16; j < 32; j += 2 {
		u1 := ascii2int(s[j])
		u2 := ascii2int(s[j+1])
		r2 = (r2 << 8) + (u1 << 4) + u2
	}

	return
}

func int2ascii(n uint64) (s string) {
	if n <= 9 {
		s = string(n + 48)
	} else {
		s = string(n + 55)
	}
	return
}

func Restore(n1 uint64, n2 uint64) (s string) {
	for n2 > 0 {
		s = int2ascii(n2 & 0x0f) + s
		n2 = n2 >> 4
	}

	for n1 > 0 {
		s = int2ascii(n1 & 0x0f) + s
		n1 = n1 >> 4
	}
	return
}
