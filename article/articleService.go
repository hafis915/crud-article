package article

import (
	"articles/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NewArticle struct {
	Title    string `json:"title" binding:"required,min=20"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3"`
	Status   string `json:"status"`
}

type ArticleUpdate struct {
	Title    string `json:"title" binding:"required,min=20"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3"`
	Status   string `json:"status"`
}

func CreateArticle(c *gin.Context) {
	var article NewArticle

	err := c.ShouldBindJSON(&article)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	newArticle := model.Article{
		Title:    article.Title,
		Content:  article.Content,
		Category: article.Category,
		Status:   article.Status,
	}

	db, err := model.Database()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = db.Create(&newArticle).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newArticle)
}

func GetArticle(c *gin.Context) {
	var article model.Article
	db, err := model.Database()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = db.Where("id=?", c.Param("id")).First(&article).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, article)
}

func GetArticles(c *gin.Context) {
	var articles []model.Article

	db, err := model.Database()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	limit, err := strconv.Atoi(c.Param("limit"))
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&articles).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)
}

func UpdateArticle(c *gin.Context) {
	var article model.Article

	db, err := model.Database()

	err = db.Where("id= ?", c.Param("id")).First(&article).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	var updateArticle ArticleUpdate
	err = c.ShouldBindJSON(&updateArticle)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	updatedArticle := model.Article{
		Title:    updateArticle.Title,
		Content:  updateArticle.Content,
		Category: updateArticle.Category,
		Status:   updateArticle.Status,
	}
	err = db.Model(&article).Updates(updatedArticle).Error
	c.JSON(200, article)
}

func DeleteCategory(c *gin.Context) {
	var article model.Article

	db, err := model.Database()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Where("id= ?", c.Param("id")).First(&article).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Data Not Found"})
		return
	}

	err = db.Delete(&article).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Article deleted"})

}
