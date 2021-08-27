package organisation

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderPresenter struct {
	*Orders
	Items []OrderItemPresenter
}

type OrderItemPresenter struct {
	*OrdersArticles
	Article ArticlesPresenter
}

func GetOrganisationOrdersHandler(ctx *gin.Context) {

	organisationID, err := strconv.Atoi(ctx.GetHeader("Tenant"))
	if err != nil {
		ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
		return
	}

	// Initializing default
	limit := 1
	offset := 1
	//search := ctx.Query("search")

	if value, err := strconv.Atoi(ctx.Query("limit")); err == nil && value > limit {
		limit = value
	}

	if value, err := strconv.Atoi(ctx.Query("offset")); err == nil && value > offset {
		offset = value
	}

	offset = (offset - 1) * limit
	orders, err := GetOrganisationOrders(uint(organisationID), limit, offset)
	if err != nil {
		return
	}

	ordersPresenter := make([]OrderPresenter, len(orders))
	for i := 0; i < len(orders); i++ {
		ordersPresenter[i].Orders = &orders[i]

		orderArticles, err := GetOrderArticles(orders[i].ID)
		if err != nil {
			return
		}

		articlesID := make([]uint, len(orders))
		for orderArticleIndex := 0; orderArticleIndex < len(orderArticles); orderArticleIndex++ {
			articlesID = append(articlesID, orderArticles[orderArticleIndex].ArticleID)
		}

		articles, err := GetArticles(articlesID)
		if err != nil {
			return
		}

		for orderArticleIndex := 0; orderArticleIndex < len(orderArticles); orderArticleIndex++ {
			// search for articles
			article := Articles{}
			for articleIndex := 0; articleIndex < len(articles); articleIndex++ {
				if articles[articleIndex].ID == orderArticles[orderArticleIndex].ArticleID {
					article = articles[articleIndex]
					break
				}
			}

			pictures, err := GetArticlePictures(article.ID)
			if err != nil {
				return
			}
			ordersPresenter[i].Items = append(ordersPresenter[i].Items, OrderItemPresenter{
				OrdersArticles: &orderArticles[orderArticleIndex],
				Article: ArticlesPresenter{
					Articles: article,
					Pictures: pictures,
				},
			})
		}
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"limit":  limit,
		"offset": offset,
		"data":   ordersPresenter,
	})
}

func CreateOrderHandler(ctx *gin.Context) {

}
func GetUserOrdersHandler(ctx *gin.Context) {

}

func GetOrderHandler(ctx *gin.Context) {

}
