/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-23 22:23:39
 * @LastEditTime: 2019-10-23 23:43:47
 * @LastEditors: Do not edit
 */
package handler

import (
	mysqlWrapper "cnbeta-go/mysql"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func FetchArticleList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	articleList := mysqlWrapper.QueryArticleList(params.ByName("page"))
}
