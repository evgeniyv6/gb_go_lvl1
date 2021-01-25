package yevpckg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type yamlExample struct {
	RetrieveResources bool     `yaml:"retrieve-resources"`
	Requests          []string `yaml:"requests"`
}

type execution struct {
	Concurrency int    `yaml:"concurrency"`
	HoldFor     string `yaml:"hold-for"`
	RampUp      string `yaml:"ramp-up"`
	Scenario    string `yaml:"scenario"`
}

type reporting struct {
	Module string `yaml:"module"`
}

// YamlConf descr
type YamlConf struct {
	Port        int         `yaml:"port"`
	DbURL       string      `yaml:"db_url"`
	JaegerURL   string      `yaml:"jaeger_url"`
	SentryURL   string      `yaml:"sentry_url"`
	KafkaBroker string      `yaml:"kafka_broker"`
	SomeAppID   string      `yaml:"some_app_id"`
	SomeAppKey  string      `yaml:"some_app_key"`
	Execution   []execution `yaml:"execution"`
	Scenarios   struct {
		YamlExample yamlExample `yaml:"yaml_example"`
	} `yaml:"scenarios"`
	Reporting []reporting `yaml:"reporting"`
	Settings  struct {
		CheckInterval   string `yaml:"check-interval"`
		DefaultExecutor string `yaml:"default-executor"`
	} `yaml:"settings"`
}

// YevYAMLParser descr
func (y *YamlConf) YevYAMLParser(filename string) *YamlConf {
	ymlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Read file error %v", err)
	}

	err = yaml.Unmarshal(ymlFile, y)
	if err != nil {
		log.Fatalf("Yaml parse error %v", err)
	}
	urlAr := []string{y.DbURL, y.JaegerURL, y.SentryURL}
	urlAr = append(urlAr, y.Scenarios.YamlExample.Requests...)
	for _, ue := range urlAr {
		if res, err := IsDbURLValid(ue); err != nil {
			log.Printf("%s: db url err %s\n", ue, err)
		} else {
			log.Printf("%s: Successful db url parsing: %s\n", ue, res)
		}
	}
	return y
}
