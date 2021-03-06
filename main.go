/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-09-26 22:47:26
 * @LastEditTime: 2019-12-02 15:58:34
 * @LastEditors: Please set LastEditors
 */
package main

import (
	confEngine "cnbeta-go/config"
	logger "cnbeta-go/log"
	mysqlWrapper "cnbeta-go/mysql"
	"cnbeta-go/spider"
	v3 "cnbeta-go/v3"
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
		for range ticker.C {
			spider.SpiderRun()
		}
	}()

	// router := v1.NewRouter(v1.GenerateRoutes())
	router := v3.NewRouter(v3.GenerateRoutes())
	http.ListenAndServe(":80", router)
}
