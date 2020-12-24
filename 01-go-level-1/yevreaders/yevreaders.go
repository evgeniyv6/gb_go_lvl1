package yevreaders

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type yamlConf struct {
	Port  int8   `yaml:"port"`
	DbURL string `db_url:"db_url"`
}

func (y *yamlConf) YevJSONParser(filename string) *yamlConf {
	ymlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Read file error %v", err)
	}

	err = yaml.Unmarshal(ymlFile, y)
	if err != nil {
		log.Fatalf("Yaml parse error %v", err)
	}

	return y
}
