package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/MDGSF/utils/log"
)

type TStudent struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

type TConfig struct {
	configName string

	ID      string   `json:"ID"`
	Student TStudent `json:"Student"`
	Addr    string   `json:"address"`
}

func NewConfig(configName string) *TConfig {
	data, err := ioutil.ReadFile(configName)
	if err != nil {
		return nil
	}

	c := &TConfig{}
	c.configName = configName
	if err := json.Unmarshal(data, c); err != nil {
		return nil
	}

	return c
}

func (c *TConfig) Save() {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return
	}

	if err := ioutil.WriteFile(c.configName, data, 0666); err != nil {
		return
	}
}

func main() {
	c := NewConfig("config.json")
	log.Info("%v", c)

	c.ID = "I'm new ID"
	log.Info("%v", c)

	c.Save()
}
