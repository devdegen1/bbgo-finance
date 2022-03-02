package ftx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/c9s/bbgo/pkg/exchange/ftx/ftxapi"
	"github.com/c9s/bbgo/pkg/fixedpoint"
	"github.com/c9s/bbgo/pkg/types"
)

const (
	restEndpoint       = "https://ftx.com"
	defaultHTTPTimeout = 15 * time.Second
)

var logger = logrus.WithField("exchange", "ftx")

// POST https://ftx.com/api/orders 429, Success: false, err: Do not send more than 2 orders on this market per 200ms
var requestLimit = rate.NewLimiter(rate.Every(220*time.Millisecond), 2)

//go:generate go run generate_symbol_map.go

type Exchange struct {
	client *ftxapi.RestClient

	key, secret  string
	subAccount   string
	restEndpoint *url.URL
}

type MarketTicker struct {
	Market types.Market
	Price  fixedpoint.Value
	Ask    fixedpoint.Value
	Bid    fixedpoint.Value
	Last   fixedpoint.Value
}

type MarketMap map[string]MarketTicker

// FTX does not have broker ID
const spotBrokerID = "BBGO"

func newSpotClientOrderID(originalID string) (clientOrderID string) {
	prefix := "x-" + spotBrokerID
	prefixLen := len(prefix)

	if originalID != "" {
		// try to keep the whole original client order ID if user specifies it.
		if prefixLen+len(originalID) > 32 {
			return originalID
		}

		clientOrderID = prefix + originalID
		return clientOrderID
	}

	clientOrderID = uuid.New().String()
	clientOrderID = prefix + clientOrderID
	if len(clientOrderID) > 32 {
		return clientOrderID[0:32]
	}

	return clientOrderID
}

func NewExchange(key, secret string, subAccount string) *Exchange {
	u, err := url.Parse(restEndpoint)
	if err != nil {
		panic(err)
	}

	client := ftxapi.NewClient()
	client.Auth(key, secret, subAccount)
	return &Exchange{
		client:       client,
		restEndpoint: u,
		key:          key,
		secret:       secret,
		subAccount:   subAccount,
	}
}

func (e *Exchange) newRest() *restRequest {
	r := newRestRequest(&http.Client{Timeout: defaultHTTPTimeout}, e.restEndpoint).Auth(e.key, e.secret)
	if len(e.subAccount) > 0 {
		r.SubAccount(e.subAccount)
	}
	return r
}

func (e *Exchange) Name() types.ExchangeName {
	return types.ExchangeFTX
}

func (e *Exchange) PlatformFeeCurrency() string {
	return toGlobalCurrency("FTT")
}

func (e *Exchange) NewStream() types.Stream {
	return NewStream(e.key, e.secret, e.subAccount, e)
}

func (e *Exchange) QueryMarkets(ctx context.Context) (types.MarketMap, error) {
	markets, err := e._queryMarkets(ctx)
	if err != nil {
		return nil, err
	}
	marketMap := types.MarketMap{}
	for k, v := range markets {
		marketMap[k] = v.Market
	}
	return marketMap, nil
}

func (e *Exchange) _queryMarkets(ctx context.Context) (MarketMap, error) {
	req := e.client.NewGetMarketsRequest()
	ftxMarkets, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}

	markets := MarketMap{}
	for _, m := range ftxMarkets {
		symbol := toGlobalSymbol(m.Name)
		symbolMap[symbol] = m.Name

		mkt2 := MarketTicker{
			Market: types.Market{
				Symbol:      symbol,
				LocalSymbol: m.Name,
				// The max precision is length(DefaultPow). For example, currently fixedpoint.DefaultPow
				// is 1e8, so the max precision will be 8.
				PricePrecision:  m.PriceIncrement.NumFractionalDigits(),
				VolumePrecision: m.SizeIncrement.NumFractionalDigits(),
				QuoteCurrency:   toGlobalCurrency(m.QuoteCurrency),
				BaseCurrency:    toGlobalCurrency(m.BaseCurrency),
				// FTX only limit your order by `MinProvideSize`, so I assign zero value to unsupported fields:
				// MinNotional, MinAmount, MaxQuantity, MinPrice and MaxPrice.
				MinNotional: fixedpoint.Zero,
				MinAmount:   fixedpoint.Zero,
				MinQuantity: m.MinProvideSize,
				MaxQuantity: fixedpoint.Zero,
				StepSize:    m.SizeIncrement,
				MinPrice:    fixedpoint.Zero,
				MaxPrice:    fixedpoint.Zero,
				TickSize:    m.PriceIncrement,
			},
			Price: m.Price,
			Bid:   m.Bid,
			Ask:   m.Ask,
			Last:  m.Last,
		}
		markets[symbol] = mkt2
	}
	return markets, nil
}

