package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b: 
  c: 2
  d: [3, 4]
`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func test1() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}

func main() {
	data, _ := ioutil.ReadFile("prometheus.yml")
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
	/*
		map[
		scrape_configs:[
		map[
		job_name:prometheus
		static_configs:[
		map[
		targets:[localhost:9090]
		]
		]
		]
		map[
		job_name:node
		static_configs:[
		map[
		targets:[localhost:9100]
		]
		]
		]
		]

		global:map[
		scrape_interval:15s
		evaluation_interval:15s
		]

		alerting:map[
		alertmanagers:[
		map[
		static_configs:[
		map[targets:<nil>]
		]
		]
		]
		]

		rule_files:<nil>
		]
	*/
}
