package mysql

import (
	configEngine "cnbeta-go/config"
	"cnbeta-go/model"
	"fmt"
	"log"

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

func InitDB() {
	dbConfig := dbConf{}
	configEngine.Engine.GetStruct("MySQL", &dbConfig)

	var err error
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.UserName, dbConfig.Password, dbConfig.LocalUrl, dbConfig.LocalPort, dbConfig.DbName)
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

func SaveArticles(articles []model.ArticleStruct) error {
	for _, article := range articles {
		if err := db.Create(&article).Error; err != nil {
			log.Println("Inserting article failed with error:", err)
			return err
		}
	}
	return nil
}

func UpdateArticle(article model.ArticleStruct) error {
	if err := db.Model(&article).Update(article).Error; err != nil {
		log.Println("Updating article failed with error:", err)
		return err
	}
	return nil
}

func QueryArticle(article *model.ArticleStruct) (model.ArticleStruct, error) {
	if err := db.Model(article).Where("sid = ？", article.Sid).First(article).Error; err != nil {
		return model.ArticleStruct{}, err
	}
	return *article, nil
}
