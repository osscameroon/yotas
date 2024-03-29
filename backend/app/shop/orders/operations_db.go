package orders

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/osscameroon/yotas/app"
	"gorm.io/gorm"
	"time"
)

type OperationType string

const (
	operationTypeRefund OperationType = "refund"
	operationTypeCredit OperationType = "credit"
	operationTypeDebit  OperationType = "debit"
)

//	CreateOperation create an operation
//	If you create the operation inside a transaction, you must provide the transaction pointer to the param transaction
//	If not you can set it to nil (we will internally use the default db.Session to create operation)
func CreateOperation(transaction *gorm.DB, walletID string, operationType OperationType, amount int64, description string) (*app.Operations, error) {
	switch operationType {
	case operationTypeRefund, operationTypeCredit, operationTypeDebit:
		break
	default:
		return nil, errors.New("unallowed operation type")
	}

	if transaction == nil {
		transaction = app.Session
	}

	var operationResult *app.Operations
	err := transaction.Transaction(func(tx *gorm.DB) error {
		// get the wallet
		var wallet app.Wallets
		err := tx.Model(&app.Wallets{}).Where("wallet_id = ?", walletID).First(&wallet).Error
		if err != nil {
			return errors.New("can't find the wallet")
		}

		// create the operation without the operation hash
		operation := app.Operations{
			Model:         app.Model{CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()},
			Amount:        amount,
			WalletId:      walletID,
			Description:   description,
			OperationType: string(operationType),
			Approved:      true,
			//due to nullity constraint in migrations we give an initial random value to our operation hash
			OperationHash: fmt.Sprintf("%v", time.Now().UTC()),
		}
		err = tx.Create(&operation).Error
		if err != nil {
			return errors.New("can't create operation")
		}

		// create the hash
		finalHash := ""

		//  we check if we have previous operations
		lastOperation := app.Operations{}
		err = tx.Model(&app.Operations{}).Where("id != ?", operation.ID).Order("created_at DESC").First(&lastOperation).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("can't create operation")
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			//previous record exist we append the new hash
			finalHash, err = combineOperationHash(lastOperation.OperationHash, operation)
			if err != nil {
				return errors.New("can't create operation")
			}
		} else {
			//no previous record exist we generate a hash referer to this operation only
			finalHash, err = generateOperationHash(operation)
			if err != nil {
				return errors.New("can't create operation")
			}
		}

		// we omit the field updated_at because we don't want his value to change automatically by gorm
		err = tx.Model(&app.Operations{}).Where("id = ?", operation.ID).Omit("updated_at").Update("operation_hash", finalHash).Error
		if err != nil {
			return errors.New("can't create operation")
		}

		operationResult = &operation

		// update the wallet balance
		switch operationType {
		case operationTypeRefund, operationTypeCredit:
			wallet.Balance += amount
			wallet.UpdatedAt = time.Now().UTC()
			err = tx.Save(&wallet).Error
			if err != nil {
				return fmt.Errorf("can't %s wallet", operationType)
			}
		case operationTypeDebit:

			// check if the wallet has enough yotas to pay the order
			if wallet.Balance < amount {
				return errors.New("wallet doesn't have enough yotas")
			}

			wallet.Balance -= amount
			wallet.UpdatedAt = time.Now().UTC()
			err = tx.Save(&wallet).Error
			if err != nil {
				return fmt.Errorf("can't %s wallet", operationType)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return operationResult, nil
}

//generateOperationHash generate a hash for an operation
func generateOperationHash(operations app.Operations) (string, error) {
	operations.OperationHash = ""
	jsonStr, err := json.Marshal(operations)
	if err != nil {
		return "", errors.New("can't marshal operation")
	}

	//We create the hash
	h := sha256.New()
	_, _ = h.Write(jsonStr)
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

//combineOperationHash is used to combine hash of a previous operation within a hash for a current operation
func combineOperationHash(previousHash string, operations app.Operations) (string, error) {
	operations.OperationHash = ""
	jsonStr, err := json.Marshal(operations)
	if err != nil {
		return "", errors.New("can't marshal operation")
	}

	dataToHash := append([]byte(previousHash), jsonStr...)

	//We create the hash
	h := sha256.New()
	_, _ = h.Write(dataToHash)
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
