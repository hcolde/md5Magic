# md5Magic

## How to get

`go get -u github.com/hcolde/md5Magic`

## Demo

```golang

package main

import (
	"fmt"
	"github.com/hcolde/md5Magic"
)

func main() {
	md5 := "AE439DEDC6A1F7FA7C4C37A01D8A4297"
	a, b, err := md5Magic.Magic(md5)
	if err != nil {
		fmt.Println(err.Error()) // Parameter must be between 0 and F (Case-insensitive),and length must be 32
		return
	}
	fmt.Println(a, b) // 12557053830604388346 8956594919762969239

	fmt.Println(md5Magic.Restore(a, b)) // AE439DEDC6A1F7FA7C4C37A01D8A4297
	fmt.Println(md5Magic.Restore(a, b, "LOWER")) // ae439dedc6a1f7fa7c4c37a01d8a4297
}

```

## Note

1. The characters are Case-insensitive,but the restored characters should be all uppercase;
2. You'll get the 2 numbers with type of unsigned bigint from Magic function;
3. Please follow the LGPL License!
