/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-23 22:23:39
 * @LastEditTime: 2019-10-23 22:42:04
 * @LastEditors: Do not edit
 */
package router

import "github.com/julienschmidt/httprouter"

type Route struct {
	Name       string
	Method     string
	Path       string
	HandleFunc httprouter.Handle
}

func generateRoutes() []Route {
	routes := []Route{
		Route{"FetchArticleList", "GET", "/timeline"},
		Route{"FetchArticle", "GET", "/article/:sid"},
	}
	return routes
}
