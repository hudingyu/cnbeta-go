/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-23 22:23:39
 * @LastEditTime: 2019-10-25 00:19:01
 * @LastEditors: Do not edit
 */
package handler

import (
	mysqlWrapper "cnbeta-go/mysql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type successResponse struct {
	status  int
	message string
	result  interface{}
}

type failureResponse struct {
	status  int
	message string
	err     interface{}
}

func FetchArticleList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pageSize, _ := strconv.Atoi(params.ByName("pageSize"))
	lastSid, _ := strconv.Atoi(params.ByName("lastSid"))
	articleList, err := mysqlWrapper.QueryArticleList(pageSize, lastSid)
	if err != nil {

	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := successResponse{
		status:  http.StatusOK,
		message: "ok",
		result:  articleList,
	}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&failureResponse{status: http.StatusInternalServerError, message: "failure", err: err})
	}
}
