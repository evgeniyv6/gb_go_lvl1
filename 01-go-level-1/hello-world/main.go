// hello-world print and example from Mark Summerfield book

package main

import (
	"fmt"
	"github.com/evgeniyv6/gb_go_lvl1/01-go-level-1/hello-world/c_algos"
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

	fmt.Println(c_algos.IsAutomorph(90626))


	// try 14 exercise
	fmt.Println("Automorphic number:")
	for i:=1; i< 2890625; i++ {
		if c_algos.IsAutomorph(i) {
			fmt.Printf("%d\t", i)
		}
	}
	fmt.Println()
	fmt.Println()

	// try 13 exercise
	fmt.Println("Random numbers")
	fmt.Printf("Standart rand func: %d\n", c_algos.StandartRand(1, 100))
	fmt.Printf("Lehmer rand func: %d\n",c_algos.KnuthRand(1, 100,4,5))

	fmt.Println()
	// try 12 exercise
	fmt.Printf("Max: %d\n",c_algos.MaxFromInt(1,2,3))
	fmt.Printf("Min: %d\n",c_algos.MinFromInt(1,-122,-3))
	fmt.Println()

	// try 6 exercise
	for i:=1; i<=150;i++{
		age, yearDes := c_algos.YearIdentifier(i)
		fmt.Printf("age - %d  %s # ", age, yearDes)
	}

}


