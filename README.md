# md5Magic

## How to get

`go get github.com/hcolde/md5Magic`

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
		fmt.Println(err.Error()) // Characters must be between 0 and F (Case-insensitive)
		return
	}
	fmt.Println(a, b) // 12557053830604388346 8956594919762969239

	c := md5Magic.Restore(a, b)
	fmt.Println(c) // AE439DEDC6A1F7FA7C4C37A01D8A4297
}

```

## Note

1. The characters are Case-insensitive,but the restored characters should be all uppercase;
2. You'll get the 2 numbers with type of unsigned bigint from Magic function;
3. Please follow the LGPL License!
