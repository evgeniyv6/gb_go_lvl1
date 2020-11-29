// hello-world print and example from Mark Summerfield book

package main

import (
	"fmt"
	"strings"
)

type bitFlag int

const (
	active bitFlag = 1 << iota // 1 << 0 == 1
	send                       // 1 << 1 == 2
)

var flag = active | send

func (flag bitFlag) String() string {
	var flags []string
	if flag&active == active {
		flags = append(flags, "Active")
	}
	if flag&send == send {
		flags = append(flags, "Send")
	}
	if len(flags) > 0 {
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
	return "0()"
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println(bitFlag(0), active, send, flag)
}
