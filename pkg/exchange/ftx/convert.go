package ftx

import (
	"fmt"
	"strings"

	"github.com/c9s/bbgo/pkg/datatype"
	"github.com/c9s/bbgo/pkg/fixedpoint"
	"github.com/c9s/bbgo/pkg/types"
)

func toGlobalCurrency(original string) string {
	return TrimUpperString(original)
}

func toGlobalSymbol(original string) string {
	return TrimUpperString(original)
}

func TrimUpperString(original string) string {
	return strings.ToUpper(strings.TrimSpace(original))
}

func TrimLowerString(original string) string {
	return strings.ToLower(strings.TrimSpace(original))
}

var errUnsupportedOrderStatus = fmt.Errorf("unsupported order status")

func toGlobalOrder(r order) (types.Order, error) {
	// In exchange/max/convert.go, it only parses these fields.
	o := types.Order{
		SubmitOrder: types.SubmitOrder{
			ClientOrderID: r.ClientId,
			Symbol:        toGlobalSymbol(r.Market),
			Side:          types.SideType(TrimUpperString(r.Side)),
			// order type definition: https://github.com/ftexchange/ftx/blob/master/rest/client.py#L122
			Type:        types.OrderType(TrimUpperString(r.Type)),
			Quantity:    r.Size,
			Price:       r.Price,
			TimeInForce: "GTC",
		},
		Exchange:         types.ExchangeFTX.String(),
		IsWorking:        r.Status == "open",
		OrderID:          uint64(r.ID),
		Status:           "",
		ExecutedQuantity: r.FilledSize,
		CreationTime:     datatype.Time(r.CreatedAt),
		UpdateTime:       datatype.Time(r.CreatedAt),
	}

	// `new` (accepted but not processed yet), `open`, or `closed` (filled or cancelled)
	switch r.Status {
	case "new":
		o.Status = types.OrderStatusNew
	case "open":
		if fixedpoint.NewFromFloat(o.ExecutedQuantity) != fixedpoint.NewFromInt(0) {
			o.Status = types.OrderStatusPartiallyFilled
		} else {
			o.Status = types.OrderStatusNew
		}
	case "closed":
		// filled or canceled
		if fixedpoint.NewFromFloat(o.Quantity) == fixedpoint.NewFromFloat(o.ExecutedQuantity) {
			o.Status = types.OrderStatusFilled
		} else {
			// can't distinguish it's canceled or rejected from order response, so always set to canceled
			o.Status = types.OrderStatusCanceled
		}
	default:
		return types.Order{}, fmt.Errorf("unsupported status %s: %w", r.Status, errUnsupportedOrderStatus)
	}

	return o, nil
}
