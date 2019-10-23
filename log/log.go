/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-08 23:42:45
 * @LastEditTime: 2019-10-09 23:16:32
 * @LastEditors: Do not edit
 */
package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

func InitLogConf() {
	y, m, d := time.Now().Date()
	if _, err := os.Stat("logs/"); os.IsNotExist(err) {
		os.Mkdir("logs/", os.ModePerm)
	}
	fileName := fmt.Sprintf("logs/%d-%d-%d_cnbeta-spider.log", y, m, d)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("file open error, %v", err)
	}

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(f)
}
