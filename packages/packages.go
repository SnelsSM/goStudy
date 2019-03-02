package main

import (
	"fmt"

	"github.com/SnelsSM/goStudy/packages/apackage"
	"github.com/SnelsSM/goStudy/packages/bpackage"
)

func main() {
	fmt.Println(apackage.Hello())
	fmt.Println(bpackage.Hello())
}
