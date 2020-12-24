package main

import (
	"fmt"
	"github.com/evgeniyv6/gb_go_lvl1/01-go-level-1/yevreaders"
)

func main() {
	var c yevreaders.yamlConf
	c.YevJSONParser()
	fmt.Println(c)
}
