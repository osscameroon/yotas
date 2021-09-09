package organisation

import (
	"errors"
	"fmt"
	"github.com/osscameroon/yotas/db"
	"gorm.io/gorm"
	"time"
)

type OrderState string
type OperationType string

const (
	orderStateNotPaid  OrderState = "not paid"
	orderStateToReview OrderState = "to review"
	orderStateAccepted OrderState = "accepted"
	orderStateDeclined OrderState = "declined"
)

func CreateOrder(order *Orders, orderItems []*OrderItemPresenter) error {
	err := db.Session.Transaction(func(tx *gorm.DB) error {

		// We first create the order
		order.State = string(orderStateNotPaid)
		order.CreatedAt = time.Now().UTC()
		err := tx.Create(order).Error
		if err != nil {
			return errors.New("can't create order")
		}

		var totalOrderAmount int64
		// we create each order OrdersArticles and summing the order amount
		for _, orderArticle := range orderItems {
			//sum the amount of article
			totalOrderAmount += int64(orderArticle.Quantity) * orderArticle.Article.Price

			orderArticle.OrderID = order.ID
			orderArticle.ArticlePrice = orderArticle.Article.Price
			orderArticle.CreatedAt = time.Now().UTC()
			err = tx.Create(orderArticle.OrdersArticles).Error
			if err != nil {
				return fmt.Errorf("can't create orders process fail when creating order item %s", orderArticle.Article.Name)
			}
		}

		// finally, we update the order amount
		order.TotalAmount = totalOrderAmount
		order.UpdatedAt = time.Now().UTC()
		err = tx.Save(order).Error
		if err != nil {
			return errors.New("can't create order")
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

func DeclineOrder(orderID uint, reason string) error {

	err := db.Session.Transaction(func(tx *gorm.DB) error {

		var order Orders
		err := tx.Model(&Orders{}).Where("id = ?", orderID).First(&order).Error
		if err != nil {
			return errors.New("can't find order")
		}

		order.State = string(orderStateDeclined)
		order.Decision = reason
		order.UpdatedAt = time.Now().UTC()

		err = tx.Save(&order).Error
		if err != nil {
			return errors.New("can't update order state")
		}

		operationDescription := fmt.Sprintf("Refund your wallet due to the organisation admin decision to cancel your order %d. \nDetails: %s", orderID, reason)
		_, err = CreateOperation(tx, order.WalletId, operationTypeRefund, order.TotalAmount, operationDescription)
		return err
	})

	return err
}

func AcceptOrder(orderID uint, reason string) error {
	err := db.Session.Transaction(func(tx *gorm.DB) error {
		var order Orders
		err := tx.Model(&Orders{}).Where("id = ?", orderID).First(&order).Error
		if err != nil {
			return errors.New("can't find order")
		}

		order.State = string(orderStateAccepted)
		order.Decision = reason
		order.UpdatedAt = time.Now().UTC()

		err = tx.Save(&order).Error
		if err != nil {
			return errors.New("can't update order state")
		}
		return nil
	})

	return err
}

func PayOrder(orderID uint) error {
	err := db.Session.Transaction(func(tx *gorm.DB) error {

		var order Orders
		err := tx.Model(&Orders{}).Where("id = ?", orderID).First(&order).Error
		if err != nil {
			return errors.New("can't find order")
		}

		// check if the wallet has enough yotas to pay the order
		var wallet Wallets
		err = tx.Model(&Wallets{}).Where("wallet_id = ?", order.WalletId).First(&wallet).Error
		if err != nil {
			return errors.New("can't find user wallet")
		}

		if wallet.Balance < order.TotalAmount {
			return errors.New("wallet doesn't have enough money")
		}

		order.State = string(orderStateToReview)
		order.UpdatedAt = time.Now().UTC()
		err = tx.Save(&order).Error
		if err != nil {
			return errors.New("can't update order state")
		}

		operationDescription := fmt.Sprintf("Debit your wallet due to your order %d.", orderID)
		_, err = CreateOperation(tx, order.WalletId, operationTypeDebit, order.TotalAmount, operationDescription)
		return err
	})

	return err
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
