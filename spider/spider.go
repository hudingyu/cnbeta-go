/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-09-26 23:19:51
 * @LastEditTime: 2019-11-24 13:41:39
 * @LastEditors: Please set LastEditors
 */
package spider

import (
	"fmt"
	"runtime"

	configEngine "cnbeta-go/config"
)

type ServerConf struct {
	ServerIp   string
	ServerPort int
}
type SiteConf struct {
	Website        string
	ListBaseURL    string
	ContentBaseURL string
	TotalPage      int
}

var siteConfig SiteConf

func SpiderRun() {
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()

	fmt.Println("start...")
	siteConfig = SiteConf{}
	configEngine.Engine.GetStruct("Site", &siteConfig)

	ArticleListCrawlerRun()
	ContentCrawlerRun()
}

// func (c *ConfigEngine) getConf() *ConfigEngine {
// 	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println("yaml read")
// 	err = yaml.Unmarshal(yamlFile, &c.data)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	return c
// }
