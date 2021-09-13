package orders

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/osscameroon/yotas/app"
	articles2 "github.com/osscameroon/yotas/app/shop/articles"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var orderFilterCriteria = map[string]OrderState{
	"not paid":  orderStateNotPaid,
	"to review": orderStateToReview,
	"accepted":  orderStateAccepted,
	"declined":  orderStateDeclined,
}

// GetOrganisationOrdersHandler Handler for organisation Orders
func GetOrganisationOrdersHandler(ctx *gin.Context) {
	organisationID, err := strconv.Atoi(ctx.GetHeader("Tenant"))
	if err != nil {
		ctx.String(http.StatusBadRequest, articles2.ErrTenantNotProvided.Error())
		return
	}

	// Initializing default
	limit := 1
	offset := 1
	stateFilter := ctx.Query("state")
	//search := ctx.Query("search")

	if value, err := strconv.Atoi(ctx.Query("limit")); err == nil && value > limit {
		limit = value
	}

	if value, err := strconv.Atoi(ctx.Query("offset")); err == nil && value > offset {
		offset = value
	}

	state := orderFilterCriteria[stateFilter]

	offset = (offset - 1) * limit
	orders, err := GetOrganisationOrders(uint(organisationID), state, limit, offset)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	}

	var ordersPresenter []app.OrderPresenter
	for orderIndex, order := range orders {
		var orderPresenter app.OrderPresenter
		orderPresenter.Orders = &orders[orderIndex]

		orderArticles, err := GetOrderArticles(order.ID)
		if err != nil {
			return
		}

		articlesID := make([]uint, len(orderArticles))
		for _, article := range orderArticles {
			articlesID = append(articlesID, article.ArticleID)
		}

		articlesList, err := articles2.GetArticles(articlesID)
		if err != nil {
			return
		}

		for orderArticleIndex, orderArticle := range orderArticles {
			// search for articles
			var articleToSave app.Articles
			for _, article := range articlesList {
				if article.ID == orderArticle.ArticleID {
					articleToSave = article
					break
				}
			}

			pictures, err := articles2.GetArticlePictures(articleToSave.ID)
			if err != nil {
				return
			}
			orderPresenter.Items = append(orderPresenter.Items, &app.OrderItemPresenter{
				OrdersArticles: &orderArticles[orderArticleIndex],
				Article: app.ArticlesPresenter{
					Articles: articleToSave,
					Pictures: pictures,
				},
			})
		}
		ordersPresenter = append(ordersPresenter, orderPresenter)
	}

	if len(ordersPresenter) == 0 {
		// to avoid return nil
		ordersPresenter = []app.OrderPresenter{}
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"limit":  limit,
		"offset": offset,
		"data":   ordersPresenter,
	})
}

func GetWalletOrdersHandler(ctx *gin.Context) {
	// Initializing default
	limit := 1
	offset := 1
	stateFilter := ctx.Query("state")
	walletID := ctx.Query("wallet")
	if strings.TrimSpace(walletID) == "" {
		ctx.String(http.StatusBadRequest, "Provide wallet id")
		return
	}

	if value, err := strconv.Atoi(ctx.Query("limit")); err == nil && value > limit {
		limit = value
	}

	if value, err := strconv.Atoi(ctx.Query("offset")); err == nil && value > offset {
		offset = value
	}

	state := orderFilterCriteria[stateFilter]

	offset = (offset - 1) * limit
	orders, err := GetWalletOrders(walletID, state, limit, offset)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	}

	var ordersPresenter []app.OrderPresenter
	for orderIndex, order := range orders {
		var orderPresenter app.OrderPresenter
		orderPresenter.Orders = &orders[orderIndex]

		orderArticles, err := GetOrderArticles(order.ID)
		if err != nil {
			return
		}

		articlesID := make([]uint, len(orderArticles))
		for _, article := range orderArticles {
			articlesID = append(articlesID, article.ArticleID)
		}

		articlesList, err := articles2.GetArticles(articlesID)
		if err != nil {
			return
		}

		for orderArticleIndex, orderArticle := range orderArticles {
			// search for articles
			var articleToSave app.Articles
			for _, article := range articlesList {
				if article.ID == orderArticle.ArticleID {
					articleToSave = article
					break
				}
			}

			pictures, err := articles2.GetArticlePictures(articleToSave.ID)
			if err != nil {
				return
			}
			orderPresenter.Items = append(orderPresenter.Items, &app.OrderItemPresenter{
				OrdersArticles: &orderArticles[orderArticleIndex],
				Article: app.ArticlesPresenter{
					Articles: articleToSave,
					Pictures: pictures,
				},
			})
		}
		ordersPresenter = append(ordersPresenter, orderPresenter)
	}

	if len(ordersPresenter) == 0 {
		// to avoid return nil
		ordersPresenter = []app.OrderPresenter{}
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"limit":  limit,
		"offset": offset,
		"data":   ordersPresenter,
	})
}

