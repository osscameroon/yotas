package articles

import (
	"errors"
	"fmt"
	"github.com/osscameroon/yotas/app"
	"strings"
	"time"
)

//CreateArticle Add an Articles for an Organisations. The return an error if something when wrong
func CreateArticle(article *app.Articles) error {
	if strings.TrimSpace(article.Name) == "" {
		return errors.New("article name can't be empty")
	}

	if article.Quantity == 0 {
		return errors.New("article quantity can't be 0")
	}

	if article.Price == 0 {
		return errors.New("article price can't be 0")
	}

	article.ID = 0
	article.CreatedAt = time.Now().UTC()
	article.UpdatedAt = article.CreatedAt

	return app.Session.Create(article).Error
}

//GetArticle Retrieve an Articles an return a *Articles. If the article are not found the will return an nil pointer with an error
func GetArticle(articleID uint) (*app.Articles, error) {
	var article app.Articles
	result := app.Session.Where("id = ?", articleID).First(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}

//GetOrganisationArticle Retrieve an Articles an return a *Articles who belong to an Organisations. If the article are not found the will return an nil pointer with an error
func GetOrganisationArticle(articleID uint, organisationID uint) (*app.Articles, error) {
	var article app.Articles
	result := app.Session.Where(
		"id = ? AND id IN (?)",
		articleID,
		app.Session.Model(&app.OrganisationsArticles{}).Where("organisation_id = ?", organisationID).Select("article_id")).
		First(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}

//GetArticles Retrieve a list of Articles and return a []Articles. If the article are not found this will return an empty slice with an error
func GetArticles(articlesID []uint) ([]app.Articles, error) {
	var articles []app.Articles
	err := app.Session.Model(&app.Articles{}).Where("id IN (?)", articlesID).Scan(&articles).Error
	return articles, err
}

//UpdateArticle update an Articles
func UpdateArticle(article *app.Articles) error {
	article.UpdatedAt = time.Now().UTC()
	result := app.Session.Save(article)
	return result.Error
}

//DeleteArticle delete an article with the given articleID
func DeleteArticle(articleID uint) error {
	result := app.Session.Model(&app.Articles{}).Where("id = ?", articleID).Update("deleted_at", time.Now().UTC())
	return result.Error
}

//CreateOrganisationArticle Create a new OrganisationsArticles
func CreateOrganisationArticle(organisationID uint, articleID uint) error {
	return app.Session.Create(&app.OrganisationsArticles{
		OrganisationId: organisationID,
		ArticleId:      articleID,
		Model:          app.Model{CreatedAt: time.Now().UTC()},
	}).Error
}

//GetOrganisationArticles Get a list of Articles related to an Organisations
func GetOrganisationArticles(organisationID uint, categoryId string, limit int, offset int, search string, priceGte int, priceLte int, sort string) ([]app.Articles, error) {
	var results []app.Articles
	search = fmt.Sprintf("%s%s%s", "%", search, "%")
	err := app.Session.Model(&app.Articles{}).
		Joins("JOIN organisations_articles on organisations_articles.article_id = articles.id and organisations_articles.organisation_id = ?", organisationID).
		Where("Lower(name) like Lower(?) or Lower(description) like Lower(?) and price >= ? and price <= ?", search, search, priceGte, priceLte).
		Limit(limit).
		Offset(offset).Scan(&results).Error

	return results, err
}

//CreateArticlePictures Add a list of ArticlesPictures for an Articles.
func CreateArticlePictures(articleID uint, picturesID []uint) error {
	if len(picturesID) == 0 {
		return nil
	}

	data := make([]app.ArticlesPictures, len(picturesID))
	for i := 0; i < len(picturesID); i++ {
		data[i].ArticleId = articleID
		data[i].PictureId = picturesID[i]
	}

	return app.Session.Create(&data).Error
}

//GetArticlePictures Get all Pictures of an Articles
func GetArticlePictures(articleID uint) ([]app.Pictures, error) {
	var results []app.Pictures
	err := app.Session.Model(&app.Pictures{}).
		Joins("JOIN articles_pictures on articles_pictures.picture_id = pictures.id and  articles_pictures.article_id = ?", articleID).
		Scan(&results).Error

	return results, err
}

//DeleteArticlePictures Delete all ArticlesPictures of an Articles
func DeleteArticlePictures(articleID uint) error {
	return app.Session.Where("article_id = ?", articleID).Delete(&app.ArticlesPictures{}).Error
}

//DeleteArticlePicturesWithID Delete all ArticlesPictures of an Articles
// if invertDeletion is true we delete all ArticlesPictures where ID is not in picturesID
func DeleteArticlePicturesWithID(articleID uint, picturesID []uint, invertDeletion ...bool) error {
	if len(picturesID) == 0 {
		return nil
	}
	if len(invertDeletion) == 0 {
		return app.Session.Where("article_id = ? AND pictures_id IN ?", articleID, picturesID).Delete(&app.ArticlesPictures{}).Error
	}

	return app.Session.Where("article_id = ? AND pictures_id NOT IN ?", articleID, picturesID).Delete(&app.ArticlesPictures{}).Error
}
