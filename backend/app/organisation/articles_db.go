package organisation

import (
	"errors"
	"fmt"
	"github.com/osscameroon/yotas/db"
	"strings"
	"time"
)

//CreateArticle Add an Articles for an Organisations. The return an error if something when wrong
func CreateArticle(article *Articles) error {

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

	return db.Session.Create(article).Error
}

//GetArticle Retrieve an Articles an return a *Articles. If the article are not found the will return an nil pointer with an error
func GetArticle(articleID uint) (*Articles, error) {

	var article Articles
	result := db.Session.Where("id = ?", articleID).First(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}

//UpdateArticle update an Articles
func UpdateArticle(article *Articles) error {

	article.UpdatedAt = time.Now().UTC()
	result := db.Session.Save(article)
	return result.Error
}

//DeleteArticle delete an article with the given articleID
func DeleteArticle(articleID uint) error {
	result := db.Session.Model(&Articles{}).Where("id = ?", articleID).Update("deleted_at", time.Now().UTC())
	return result.Error
}

//CreateOrganisationArticle Create a new OrganisationsArticles
func CreateOrganisationArticle(organisationID uint, articleID uint) error {

	return db.Session.Create(&OrganisationsArticles{
		OrganisationId: organisationID,
		ArticleId:      articleID,
		Model:          db.Model{CreatedAt: time.Now().UTC()},
	}).Error
}

//GetOrganisationArticles Get a list of Articles related to an Organisations
func GetOrganisationArticles(organisationID uint, categoryId string, limit int, offset int, search string, priceGte int, priceLte int, sort string) ([]Articles, error) {

	results := []Articles{}
	search = fmt.Sprintf("%s%s%s", "%", search, "%")
	err := db.Session.Model(&Articles{}).
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

	data := make([]ArticlesPictures, len(picturesID))
	for i := 0; i < len(picturesID); i++ {
		data[i].ArticleId = articleID
		data[i].PictureId = picturesID[i]
	}

	return db.Session.Create(&data).Error
}

//GetArticlePictures Get all Pictures of an Articles
func GetArticlePictures(articleID uint) ([]Pictures, error) {

	results := []Pictures{}
	err := db.Session.Model(&Pictures{}).
		Joins("JOIN articles_pictures on articles_pictures.picture_id = pictures.id and  articles_pictures.article_id = ?", articleID).
		Distinct("pictures.id").
		Scan(&results).Error

	return results, err
}

//DeleteArticlePictures Delete all ArticlesPictures of an Articles
func DeleteArticlePictures(articleID uint) error {
	return db.Session.Where("article_id = ?", articleID).Delete(&ArticlesPictures{}).Error
}

//DeleteArticlePicturesWithID Delete all ArticlesPictures of an Articles
// if invertDeletion is true we delete all ArticlesPictures where ID is not in picturesID
func DeleteArticlePicturesWithID(articleID uint, picturesID []uint, invertDeletion ...bool) error {

	if len(picturesID) == 0 {
		return nil
	}
	if len(invertDeletion) == 0 {
		return db.Session.Where("article_id = ? AND pictures_id IN ?", articleID, picturesID).Delete(&ArticlesPictures{}).Error
	}

	return db.Session.Where("article_id = ? AND pictures_id NOT IN ?", articleID, picturesID).Delete(&ArticlesPictures{}).Error

}
