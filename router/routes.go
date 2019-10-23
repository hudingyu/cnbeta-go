package router

type Route struct {
	Name   string
	Method string
	Path   string
}

func generateRoutes() []Route {
	routes := []Route{
		Route{"FetchArticleList", "GET", "/timeline"},
		Route{"FetchArticle", "GET", "/article/:sid"},
	}
	return routes
}
