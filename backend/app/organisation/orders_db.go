package organisation

import (
	"errors"
	"fmt"
	"github.com/osscameroon/yotas/db"
	"gorm.io/gorm"
	"time"
)

type OrderState string

const (
	orderStateToReview OrderState = "to review"
	orderStateAccepted OrderState = "accepted"
	orderStateDeclined OrderState = "declined"
)

func CreateOrder(order *Orders, orderItems []*OrderItemPresenter) error {
	err := db.Session.Transaction(func(tx *gorm.DB) error {

		order.State = string(orderStateToReview)
		err := tx.Create(order).Error
		if err != nil {
			return errors.New("can't create orders")
		}

		for _, orderArticle := range orderItems {
			orderArticle.OrderID = order.ID
			err = tx.Create(orderArticle.OrdersArticles).Error
			if err != nil {
				return fmt.Errorf("can't create orders process fail when creating order item %s", orderArticle.Article.Name)
			}
		}

		return nil
	})

	return err
}

func GetOrder(orderID uint) (*Orders, error) {
	var order Orders
	result := db.Session.Where("id = ?", orderID).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func GetWalletOrders(walletID string, limit int, offset int) ([]Orders, error) {
	results := []Orders{}
	err := db.Session.Model(&Orders{}).
		Joins("JOINS wallets on wallets.wallet_id = orders.wallet_id and orders.wallet_id = ?", walletID).
		Limit(limit).
		Offset(offset).
		Scan(&results).Error

	return results, err
}

func GetOrganisationOrders(organisationID uint, limit int, offset int) ([]Orders, error) {
	results := []Orders{}
	err := db.Session.Model(&Orders{}).
		Joins("JOIN wallets on wallets.wallet_id = orders.wallet_id and wallets.organisation_id = ?", organisationID).
		Limit(limit).
		Offset(offset).
		Scan(&results).Error

	return results, err
}

func GetOrderArticles(orderID uint) ([]OrdersArticles, error) {
	results := []OrdersArticles{}
	err := db.Session.Model(&OrdersArticles{}).
		Where("order_id = ?", orderID).
		Scan(&results).Error

	return results, err
}

func DeclineOrder(orderID uint, state OrderState) {

}

func AcceptOrder() {

}

func PayOrder(orderID uint) {

}

func DeleteOrder(orderID uint) error {
	err := db.Session.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&OrdersArticles{}).Where("order_id = ?", orderID).Update("deleted_at", time.Now().UTC()).Error
		if err != nil {
			return errors.New("can't delete order items")
		}

		err = tx.Model(&Orders{}).Where("id = ?", orderID).Update("deleted_at", time.Now().UTC()).Error
		if err != nil {
			return errors.New("can't delete order")
		}

		return nil
	})
	return err
}
