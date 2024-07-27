package types

import (
	"context"

	"github.com/tobiasprima/kitchen/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}