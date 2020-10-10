package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/adshao/go-binance"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/c9s/bbgo/accounting"
	"github.com/c9s/bbgo/bbgo"
	binance2 "github.com/c9s/bbgo/exchange/binance"
	"github.com/c9s/bbgo/service"
	"github.com/c9s/bbgo/types"
	"github.com/c9s/bbgo/util"
)

func init() {
	pnlCmd.Flags().String("symbol", "BTCUSDT", "trading symbol")
	pnlCmd.Flags().String("since", "", "pnl since time")
	RootCmd.AddCommand(pnlCmd)
}

var pnlCmd = &cobra.Command{
	Use:          "pnl",
	Short:        "pnl calculator",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		symbol, err := cmd.Flags().GetString("symbol")
		if err != nil {
			return err
		}

		binanceKey := viper.GetString("bn-key")
		binanceSecret := viper.GetString("bn-secret")

		binanceClient := binance.NewClient(binanceKey, binanceSecret)
		binanceExchange := &binance2.Exchange{Client: binanceClient}

		mysqlURL := viper.GetString("mysql-url")
		mysqlURL = fmt.Sprintf("%s?parseTime=true", mysqlURL)
		db, err := sqlx.Connect("mysql", mysqlURL)
		if err != nil {
			return err
		}

		since, err := cmd.Flags().GetString("since")
		if err != nil {
			return err
		}

		var startTime = time.Now().AddDate(-2, 0, 0)
		if len(since) > 0 {
			loc, err := time.LoadLocation("Asia/Taipei")
			if err != nil {
				return err
			}

			startTime, err = time.ParseInLocation("2006-01-02", since, loc)
			if err != nil {
				return err
			}
		}

		tradeService := &service.TradeService{DB: db}
		tradeSync := &service.TradeSync{Service: tradeService}

		logrus.Info("syncing trades...")
		if err := tradeSync.Sync(ctx, binanceExchange, symbol, startTime); err != nil {
			return err
		}

		var trades []types.Trade
		tradingFeeCurrency := binanceExchange.PlatformFeeCurrency()
		if strings.HasPrefix(symbol, tradingFeeCurrency) {
			logrus.Infof("loading all trading fee currency related trades: %s", symbol)
			trades, err = tradeService.QueryForTradingFeeCurrency(symbol, tradingFeeCurrency)
		} else {
			trades, err = tradeService.Query(symbol)
		}

		if err != nil {
			return err
		}

		logrus.Infof("%d trades loaded", len(trades))

		stockManager := &bbgo.StockDistribution{
			Symbol:             symbol,
			TradingFeeCurrency: tradingFeeCurrency,
		}

		checkpoints, err := stockManager.AddTrades(trades)
		if err != nil {
			return err
		}

		logrus.Infof("found checkpoints: %+v", checkpoints)
		logrus.Infof("stock: %f", stockManager.Stocks.Quantity())

		// query the last average price so that we can calculate the pnl
		resp, err := binanceClient.NewAveragePriceService().Symbol(symbol).Do(ctx)
		if err != nil {
			return err
		}
		currentPrice := util.MustParseFloat(resp.Price)

		calculator := &accounting.ProfitAndLossCalculator{
			TradingFeeCurrency: tradingFeeCurrency,
			Symbol:             symbol,
			StartTime:          startTime,
			CurrentPrice:       currentPrice,
			Trades:             trades,
		}
		report := calculator.Calculate()
		report.Print()
		return nil
	},
}
