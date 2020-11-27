package main

import (
	"fmt"
	"strings"
)

type BitFlag int

const (
	Active BitFlag = 1 << iota // 1 << 0 == 1
	Send                       // 1 << 1 == 2
)

var flag = Active | Send

func (flag BitFlag) String() string {
	var flags []string
	if flag&Active == Active {
		flags = append(flags, "Active")
	}
	if flag&Send == Send {
		flags = append(flags, "Send")
	}
	if len(flags) > 0 {
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
	return "0()"
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println(BitFlag(0), Active, Send, flag)
}
