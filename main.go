/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-09-26 22:47:26
 * @LastEditTime: 2019-10-06 16:19:49
 * @LastEditors: Please set LastEditors
 */
package main

import (
	confEngine "cnbeta-go/config"
	logger "cnbeta-go/log"
	mysqlWrapper "cnbeta-go/mysql"
	"cnbeta-go/spider"
	"time"
)

func main() {
	logger.InitLogConf()
	confEngine.InitConfEngine()
	mysqlWrapper.InitDB()

	ticker := time.NewTicker(10 * time.Second)
	for _ = range ticker.C {
		spider.SpiderRun()
	}

	mysqlWrapper.CloseDB()
}
