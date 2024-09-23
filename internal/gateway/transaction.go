package gateway

import "github.com/gaspartv/go-fullcycle-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
