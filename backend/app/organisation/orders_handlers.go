package organisation

type OrderPresenter struct {
	Orders
	Items []OrderItemPresenter
}

type OrderItemPresenter struct {
	*OrdersArticles
	Article ArticlesPresenter
}
