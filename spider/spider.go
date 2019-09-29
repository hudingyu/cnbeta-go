/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-09-26 23:19:51
 * @LastEditTime: 2019-09-27 13:38:24
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ConfigEngine struct {
	data map[interface{}]interface{}
}
type conf struct {
	website        string `yaml:"website"`
	listBaseUrl    string `yaml:"listBaseUrl"`
	contentBaseUrl string	`yaml:"contentBaseUrl"`
}

func main() {
	fmt.Println("start...")
	c := ConfigEngine{}
	conf := c.getConf()
	fmt.Println(conf)
}

func (c *ConfigEngine) getConf() *ConfigEngine {
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("yaml read")
	err = yaml.Unmarshal(yamlFile, &c.data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
