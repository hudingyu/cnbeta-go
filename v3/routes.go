/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-23 22:23:39
 * @LastEditTime: 2019-11-27 20:58:12
 * @LastEditors: Please set LastEditors
 */
package v3

import "github.com/gin-gonic/gin"

type Route struct {
	Name       string
	Path       string
	HandleFunc gin.HandlerFunc
}

func GenerateRoutes() []Route {
	routes := []Route{
		Route{"FetchArticleList", "/api/timeline", FetchArticleList},
		Route{"FetchArticle", "/api/article/:sid", FetchArticle},
		Route{"ProxyArticleContent", "/api/news/:sid", proxyArticleContent},
		Route{"ProxyVideoList", "/api/videolist", proxyVideoList},
	}
	return routes
}
