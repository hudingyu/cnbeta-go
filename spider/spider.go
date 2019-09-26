/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-09-26 23:19:51
 * @LastEditTime: 2019-09-27 00:55:41
 * @LastEditors: Do not edit
 */
package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type conf struct {
	website        string
	listBaseUrl    string
	contentBaseUrl string
}

func main() {
	fmt.Println("start...")
	var c conf
	conf := c.getConf()
	fmt.Println(conf.website)
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("../config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("yaml read")
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c)
	return c
}