func GetOrderHandler(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Order id must be an int")
		return
	}

	// TODO check if the issuer of this request is the owner of the order or a manager of the organisation

	order, err := GetOrder(uint(orderID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.String(http.StatusNotFound, "Ressource not found")
		return
	}

	orderPresenter := app.OrderPresenter{}
	orderPresenter.Orders = order

	orderArticles, err := GetOrderArticles(order.ID)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	}

	articlesID := make([]uint, len(orderArticles))
	for _, orderArticle := range orderArticles {
		articlesID = append(articlesID, orderArticle.ArticleID)
	}

	articlesList, err := articles2.GetArticles(articlesID)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	}

	for _, orderArticle := range orderArticles {
		// search for articles
		var articleToSave app.Articles
		for _, article := range articlesList {
			if article.ID == orderArticle.ArticleID {
				articleToSave = article
				break
			}
		}

		pictures, err := articles2.GetArticlePictures(articleToSave.ID)
		if err != nil {
			log.Println(err)
			ctx.String(http.StatusInternalServerError, "An error occur")
			return
		}
		orderPresenter.Items = append(orderPresenter.Items, &app.OrderItemPresenter{
			OrdersArticles: &orderArticle,
			Article: app.ArticlesPresenter{
				Articles: articleToSave,
				Pictures: pictures,
			},
		})
	}

	ctx.JSON(http.StatusOK, orderPresenter)
}

func CreateOrderHandler(ctx *gin.Context) {
	organisationID, err := strconv.Atoi(ctx.GetHeader("Tenant"))
	if err != nil || organisationID == 0 {
		ctx.String(http.StatusBadRequest, articles2.ErrTenantNotProvided.Error())
		return
	}

	order := app.OrderPresenter{}
	err = ctx.BindJSON(&order)
	if err != nil {
		ctx.String(http.StatusNotAcceptable, "Bad content")
		return
	}

	if len(order.Items) == 0 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"messages": []string{"Can't create order with no items"}})
		return
	}

	// Check if each article belong to this organisation
	providedArticle := map[uint]uint{}
	for _, item := range order.Items {
		article, err := articles2.GetOrganisationArticle(item.ArticleID, uint(organisationID))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"messages": []string{
				fmt.Sprintf("The %v item doesn't belong to this organisation", item.ArticleID)}})
			return
		}

		// check if we already have an orderItem for this article
		if _, exist := providedArticle[item.ArticleID]; exist {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"messages": []string{
				fmt.Sprintf("Please provide one orderitem per article. Duplicate on article %s", article.Name)}})
			return
		}

		providedArticle[item.ArticleID] = item.ArticleID
		item.Article.Articles = *article
		item.Article.Pictures, _ = articles2.GetArticlePictures(item.ArticleID)
	}

	err = CreateOrder(order.Orders, order.Items)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"messages": []string{"Can't create order an error occur" + err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func DeleteOrderHandler(ctx *gin.Context) {
	//TODO check if the user is the owner of the order or an admin
	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Order id must be an int")
		return
	}

	order, err := GetOrder(uint(orderID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.String(http.StatusNotFound, "Ressource not found")
		return
	}

	if order.State != string(orderStateNotPaid) {
		ctx.String(http.StatusForbidden, "Can't delete order with state "+order.State)
		return
	}

	err = DeleteOrder(order.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "An error occur when deleting the order")
		return
	}

	ctx.String(http.StatusOK, "Order deleted")
}

