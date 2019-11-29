/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-23 22:23:39
 * @LastEditTime: 2019-11-29 11:30:56
 * @LastEditors: Please set LastEditors
 */
package v1

import (
	mysqlWrapper "cnbeta-go/mysql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		return
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
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&successResponse{StatusCode: http.StatusOK, Message: "ok", Result: article}); err != nil {
		handleErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

func proxyArticleContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	sid := params.ByName("sid")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://www.cnbeta.com/articles/%s.htm", sid), nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HTTP request failed, err:", err)
		handleErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Http status code:", resp.StatusCode)
		handleErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	defer resp.Body.Close()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	data, _ := ioutil.ReadAll(resp.Body)
	w.Write(data)
}

func proxyVideoList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	query := r.URL.Query()

	var url string
	if bid, ok := query["b_id"]; !ok {
		url = "https://36kr.com/pp/api/info-flow/channel_video/videos?per_page=20"
	} else {
		lastBid, _ := strconv.Atoi(bid[0])
		url = fmt.Sprintf("https://36kr.com/pp/api/info-flow/channel_video/videos?b_id=%d&per_page=20", lastBid)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HTTP request failed, err:", err)
		handleErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Http status code:", resp.StatusCode)
		handleErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, _ := ioutil.ReadAll(resp.Body)
	w.Write(data)
}

func handleErrorResponse(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&failureResponse{StatusCode: status, Err: errMsg})
}
