package orders

import (
	"github.com/osscameroon/yotas/app"
)

func OrderRouter() {
	router := app.GetApiRouter()

	router.GET("/orders", GetOrganisationOrdersHandler)
	router.GET("/orders/wallet", GetWalletOrdersHandler)
	router.GET("/orders/:orderID", GetOrderHandler)
	router.POST("/orders", CreateOrderHandler)
	router.POST("/orders/:orderID/process", ProcessOrderHandler)
	router.POST("/orders/:orderID/pay", PayOrderHandler)
	router.PUT("/orders/:orderID", UpdateOrderHandler)
	router.DELETE("/orders/:orderID", DeleteOrderHandler)
}
