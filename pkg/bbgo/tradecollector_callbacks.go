// Code generated by "callbackgen -type TradeCollector"; DO NOT EDIT.

package bbgo

import (
	"github.com/c9s/bbgo/pkg/types"
)

func (c *TradeCollector) OnTrade(cb func(trade types.Trade)) {
	c.tradeCallbacks = append(c.tradeCallbacks, cb)
}

func (c *TradeCollector) EmitTrade(trade types.Trade) {
	for _, cb := range c.tradeCallbacks {
		cb(trade)
	}
}

func (c *TradeCollector) OnPositionUpdate(cb func(position *Position)) {
	c.positionUpdateCallbacks = append(c.positionUpdateCallbacks, cb)
}

func (c *TradeCollector) EmitPositionUpdate(position *Position) {
	for _, cb := range c.positionUpdateCallbacks {
		cb(position)
	}
}
