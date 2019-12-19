/*
 * @Description:
 * @Author: hudingyu
 * @Date: 2019-10-08 23:42:45
 * @LastEditTime : 2019-12-19 20:29:06
 * @LastEditors  : Please set LastEditors
 */
package mysql

import (
	configEngine "cnbeta-go/config"
	"cnbeta-go/model"
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dbConf struct {
	UserName  string
	Password  string
	LocalUrl  string
	LocalPort int
	OriginIp  string
	DbName    string
}

var db *gorm.DB

func init() {

}
func InitDB() {
	dbConfig := dbConf{}
	configEngine.Engine.GetStruct("MySQL", &dbConfig)

	var err error
	// dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.UserName, dbConfig.Password, dbConfig.LocalUrl, dbConfig.LocalPort, dbConfig.DbName)
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.UserName, dbConfig.Password, dbConfig.OriginIp, dbConfig.LocalPort, dbConfig.DbName)
	db, err = gorm.Open("mysql", dbUrl)
	// defer db.Close()
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
	db.AutoMigrate(&model.ArticleStruct{})
}

func CloseDB() {
	db.Close()
}

func CacheArticles(articles []model.ArticleStruct) error {
	for _, article := range articles {
		var count int
		db.Model(&model.ArticleStruct{}).Where("sid = ?", article.Sid).Count(&count)
		if count == 0 {
			if err := db.Create(&article).Error; err != nil {
				log.Println("Inserting article failed with error:", err)
				return err
			}
		} else {
			// if err := db.Model(&article).Update(article).Error; err != nil {
			// 	log.Println("Updating article failed with error:", err)
			// 	return err
			// }
			// 执行原生SQL, 只更新部分需要更新的字段
			sql, arguments := generateSQLForUpdatingArticle(article)
			if err := db.Exec(sql, arguments...).Error; err != nil {
				log.Println("Updating article failed with error:", err)
				return err
			}
		}
	}
	return nil
}

func UpdateArticle(article model.ArticleStruct) error {
	sql, arguments := generateSQLForUpdatingArticle(article)
	if err := db.Exec(sql, arguments...).Error; err != nil {
		log.Println("Updating article failed with error:", err)
		return err
	}
	return nil
}

func QueryArticle(sid string) (model.ArticleStruct, error) {
	article := model.ArticleStruct{}
	if err := db.Model(article).Where("sid = ?", sid).First(&article).Error; err != nil {
		return model.ArticleStruct{}, err
	}
	return article, nil
}

func QueryUncachedArticles() ([]model.ArticleStruct, error) {
	articleList := []model.ArticleStruct{}
	err := db.Model(&model.ArticleStruct{}).Where("content = ?", "").Find(&articleList).Error
	return articleList, err
}

func QueryArticleList(limit int, lastSid int) ([]model.ArticleStruct, error) {
	articleList := []model.ArticleStruct{}
	var err error
	if lastSid != 0 {
		err = db.Model(&model.ArticleStruct{}).Select("sid, title, pub_time, summary, thumb, comment_count").Where("sid < ?", lastSid).Order("sid desc").Limit(limit).Find(&articleList).Error
	} else {
		err = db.Model(&model.ArticleStruct{}).Select("sid, title, pub_time, summary, thumb, comment_count").Order("sid desc").Limit(limit).Find(&articleList).Error
	}
	return articleList, err
}

func generateSQLForUpdatingArticle(article model.ArticleStruct) (string, []interface{}) {
	var columns = make([]string, 0)
	var arguments = make([]interface{}, 0)

	if len(article.CommentCount) > 0 {
		columns = append(columns, "comment_count = ?")
		arguments = append(arguments, article.CommentCount)
	}

	if len(article.Source) > 0 {
		columns = append(columns, "source = ?")
		arguments = append(arguments, article.Source)
	}

	if len(article.Summary) > 0 {
		columns = append(columns, "summary = ?")
		arguments = append(arguments, article.Summary)
	}

	if len(article.Content) > 0 {
		columns = append(columns, "content = ?")
		arguments = append(arguments, article.Content)
	}

	sql := fmt.Sprintf("UPDATE article_structs SET %s WHERE sid = %s", strings.Join(columns, ","), article.Sid)
	return sql, arguments
}
