package organisation

import (
	"errors"
	"github.com/osscameroon/yotas/db"
	"gorm.io/gorm"
)

const (
	operationTypeRefund OperationType = "refund"
	operationTypeCredit OperationType = "credit"
	operationTypeDebit  OperationType = "debit"
)

//	CreateOperation create an operation
//	If you create the operation inside a transaction, you must provide the transaction pointer to the param transaction
//	If not you can set it to nil (we will internally use the default db.Session to create operation)
func CreateOperation(transaction *gorm.DB, walletID string, operationType OperationType, amount int64) (string, error) {

	switch operationType {
	case operationTypeRefund, operationTypeCredit, operationTypeDebit:
		break
	default:
		return "", errors.New("unallowed operation type")
	}

	if transaction == nil {
		transaction = db.Session
	}

	transaction.Transaction(func(tx *gorm.DB) error {
		return nil
	})

	return "", nil
}