func (e *Exchange) QueryAccount(ctx context.Context) (*types.Account, error) {

	req := e.client.NewGetAccountRequest()
	ftxAccount, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}

	a := &types.Account{
		MakerCommission:   ftxAccount.MakerFee,
		TakerCommission:   ftxAccount.TakerFee,
		TotalAccountValue: ftxAccount.TotalAccountValue,
	}

	balances, err := e.QueryAccountBalances(ctx)
	if err != nil {
		return nil, err
	}

	a.UpdateBalances(balances)
	return a, nil
}

func (e *Exchange) QueryAccountBalances(ctx context.Context) (types.BalanceMap, error) {
	balanceReq := e.client.NewGetBalancesRequest()
	ftxBalances, err := balanceReq.Do(ctx)
	if err != nil {
		return nil, err
	}

	var balances = make(types.BalanceMap)
	for _, r := range ftxBalances {
		balances[toGlobalCurrency(r.Coin)] = types.Balance{
			Currency:  toGlobalCurrency(r.Coin),
			Available: r.Free,
			Locked:    r.Total.Sub(r.Free),
		}
	}

	return balances, nil
}

// resolution field in api
// window length in seconds. options: 15, 60, 300, 900, 3600, 14400, 86400, or any multiple of 86400 up to 30*86400
var supportedIntervals = map[types.Interval]int{
	types.Interval1m:  1,
	types.Interval5m:  5,
	types.Interval15m: 15,
	types.Interval1h:  60,
	types.Interval1d:  60 * 24,
	types.Interval3d:  60 * 24 * 3,
}

func (e *Exchange) SupportedInterval() map[types.Interval]int {
	return supportedIntervals
}

func (e *Exchange) IsSupportedInterval(interval types.Interval) bool {
	return isIntervalSupportedInKLine(interval)
}

func (e *Exchange) QueryKLines(ctx context.Context, symbol string, interval types.Interval, options types.KLineQueryOptions) ([]types.KLine, error) {
	var klines []types.KLine
	var since, until, currentEnd time.Time
	if options.StartTime != nil {
		since = *options.StartTime
	}
	if options.EndTime != nil {
		until = *options.EndTime
	} else {
		until = time.Now()
	}

	currentEnd = until

	for {

		// the fetch result is from newest to oldest
		endTime := currentEnd.Add(interval.Duration())
		options.EndTime = &endTime
		lines, err := e._queryKLines(ctx, symbol, interval, types.KLineQueryOptions{
			StartTime: &since,
			EndTime:   &currentEnd,
		})

		if err != nil {
			return nil, err
		}

		if len(lines) == 0 {
			break
		}

		for _, line := range lines {

			if line.StartTime.Unix() < currentEnd.Unix() {
				currentEnd = line.StartTime.Time()
			}

			if line.StartTime.Unix() > since.Unix() {
				klines = append(klines, line)
			}
		}

		if len(lines) == 1 && lines[0].StartTime.Unix() == currentEnd.Unix() {
			break
		}

		outBound := currentEnd.Add(interval.Duration()*-1).Unix() <= since.Unix()
		if since.IsZero() || currentEnd.Unix() == since.Unix() || outBound {
			break
		}

		if options.Limit != 0 && options.Limit <= len(lines) {
			break
		}
	}
	sort.Slice(klines, func(i, j int) bool { return klines[i].StartTime.Unix() < klines[j].StartTime.Unix() })

	if options.Limit != 0 {
		limitedItems := len(klines) - options.Limit
		if limitedItems > 0 {
			return klines[limitedItems:], nil
		}
	}

	return klines, nil
}

func (e *Exchange) _queryKLines(ctx context.Context, symbol string, interval types.Interval, options types.KLineQueryOptions) ([]types.KLine, error) {
	var since, until time.Time
	if options.StartTime != nil {
		since = *options.StartTime
	}
	if options.EndTime != nil {
		until = *options.EndTime
	} else {
		until = time.Now()
	}
	if since.After(until) {
		return nil, fmt.Errorf("invalid query klines time range, since: %+v, until: %+v", since, until)
	}
	if !isIntervalSupportedInKLine(interval) {
		return nil, fmt.Errorf("interval %s is not supported", interval.String())
	}

	if err := requestLimit.Wait(ctx); err != nil {
		return nil, err
	}

	resp, err := e.newRest().HistoricalPrices(ctx, toLocalSymbol(symbol), interval, 0, since, until)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return nil, fmt.Errorf("ftx returns failure")
	}

	var klines []types.KLine
	for _, r := range resp.Result {
		globalKline, err := toGlobalKLine(symbol, interval, r)
		if err != nil {
			return nil, err
		}
		klines = append(klines, globalKline)
	}
	return klines, nil
}

