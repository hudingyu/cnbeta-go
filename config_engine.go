/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-09-29 22:57:59
 * @LastEditTime: 2019-09-29 23:45:10
 * @LastEditors: Do not edit
 */
package config_engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

type ConfigEngine struct {
	data map[interface{}]interface{}
}

func (c *ConfigEngine) LoadYaml(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("can't read file:" + path)
	}
	err = yaml.Unmarshal(yamlFile, &c.data)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("can not parse " + path + " config")
	}
	return nil
}

func (c *ConfigEngine) Get(name string) interface{} {
	path := strings.Split(name, ".")
	data := c.data
	for index, value := range path {
		v, ok := data[value]
		if !ok {
			break
		}
		if index+1 == len(path) {
			return v
		}
		if reflect.TypeOf(v).String() == "map[interface {}] interface {}" {
			data = v.(map[interface{}]interface{})
		}
	}
	return nil
}

func (c *ConfigEngine) GetStruct(name string, desObj interface{}) interface{} {
	data := c.Get(name)
	switch data.(type) {
	case string:
		c.setField(desObj, name, data)
	case map[interface{}]interface{}:
		c.mapToStruct(data.(map[interface{}]interface{}), desObj)
	}
	return desObj
}

func (c *ConfigEngine) mapToStruct(sourceMap map[interface{}]interface{}, des interface{}) interface{} {
	for key, value := range sourceMap {
		c.setField(des, key.(string), value)
	}
	return des
}

func (c *ConfigEngine) setField(desStruct interface{}, name string, newValue interface{}) error {
	desStructValue := reflect.Indirect(reflect.ValueOf(desStruct))
	curFieldValue := desStructValue.FieldByName(name)

	curFieldType := curFieldValue.Type()
	nextValue := reflect.ValueOf(newValue)

	if curFieldType.Kind() == reflect.Struct && nextValue.Kind() == reflect.Map {
		for key, value := range newValue.(map[interface{}]interface{}) {
			c.setField(curFieldValue.Addr().Interface(), key.(string), value)
		}
	} else {
		if curFieldType != nextValue.Type() {
			return errors.New("Provided value type didn't match desination struct field type")
		}

		curFieldValue.Set(nextValue)
	}

	return nil
}
