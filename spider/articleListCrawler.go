package spider

import (
	"cnbeta-go/model"
	mysqlWrapper "cnbeta-go/mysql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// {"sid":"895663","catid":"7","topic":"717","aid":"ugmbbc","user_id":"0","title":"Q3完成9.7万辆交付 特斯拉离完成全年目标还有多远","keywords":"","comments":"0","counter":"19","mview":"19","collectnum":"0","good":"0","bad":"0","score":"0","ratings":"0","score_story":"0","ratings_story":"0","pollid":"0","queueid":"0","inputtime":"2019-10-3 17:34","updatetime":"0","thumb":"https://static.cnbetacdn.com/thumb/mini/article/2019/1003/3836c507fd26e52.jpg","source":"新京报","sourceid":"1005","premium":"0","url_show":"http://www.cnbeta.com/articles/tech/895663.htm","label":"科技","url":"/touch/articles/895663.htm"
type articleListRespStruct struct {
	state   string
	message string
	Result  struct {
		List []model.ArticleStruct `json:"list"`
	} `json:"result"`
}

/*
 * @description: 初始化方法 抓取文章列表
 * @param {type}
 * @return:
 */
func ArticleListCrawlerRun() {
	fmt.Println("grabbing article list starts...")
	start := time.Now()
	pageUrlList := getPageUrlList(siteConfig.TotalPage)

	var wg sync.WaitGroup
	wg.Add(len(pageUrlList))

	for _, pageUrl := range pageUrlList {
		go func(url string) {
			defer wg.Done()

			defer func() {
				// 发生宕机时，获取panic传递的上下文并打印
				err := recover()
				switch err.(type) {
				case runtime.Error: // 运行时错误
					fmt.Println("runtime error:", err)
				default: // 非运行时错误
					fmt.Println("error:", err)
				}
			}()

			getArticleList(url)
		}(pageUrl)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("ArticleList crawlling took %s\n", elapsed)
}

func getPageUrlList(totalPage int) []string {
	var pageUrlList []string = make([]string, 0)
	for page := 1; page <= totalPage; page++ {
		pageUrlList = append(pageUrlList, fmt.Sprintf("%s%d", siteConfig.ListBaseURL, page))
	}
	return pageUrlList
}

func getArticleList(pageUrl string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", pageUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Http get err:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Http status code:", resp.StatusCode)
		return
	}

	// defer mysqlWrapper.CloseDB()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("read from resp.Body failed,err:", err)
	// 	return
	// }
	// fmt.Printf("%s\n", body)

	data := articleListRespStruct{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("JSON parsing err:", err)
		return
	}
	articleList := data.Result.List
	// fmt.Println(articleList[0])
	if err := mysqlWrapper.CacheArticles(articleList); err != nil {
		fmt.Println("Saving to mysql failed, err:", err)
	} else {
		fmt.Println("Saving to mysql succeeded!")
	}
}
