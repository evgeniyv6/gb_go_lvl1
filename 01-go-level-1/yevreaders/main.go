package main

import (
	"flag"
	"fmt"
	"github.com/evgeniyv6/gb_go_lvl1/01-go-level-1/yevreaders/yevpckg"
	"log"
)

var confPath = flag.String("yaml-path", "", "path to the yaml config file")

func main() {
	flag.Parse()
	var c yevpckg.YamlConf
	if *confPath == "" {
		log.Fatal("Path cannot be empty")
	} else {
		c.YevYAMLParser(*confPath)
		fmt.Printf("Parse result:\n%+v", c)
	}
}