func isIntervalSupportedInKLine(interval types.Interval) bool {
	_, ok := supportedIntervals[interval]
	return ok
}

func (e *Exchange) QueryTrades(ctx context.Context, symbol string, options *types.TradeQueryOptions) ([]types.Trade, error) {
	var since, until time.Time
	if options.StartTime != nil {
		since = *options.StartTime
	}
	if options.EndTime != nil {
		until = *options.EndTime
	} else {
		until = time.Now()
	}

	if since.After(until) {
		return nil, fmt.Errorf("invalid query trades time range, since: %+v, until: %+v", since, until)
	}

	if options.Limit == 1 {
		// FTX doesn't provide pagination api, so we have to split the since/until time range into small slices, and paginate ourselves.
		// If the limit is 1, we always get the same data from FTX.
		return nil, fmt.Errorf("limit can't be 1 which can't be used in pagination")
	}

	limit := options.Limit
	if limit == 0 {
		limit = 200
	}

	tradeIDs := make(map[uint64]struct{})

	lastTradeID := options.LastTradeID
	var trades []types.Trade
	symbol = strings.ToUpper(symbol)

	for since.Before(until) {
		req := e.client.NewGetFillsRequest()
		req.Market(toLocalSymbol(symbol))
		req.StartTime(since)
		req.EndTime(until)
		req.Order("asc")
		fills, err := req.Do(ctx)
		if err != nil {
			return nil, err
		}

		sort.Slice(fills, func(i, j int) bool {
			return fills[i].Id < fills[j].Id
		})

		for _, fill := range fills {
			// always update since to avoid infinite loop
			since = fill.Time

			if _, ok := tradeIDs[fill.Id]; ok {
				continue
			}

			if fill.Id <= lastTradeID || fill.Time.Before(since) || fill.Time.After(until) {
				continue
			}

			tradeIDs[fill.Id] = struct{}{}
			lastTradeID = fill.Id

			t, err := toGlobalTrade(fill)
			if err != nil {
				return nil, err
			}
			trades = append(trades, t)
		}

		if int64(len(fills)) < limit {
			return trades, nil
		}
	}

	return trades, nil
}

func (e *Exchange) QueryDepositHistory(ctx context.Context, asset string, since, until time.Time) (allDeposits []types.Deposit, err error) {
	if until == (time.Time{}) {
		until = time.Now()
	}
	if since.After(until) {
		return nil, fmt.Errorf("invalid query deposit history time range, since: %+v, until: %+v", since, until)
	}
	asset = TrimUpperString(asset)

	resp, err := e.newRest().DepositHistory(ctx, since, until, 0)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return nil, fmt.Errorf("ftx returns failure")
	}
	sort.Slice(resp.Result, func(i, j int) bool {
		return resp.Result[i].Time.Before(resp.Result[j].Time.Time)
	})
	for _, r := range resp.Result {
		d, err := toGlobalDeposit(r)
		if err != nil {
			return nil, err
		}
		if d.Asset == asset && !since.After(d.Time.Time()) && !until.Before(d.Time.Time()) {
			allDeposits = append(allDeposits, d)
		}
	}
	return
}

func (e *Exchange) SubmitOrders(ctx context.Context, orders ...types.SubmitOrder) (types.OrderSlice, error) {
	var createdOrders types.OrderSlice
	// TODO: currently only support limit and market order
	// TODO: support time in force
	for _, so := range orders {
		if err := requestLimit.Wait(ctx); err != nil {
			logrus.WithError(err).Error("rate limit error")
		}

		orderType, err := toLocalOrderType(so.Type)
		if err != nil {
			logrus.WithError(err).Error("type error")
		}

		or, err := e.newRest().PlaceOrder(ctx, PlaceOrderPayload{
			Market:     toLocalSymbol(TrimUpperString(so.Symbol)),
			Side:       TrimLowerString(string(so.Side)),
			Price:      so.Price,
			Type:       string(orderType),
			Size:       so.Quantity,
			ReduceOnly: false,
			IOC:        so.TimeInForce == types.TimeInForceIOC,
			PostOnly:   so.Type == types.OrderTypeLimitMaker,
			ClientID:   newSpotClientOrderID(so.ClientOrderID),
		})

		if err != nil {
			return createdOrders, fmt.Errorf("failed to place order %+v: %w", so, err)
		}

		if !or.Success {
			return createdOrders, fmt.Errorf("ftx returns placing order failure")
		}

		globalOrder, err := toGlobalOrder(or.Result)
		if err != nil {
			return createdOrders, fmt.Errorf("failed to convert response to global order")
		}

		createdOrders = append(createdOrders, globalOrder)
	}
	return createdOrders, nil
}

