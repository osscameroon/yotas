package organisation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type OrderPresenter struct {
	*Orders
	Items []*OrderItemPresenter `json:"items"`
}

type OrderItemPresenter struct {
	*OrdersArticles
	Article ArticlesPresenter `json:"article"`
}

type OrderDecision struct {
	Accepted bool   `json:"accepted"`
	Reason   string `json:"reason"`
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

		articlesID := make([]uint, len(orderArticles))
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
			ordersPresenter[i].Items = append(ordersPresenter[i].Items, &OrderItemPresenter{
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

func GetUserOrdersHandler(ctx *gin.Context) {

	// TODO Retreive user information an user the same flow in GetOrganisationOrdersHandler
}

func GetOrderHandler(ctx *gin.Context) {

	/*	organisationID, err := strconv.Atoi(ctx.GetHeader("Tenant"))
		if err != nil {
			ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
			return
		}*/

	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Order id must be an int")
		return
	}

	// TODO check if the issuer of this request is the owner of the order or a manager of the organisation

	order, err := GetOrder(uint(orderID))
	if err != nil {
		return
	}
	orderPresenter := OrderPresenter{}
	orderPresenter.Orders = order

	orderArticles, err := GetOrderArticles(order.ID)
	if err != nil {
		return
	}

	articlesID := make([]uint, len(orderArticles))
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
		orderPresenter.Items = append(orderPresenter.Items, &OrderItemPresenter{
			OrdersArticles: &orderArticles[orderArticleIndex],
			Article: ArticlesPresenter{
				Articles: article,
				Pictures: pictures,
			},
		})
	}

	ctx.JSON(http.StatusOK, orderPresenter)
}

func CreateOrderHandler(ctx *gin.Context) {

	organisationID, err := strconv.Atoi(ctx.GetHeader("Tenant"))
	if err != nil || organisationID == 0 {
		ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
		return
	}

	order := OrderPresenter{}
	err = ctx.BindJSON(&order)
	if err != nil {
		ctx.String(http.StatusNotAcceptable, "Bad content")
		return
	}

	// TODO check if wallet id belong to this organisation

	if len(order.Items) == 0 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"messages": []string{"Can't create order with no items"}})
		return
	}

	// Check if each article belong to this organisation
	for index := 0; index < len(order.Items); index++ {
		article, err := GetOrganisationArticle(order.Items[index].ArticleID, uint(organisationID))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"messages": []string{
				fmt.Sprintf("The %v item doesn't belong to this organisation", index)}})
			return
		}

		order.Items[index].Article.Articles = *article
		order.Items[index].Article.Pictures, _ = GetArticlePictures(order.Items[index].ArticleID)
	}

	err = CreateOrder(order.Orders, order.Items)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"messages": []string{"Can't create order an error occur" + err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func ProcessOrderHandler(ctx *gin.Context) {

	/*	organisationID, err := strconv.Atoi(ctx.GetHeader("Tenant"))
		if err != nil || organisationID == 0 {
			ctx.String(http.StatusBadRequest, ErrTenantNotProvided.Error())
			return
		}*/

	//TODO check if the user is an admin of the organisation

	var decision OrderDecision
	err := ctx.BindJSON(&decision)
	if err != nil {
		ctx.String(http.StatusNotAcceptable, "Bad content")
		return
	}

	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Order id must be an int")
		return
	}

	_, err = GetOrder(uint(orderID))
	if err != nil {
		ctx.String(http.StatusNotFound, "Ressource not found")
		return
	}

	if decision.Accepted {
		err := AcceptOrder(uint(orderID), decision.Reason)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "An error occur when processing the order")
			return
		}
	} else {
		if strings.TrimSpace(decision.Reason) == "" {
			ctx.String(http.StatusBadRequest, "Please provide a reason when declining an order")
			return
		}
		err := DeclineOrder(uint(orderID), decision.Reason)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "An error occur when processing the order")
			return
		}
	}

	ctx.String(http.StatusOK, "Order processed")
}
