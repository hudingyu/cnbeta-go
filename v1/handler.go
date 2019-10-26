/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-23 22:23:39
 * @LastEditTime: 2019-10-26 19:04:48
 * @LastEditors: Please set LastEditors
 */
package v1

import (
	mysqlWrapper "cnbeta-go/mysql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type successResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}

type failureResponse struct {
	StatusCode int    `json:"status_code"`
	Err        string `json:"err"`
}

func FetchArticleList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// pageSize, _ := strconv.Atoi(params.ByName("page"))
	// lastSid, _ := strconv.Atoi(params.ByName("last_sid"))
	lastSid, _ := strconv.Atoi(r.URL.Query()["last_sid"][0])
	articleList, err := mysqlWrapper.QueryArticleList(35, lastSid)
	if err != nil {
		handleErrorResponse(w, http.StatusNotFound, "Record Not Found")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := successResponse{
		StatusCode: http.StatusOK,
		Message:    "ok",
		Result:     articleList,
	}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		handleErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

func FetchArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	sid := params.ByName("sid")
	article, err := mysqlWrapper.QueryArticle(sid)
	if err != nil {
		handleErrorResponse(w, http.StatusNotFound, "Record Not Found")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&successResponse{StatusCode: http.StatusOK, Message: "ok", Result: article}); err != nil {
		handleErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

func handleErrorResponse(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&failureResponse{StatusCode: http.StatusInternalServerError, Err: errMsg})
}