func (e *Exchange) QueryOrder(ctx context.Context, q types.OrderQuery) (*types.Order, error) {
	orderID, err := strconv.ParseInt(q.OrderID, 10, 64)
	if err != nil {
		return nil, err
	}
	_ = orderID
	return nil, nil
}

func (e *Exchange) QueryOpenOrders(ctx context.Context, symbol string) (orders []types.Order, err error) {
	// TODO: invoke open trigger orders

	req := e.client.NewGetOpenOrdersRequest(toLocalSymbol(symbol))
	ftxOrders, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, ftxOrder := range ftxOrders {
		o, err := toGlobalOrderNew(ftxOrder)
		if err != nil {
			return orders, err
		}

		orders = append(orders, o)
	}
	return orders, nil
}

// symbol, since and until are all optional. FTX can only query by order created time, not updated time.
// FTX doesn't support lastOrderID, so we will query by the time range first, and filter by the lastOrderID.
func (e *Exchange) QueryClosedOrders(ctx context.Context, symbol string, since, until time.Time, lastOrderID uint64) (orders []types.Order, err error) {
	symbol = TrimUpperString(symbol)

	req := e.client.NewGetOrderHistoryRequest(toLocalSymbol(symbol))

	if since != (time.Time{}) {
		req.StartTime(since)
	} else if until != (time.Time{}) {
		req.EndTime(until)
	}

	ftxOrders, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}

	sort.Slice(ftxOrders, func(i, j int) bool {
		return ftxOrders[i].CreatedAt.Before(ftxOrders[j].CreatedAt)
	})

	for _, ftxOrder := range ftxOrders {
		o, err := toGlobalOrderNew(ftxOrder)
		if err != nil {
			return orders, err
		}

		orders = append(orders, o)
	}
	return orders, nil
}

func (e *Exchange) CancelOrders(ctx context.Context, orders ...types.Order) error {
	for _, o := range orders {
		if err := requestLimit.Wait(ctx); err != nil {
			logrus.WithError(err).Error("rate limit error")
		}

		if len(o.ClientOrderID) > 0 {
			req := e.client.NewCancelOrderByClientOrderIdRequest(o.ClientOrderID)
			_, err := req.Do(ctx)
			if err != nil {
				return err
			}
		} else {
			req := e.client.NewCancelOrderRequest(strconv.FormatUint(o.OrderID, 10))
			_, err := req.Do(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *Exchange) QueryTicker(ctx context.Context, symbol string) (*types.Ticker, error) {
	ticketMap, err := e.QueryTickers(ctx, symbol)
	if err != nil {
		return nil, err
	}

	if ticker, ok := ticketMap[symbol]; ok {
		return &ticker, nil
	}
	return nil, fmt.Errorf("ticker %s not found", symbol)
}

func (e *Exchange) QueryTickers(ctx context.Context, symbol ...string) (map[string]types.Ticker, error) {

	var tickers = make(map[string]types.Ticker)

	markets, err := e._queryMarkets(ctx)
	if err != nil {
		return nil, err
	}

	m := make(map[string]struct{})
	for _, s := range symbol {
		m[toGlobalSymbol(s)] = struct{}{}
	}

	rest := e.newRest()

	for k, v := range markets {

		// if we provide symbol as condition then we only query the gieven symbol ,
		// or we should query "ALL" symbol in the market.
		if _, ok := m[toGlobalSymbol(k)]; len(symbol) != 0 && !ok {
			continue
		}

		if err := requestLimit.Wait(ctx); err != nil {
			logrus.WithError(err).Errorf("order rate limiter wait error")
		}

		// ctx context.Context, market string, interval types.Interval, limit int64, start, end time.Time
		prices, err := rest.HistoricalPrices(ctx, v.Market.LocalSymbol, types.Interval1h, 1, time.Now().Add(time.Duration(-1)*time.Hour), time.Now())
		if err != nil || !prices.Success || len(prices.Result) == 0 {
			continue
		}

		lastCandle := prices.Result[0]
		tickers[toGlobalSymbol(k)] = types.Ticker{
			Time:   lastCandle.StartTime.Time,
			Volume: lastCandle.Volume,
			Last:   v.Last,
			Open:   lastCandle.Open,
			High:   lastCandle.High,
			Low:    lastCandle.Low,
			Buy:    v.Bid,
			Sell:   v.Ask,
		}
	}

	return tickers, nil
}

func (e *Exchange) Transfer(ctx context.Context, coin string, size float64, destination string) (string, error) {
	payload := TransferPayload{
		Coin:        coin,
		Size:        size,
		Source:      e.subAccount,
		Destination: destination,
	}
	resp, err := e.newRest().Transfer(ctx, payload)
	if err != nil {
		return "", err
	}
	if !resp.Success {
		return "", fmt.Errorf("ftx returns transfer failure")
	}
	return resp.Result.String(), nil
}
