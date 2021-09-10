package organisation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

var organisationArticlesSortCriteria = map[string]string{
	"date":       "created_at",
	"price":      "price",
	"popularity": "date",
	"name":       "name",
}

var ErrTenantNotProvided error = errors.New("you must provide Tenant (organisation id) on Header")

type ArticlesPresenter struct {
	Articles
	Pictures []Pictures `json:"pictures"`
}

//CreateArticleHandler is a handler for CreateArticle
func CreateArticleHandler(ctx *gin.Context) {
	if strings.TrimSpace(ctx.GetHeader("Tenant")) == "" {
		ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
		return
	}

	organisationID, err := strconv.Atoi(ctx.GetHeader("Tenant"))
	if err != nil {
		ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
		return
	}

	article := ArticlesPresenter{}
	err = ctx.BindJSON(&article)
	if err != nil {
		ctx.String(http.StatusNotAcceptable, "You must provide an Articles in json format on the body")
		return
	}

	err = CreateArticle(&article.Articles)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Can't create Article")
		return
	}

	err = CreateOrganisationArticle(uint(organisationID), article.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Can't link Article to the organisation")
		//delete article to avoid zombie articles
		_ = DeleteArticle(article.ID)
		return
	}

	var savedPictures []Pictures
	var picturesIDToSaved []uint
	for i := 0; i < len(article.Pictures); i++ {
		pictureID := article.Pictures[i].ID

		// If the Pictures not exist an not linked to the Organisations we skip the Pictures
		pic, err := GetOrganisationPicture(pictureID, uint(organisationID))
		if err != nil {
			continue
		}

		savedPictures = append(savedPictures, *pic)
		picturesIDToSaved = append(picturesIDToSaved, pic.ID)
	}

	err = CreateArticlePictures(article.ID, picturesIDToSaved)
	if err != nil {
		article.Pictures = []Pictures{}
	}

	article.Pictures = savedPictures

	ctx.JSON(http.StatusOK, article)
}

//GetArticleHandler is a handler for GetArticle
func GetArticleHandler(ctx *gin.Context) {
	if strings.TrimSpace(ctx.GetHeader("Tenant")) == "" {
		ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
		return
	}

	_, err := strconv.Atoi(ctx.GetHeader("Tenant"))
	if err != nil {
		ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
		return
	}

	articleID, err := strconv.Atoi(ctx.Param("articleID"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Article id must be an int")
		return
	}

	article, err := GetArticle(uint(articleID))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	result := ArticlesPresenter{Articles: *article}
	result.Pictures, _ = GetArticlePictures(uint(articleID))

	ctx.JSON(http.StatusOK, result)
}

//GetOrganisationArticlesHandler is a handler for GetOrganisationArticles
func GetOrganisationArticlesHandler(ctx *gin.Context) {
	organisationID, err := strconv.Atoi(ctx.GetHeader("Tenant"))
	if err != nil {
		ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
		return
	}

	// Initializing default
	limit := 1
	offset := 1
	search := ""
	priceGte := 0
	priceLte := 9999999
	categoryId := ""
	sort := ""

	if value, err := strconv.Atoi(ctx.Query("limit")); err == nil && value > limit {
		limit = value
	}

	if value, err := strconv.Atoi(ctx.Query("offset")); err == nil && value > offset {
		offset = value
	}

	search = ctx.Query("search")

	if value, err := strconv.Atoi(ctx.Query("price_gte")); err == nil {
		priceGte = value
	}

	if value, err := strconv.Atoi(ctx.Query("price_lte")); err == nil {
		priceLte = value
	}

	if value, err := strconv.Atoi(ctx.Query("offset")); err == nil {
		offset = value
	}

	categoryId = ctx.Query("category")

	sort, isValidSortCriteria := organisationArticlesSortCriteria[ctx.Query("sort")]
	if !isValidSortCriteria {
		sort = ""
	}

	offset = (offset - 1) * limit
	articles, err := GetOrganisationArticles(uint(organisationID), categoryId, limit, offset, search, priceGte, priceLte, sort)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	results := make([]ArticlesPresenter, len(articles))
	for i := 0; i < len(articles); i++ {
		results[i].Articles = articles[i]
		results[i].Pictures, _ = GetArticlePictures(articles[i].ID)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"limit": limit, "offset": offset, "data": results})
}

//UpdateArticleHandler is a handler for UpdateArticle
func UpdateArticleHandler(ctx *gin.Context) {
	if strings.TrimSpace(ctx.GetHeader("Tenant")) == "" {
		ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
		return
	}

	organisationID, err := strconv.Atoi(ctx.GetHeader("Tenant"))
	if err != nil {
		ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
		return
	}

	articleID, err := strconv.Atoi(ctx.Param("articleID"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Article id must be an int")
		return
	}

	// We retrieve the the saved article
	storedArticle, err := GetArticle(uint(articleID))
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		ctx.String(http.StatusNotFound, "Ressource not found")
		return
	default:
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	}
	storedPictures, _ := GetArticlePictures(uint(articleID))

	var providedArticle ArticlesPresenter
	err = ctx.BindJSON(&providedArticle)
	if err != nil {
		ctx.String(http.StatusNotAcceptable, "You must provide an Articles in json format on the body")
		return
	}

	if strings.TrimSpace(providedArticle.Name) == "" {
		ctx.String(http.StatusBadRequest, "article name can't be empty")
		return
	}

	if providedArticle.Quantity == 0 {
		ctx.String(http.StatusBadRequest, "article quantity can't be 0")
		return
	}

	if providedArticle.Price == 0 {
		ctx.String(http.StatusBadRequest, "article price can't be 0")
		return
	}

	var picturesSaved []Pictures
	var picturesIDToSave []uint
	for i := 0; i < len(storedPictures); i++ {
		pictureID := storedPictures[i].ID

		// If the Pictures not exist an not linked to the Organisations we skip the Pictures
		pic, err := GetOrganisationPicture(pictureID, uint(organisationID))
		if err != nil {
			continue
		}

		picturesSaved = append(picturesSaved, *pic)
		picturesIDToSave = append(picturesIDToSave, pic.ID)
	}

	// if saving the new pictures fails we return the last stored pictures
	err = CreateArticlePictures(storedArticle.ID, picturesIDToSave)
	if err != nil {
		providedArticle.Pictures = storedPictures
	} else {
		providedArticle.Pictures = append(storedPictures, picturesSaved...)
	}

	providedArticle.Articles.Model = storedArticle.Model
	err = UpdateArticle(&providedArticle.Articles)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	}

	ctx.JSON(http.StatusOK, providedArticle)
}

//DeleteArticleHandler is a handler for DeleteArticle
func DeleteArticleHandler(ctx *gin.Context) {
	articleID, err := strconv.Atoi(ctx.Param("articleID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Article id must be an int")
		return
	}

	err = DeleteArticle(uint(articleID))
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		ctx.String(http.StatusNotFound, "Ressource not found")
		return
	default:
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	}

	ctx.String(http.StatusOK, "Article Deleted")
}
