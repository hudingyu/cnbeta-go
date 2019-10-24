/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-23 22:23:39
 * @LastEditTime: 2019-10-24 12:15:20
 * @LastEditors: Please set LastEditors
 */
package handler

import (
	mysqlWrapper "cnbeta-go/mysql"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func FetchArticleList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pageSize, _ := strconv.Atoi(params.ByName("pageSize"))
	lastSid, _ := strconv.Atoi(params.ByName("lastSid"))
	articleList, err := mysqlWrapper.QueryArticleList(pageSize, lastSid)
	if err != nil {

	}

}
