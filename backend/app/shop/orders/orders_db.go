package orders

import (
	"errors"
	"fmt"
	"github.com/osscameroon/yotas/app"
	"gorm.io/gorm"
	"time"
)

type OrderState string

const (
	orderStateNotPaid  OrderState = "not paid"
	orderStateToReview OrderState = "to review"
	orderStateAccepted OrderState = "accepted"
	orderStateDeclined OrderState = "declined"
)

func CreateOrder(order *app.Orders, orderItems []*app.OrderItemPresenter) error {
	return app.Session.Transaction(func(tx *gorm.DB) error {
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
				return fmt.Errorf("can't create order process fail when creating order item %s", orderArticle.Article.Name)
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
}

func GetOrder(orderID uint) (*app.Orders, error) {
	var order app.Orders
	result := app.Session.Where("id = ?", orderID).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func GetWalletOrders(walletID string, orderStateFilter OrderState, limit int, offset int) ([]app.Orders, error) {
	var results []app.Orders
	req := app.Session.Model(&app.Orders{}).
		Joins("JOIN wallets on wallets.wallet_id = orders.wallet_id and orders.wallet_id = ?", walletID)

	if string(orderStateFilter) != "" {
		req = req.Where("state = ?", string(orderStateFilter))
	}

	err := req.Limit(limit).Offset(offset).Scan(&results).Error

	return results, err
}

func GetOrganisationOrders(organisationID uint, orderStateFilter OrderState, limit int, offset int) ([]app.Orders, error) {
	var results []app.Orders
	req := app.Session.Model(&app.Orders{}).
		Joins("JOIN wallets on wallets.wallet_id = orders.wallet_id and wallets.organisation_id = ?", organisationID)

	if string(orderStateFilter) != "" {
		req = req.Where("state = ?", string(orderStateFilter))
	}

	err := req.Limit(limit).Offset(offset).Scan(&results).Error

	return results, err
}

func GetOrderArticles(orderID uint) ([]app.OrdersArticles, error) {
	var results []app.OrdersArticles
	err := app.Session.Model(&app.OrdersArticles{}).
		Where("order_id = ?", orderID).
		Scan(&results).Error

	return results, err
}

func DeclineOrder(orderID uint, reason string) error {
	err := app.Session.Transaction(func(tx *gorm.DB) error {
		var order app.Orders
		err := tx.Model(&app.Orders{}).Where("id = ?", orderID).First(&order).Error
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
	err := app.Session.Transaction(func(tx *gorm.DB) error {
		var order app.Orders
		err := tx.Model(&app.Orders{}).Where("id = ?", orderID).First(&order).Error
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
	err := app.Session.Transaction(func(tx *gorm.DB) error {
		var order app.Orders
		err := tx.Model(&app.Orders{}).Where("id = ?", orderID).First(&order).Error
		if err != nil {
			return errors.New("can't find order")
		}

		// check if the wallet has enough yotas to pay the order
		var wallet app.Wallets
		err = tx.Model(&app.Wallets{}).Where("wallet_id = ?", order.WalletId).First(&wallet).Error
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
	err := app.Session.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&app.OrdersArticles{}).Where("order_id = ?", orderID).Update("deleted_at", time.Now().UTC()).Error
		if err != nil {
			return errors.New("can't delete order items")
		}

		err = tx.Model(&app.Orders{}).Where("id = ?", orderID).Update("deleted_at", time.Now().UTC()).Error
		if err != nil {
			return errors.New("can't delete order")
		}

		return nil
	})
	return err
}

func UpdateOrder(orderID uint, orderItems []*app.OrderItemPresenter) (*app.Orders, error) {
	var order app.Orders
	err := app.Session.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&app.Orders{}).Where("id = ?", orderID).First(&order).Error
		if err != nil {
			return err
		}

		var storedOrderArticles []app.OrdersArticles
		err = tx.Model(&app.OrdersArticles{}).
			Where("order_id = ?", orderID).
			Scan(&storedOrderArticles).Error
		if err != nil {
			return err
		}

		// we create each order OrdersArticles and summing the order amount
		var finalOrderItemsID []uint
		var totalOrderAmount int64
		for _, orderItem := range orderItems {
			//sum the amount of article
			//

			// find if this orderItems already exist in the list of the stored items
			orderItemAlreadyExist := false
			for _, storedOrderArticle := range storedOrderArticles {
				if orderItem.ArticleID != storedOrderArticle.ArticleID {
					continue
				}

				orderItemAlreadyExist = true
				storedOrderArticle.UpdatedAt = time.Now().UTC()
				storedOrderArticle.Quantity = orderItem.Quantity
				storedOrderArticle.ArticlePrice = orderItem.Article.Price
				err = tx.Save(&storedOrderArticle).Error
				if err != nil {
					return err
				}

				orderItem.OrdersArticles = &storedOrderArticle
				totalOrderAmount += int64(storedOrderArticle.Quantity) * storedOrderArticle.ArticlePrice
				finalOrderItemsID = append(finalOrderItemsID, storedOrderArticle.ID)
				break
			}

			if orderItemAlreadyExist {
				continue
			}

			orderItem.OrdersArticles.OrderID = order.ID
			orderItem.OrdersArticles.ArticlePrice = orderItem.Article.Price
			orderItem.OrdersArticles.CreatedAt = time.Now().UTC()
			orderItem.OrdersArticles.UpdatedAt = time.Now().UTC()
			err = tx.Create(orderItem.OrdersArticles).Error
			if err != nil {
				return fmt.Errorf("can't create order process fail when creating order item %s", orderItem.Article.Name)
			}

			totalOrderAmount += int64(orderItem.Quantity) * orderItem.ArticlePrice
			finalOrderItemsID = append(finalOrderItemsID, orderItem.ID)
		}

		// we delete unused previous store order articles
		err = tx.Model(&app.OrdersArticles{}).Where("order_id = ? AND id NOT IN ?", order.ID, finalOrderItemsID).Delete(&app.OrdersArticles{}).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		order.TotalAmount = totalOrderAmount
		order.UpdatedAt = time.Now().UTC()
		return tx.Save(&order).Error
	})

	if err != nil {
		return nil, err
	}

	return &order, nil
}
