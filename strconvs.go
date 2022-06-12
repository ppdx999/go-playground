package main

import (
	"fmt"
	"strconv"
)

func main_strconvs() {
	reslut1, _ := strconv.Atoi("0000101")
	reslut2, _ := strconv.Atoi("aiueo")
	fmt.Printf("%d, %d\n", reslut1, reslut2)
}
