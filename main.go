package main

import (
	"articles/article"
	"articles/model"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	db.DB()

	router := gin.Default()

	router.GET("/articles/:limit/:offset", article.GetArticles)
	router.GET("/article/:id", article.GetArticle)
	router.POST("/article", article.CreateArticle)
	router.PUT("/article/:id", article.UpdateArticle)
	router.DELETE("/article/:id", article.DeleteCategory)

	log.Fatal(router.Run(":8000"))
}
