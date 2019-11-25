/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-23 22:23:39
 * @LastEditTime: 2019-11-25 14:35:52
 * @LastEditors: Please set LastEditors
 */
package v1

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name       string
	Method     string
	Path       string
	HandleFunc httprouter.Handle
}

func GenerateRoutes() []Route {
	routes := []Route{
		Route{"FetchArticleList", "GET", "/api/timeline", FetchArticleList},
		Route{"FetchArticle", "GET", "/api/article/:sid", FetchArticle},
		Route{"ProxyArticleContent", "GET", "/api/news/:sid", proxyArticleContent},
		Route{"ProxyVideoList", "GET", "/api/videolist", proxyVideoList},
	}
	return routes
}
