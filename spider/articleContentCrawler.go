package spider

import (
	"cnbeta-go/model"
	mysqlWrapper "cnbeta-go/mysql"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type cString string

func ContentCrawlerRun() {
	fmt.Println("grabbing article content starts...")
	start := time.Now()

	articleList, err := mysqlWrapper.QueryUncachedArticles()
	if err != nil {
		log.Println("Fail to query uncached articles, err:", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(articleList))

	for _, article := range articleList {
		go func(article model.ArticleStruct) {
			defer wg.Done()

			// defer func() {
			// 	// 发生宕机时，获取panic传递的上下文并打印
			// 	err := recover()
			// 	switch err.(type) {
			// 	case runtime.Error: // 运行时错误
			// 		fmt.Println("runtime error:", err)
			// 	default: // 非运行时错误
			// 		fmt.Println("error:", err)
			// 	}
			// }()

			crawlArticleContent(article)
		}(article)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Article content crawlling took", elapsed)
}

func crawlArticleContent(article model.ArticleStruct) {
	client := &http.Client{}
	pageUrl := fmt.Sprintf("%s%s.htm", siteConfig.ContentBaseURL, article.Sid)
	req, _ := http.NewRequest("GET", pageUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("HTTP request failed, err:", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("Http status code:", resp.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Fail to parse html, err:", err)
		return
	}

	styleReg := regexp.MustCompile(` style="[^"]*"`)
	scriptReg := regexp.MustCompile(`<script(.*?)</script>`)
	if source := doc.Find(".article-byline span a").First().Text(); len(source) > 0 {
		article.Source = source
	} else {
		article.Source = doc.Find(".article-byline span").First().Text()
	}
	article.Summary = doc.Find(".article-summ p").First().Text()
	content, _ := doc.Find(".articleCont").First().Html()
	article.Content = cString(content).replace(styleReg, "").replace(scriptReg, "").toString()

	if err := mysqlWrapper.UpdateArticle(article); err != nil {
		fmt.Println("Fail to update article, err:", err)
	} else {
		fmt.Println("Succeed to update article!")
	}
}

func (str cString) replace(reg *regexp.Regexp, replaceValue string) cString {
	return cString(reg.ReplaceAllString(string(str), replaceValue))
}

func (str cString) toString() string {
	return string(str)
}
