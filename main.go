/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-09-26 22:47:26
 * @LastEditTime: 2019-10-26 16:48:05
 * @LastEditors: Please set LastEditors
 */
package main

import (
	confEngine "cnbeta-go/config"
	logger "cnbeta-go/log"
	mysqlWrapper "cnbeta-go/mysql"
	"cnbeta-go/spider"
	v1 "cnbeta-go/v1"
	"net/http"
	"time"
)

func main() {
	logger.InitLogConf()
	confEngine.InitConfEngine()
	mysqlWrapper.InitDB()

	go func() {
		spider.SpiderRun()
		ticker := time.NewTicker(10 * 60 * time.Second)
		for _ = range ticker.C {
			spider.SpiderRun()
		}
	}()

	router := v1.NewRouter(v1.GenerateRoutes())
	http.ListenAndServe(":8089", router)

	mysqlWrapper.CloseDB()
}
