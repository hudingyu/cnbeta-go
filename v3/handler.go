/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-23 22:23:39
 * @LastEditTime: 2019-12-03 12:16:39
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

	"github.com/PuerkitoBio/goquery"
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
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Fail to parse html, err:", err)
	}
	title, _ := doc.Find(".title h1").Html()
	title = fmt.Sprintf("<header class=\"title\"><h1>%s</h1></header>", title)
	summary, _ := doc.Find(".article-summary p").Html()
	summary = fmt.Sprintf("<div class=\"article-summary\"><p>%s</p></div>", summary)
	content, _ := doc.Find(".article-content").Html()
	content = fmt.Sprintf("<div class=\"article-content\">%s</div>", content)
	author, _ := doc.Find(".cnbeta-article-vote-tags").Html()
	article := fmt.Sprintf("%s%s%s", summary, content, author)
	// data, _ := ioutil.ReadAll(resp.Body)
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, article)
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
