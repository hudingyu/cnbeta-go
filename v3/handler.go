/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-23 22:23:39
 * @LastEditTime: 2019-11-29 11:38:34
 * @LastEditors: Please set LastEditors
 */
package v3

import (
	mysqlWrapper "cnbeta-go/mysql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchArticleList(c *gin.Context) {
	lastSid, _ := strconv.Atoi(c.DefaultQuery("last_sid", "0"))
	articleList, err := mysqlWrapper.QueryArticleList(35, lastSid)
	if err != nil {
		c.String(http.StatusNotFound, "Record Not Found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "ok",
		"result":      articleList,
	})
}

func FetchArticle(c *gin.Context) {
	sid := c.Param("sid")
	article, err := mysqlWrapper.QueryArticle(sid)
	if err != nil {
		c.String(http.StatusNotFound, "Record Not Found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "ok",
		"result":      article,
	})
}

func proxyArticleContent(c *gin.Context) {
	sid := c.Param("sid")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://www.cnbeta.com/articles/%s.htm", sid), nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HTTP request failed, err:", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Http status code:", resp.StatusCode)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, string(data))
}

func proxyVideoList(c *gin.Context) {
	bid := c.Query("b_id")
	var url string
	if bid == "" {
		url = "https://36kr.com/pp/api/info-flow/channel_video/videos?per_page=20"
	} else {
		lastBid, _ := strconv.Atoi(bid)
		url = fmt.Sprintf("https://36kr.com/pp/api/info-flow/channel_video/videos?b_id=%d&per_page=20", lastBid)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HTTP request failed, err:", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Http status code:", resp.StatusCode)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	defer resp.Body.Close()

	c.Header("Content-Type", "application/json; charset=UTF-8")
	data, _ := ioutil.ReadAll(resp.Body)
	c.String(http.StatusOK, string(data))
}
