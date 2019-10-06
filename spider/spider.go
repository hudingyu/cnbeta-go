/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-09-26 23:19:51
 * @LastEditTime: 2019-10-05 14:55:56
 * @LastEditors: Please set LastEditors
 */
package spider

import (
	"fmt"

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
