/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-09-26 22:47:26
 * @LastEditTime: 2019-11-10 12:59:37
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
	http.ListenAndServe("127.0.0.0:8080", router)
}