func ProcessOrderHandler(ctx *gin.Context) {
	//TODO check if the user is an admin of the organisation
	var decision app.OrderDecision
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

	order, err := GetOrder(uint(orderID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.String(http.StatusNotFound, "Ressource not found")
		return
	}

	if order.State != string(orderStateToReview) {
		ctx.String(http.StatusForbidden, "Can't process order with state "+order.State)
		return
	}

	if decision.Accepted {
		err = AcceptOrder(uint(orderID), decision.Reason)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "An error occur when processing the order")
			return
		}
	} else {
		if strings.TrimSpace(decision.Reason) == "" {
			ctx.String(http.StatusBadRequest, "Please provide a reason when declining an order")
			return
		}
		err = DeclineOrder(uint(orderID), decision.Reason)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "An error occur when processing the order")
			return
		}
	}

	ctx.String(http.StatusOK, "Order processed")
}

func PayOrderHandler(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Order id must be an int")
		return
	}

	//TODO check if the user is the owner of the order
	order, err := GetOrder(uint(orderID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.String(http.StatusNotFound, "Ressource not found")
		return
	}

	if order.State != string(orderStateNotPaid) {
		ctx.String(http.StatusForbidden, "Can't pay order with state "+order.State)
		return
	}

	err = PayOrder(order.ID)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.String(http.StatusOK, "Payment processed successfully")
}

func UpdateOrderHandler(ctx *gin.Context) {
	organisationID, err := strconv.Atoi(ctx.GetHeader("Tenant"))
	if err != nil || organisationID == 0 {
		ctx.String(http.StatusBadRequest, articles2.ErrTenantNotProvided.Error())
		return
	}

	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Order id must be an int")
		return
	}

	//TODO check if the user is the owner of the order
	order, err := GetOrder(uint(orderID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "An error occur")
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.String(http.StatusNotFound, "Ressource not found")
		return
	}

	if order.State != string(orderStateNotPaid) {
		ctx.String(http.StatusForbidden, "Can't update order with state "+order.State)
		return
	}

	var orderPresenter app.OrderPresenter
	err = ctx.BindJSON(&orderPresenter)
	if err != nil {
		ctx.String(http.StatusNotAcceptable, "Bad content")
		return
	}

	if len(orderPresenter.Items) == 0 {
		ctx.String(http.StatusBadRequest, "Provide at least one order item")
		return
	}

	// Check if each article belong to this organisation
	providedArticle := map[uint]uint{}
	for _, item := range orderPresenter.Items {
		article, err := articles2.GetOrganisationArticle(item.ArticleID, uint(organisationID))
		if err != nil {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("The %v item doesn't belong to this organisation", item.ArticleID))
			return
		}

		// check if we already have an orderItem for this article
		if _, exist := providedArticle[item.ArticleID]; exist {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("Please provide one orderitem per article. Duplicate on article %s", article.Name))
			return
		}

		providedArticle[item.ArticleID] = item.ArticleID
		item.Article.Articles = *article
		item.Article.Pictures, _ = articles2.GetArticlePictures(item.ArticleID)
	}

	updatedOrder, err := UpdateOrder(order.ID, orderPresenter.Items)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, "An error occur try again later")
		return
	}

	orderPresenter.Orders = updatedOrder
	ctx.JSON(http.StatusOK, orderPresenter)
}
