package orders

import (
	"github.com/osscameroon/yotas/app"
)

func OrderRouter() {
	router := app.GetApiRouter()

	router.GET("/shop", GetOrganisationOrdersHandler)
	router.GET("/shop/wallet", GetWalletOrdersHandler)
	router.GET("/shop/:orderID", GetOrderHandler)
	router.POST("/shop", CreateOrderHandler)
	router.POST("/shop/:orderID/process", ProcessOrderHandler)
	router.POST("/shop/:orderID/pay", PayOrderHandler)
	router.PUT("/shop/:orderID", UpdateOrderHandler)
	router.DELETE("/shop/:orderID", DeleteOrderHandler)
}
