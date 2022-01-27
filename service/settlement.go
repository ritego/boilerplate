package service

import (
	"fmt"

	"github.com/aellacredit/jara/cache"
	"github.com/aellacredit/jara/config"
	settlementStore "github.com/aellacredit/jara/store/settlement"
	walletStore "github.com/aellacredit/jara/store/wallet"
	"github.com/aellacredit/jara/utils"
)

type settlement struct{}

func (s *settlement) Run(class string, period string, from string, to string, chunk int64) {
	switch class {
	case utils.ALL_CLASS:
		go s.hasInterestRun(period, from, to, chunk)
		go s.noInterestRun(period, from, to, chunk)
	case utils.HAS_INTEREST_CLASS:
		go s.hasInterestRun(period, from, to, chunk)
	case utils.NO_INTEREST_CLASS:
		go s.noInterestRun(period, from, to, chunk)
	}
}

func (s *settlement) hasInterestRun(period string, from string, to string, chunk int64) {

	//1. count
	total := walletStore.CountHasInterest(from, to)
	var offset int64 = 0
	limit := chunk

	//2. for each iteration per chunk
	// - fetch wallets records
	// - - for each record
	// - - - skip if its on cache register for the day [processing|completed]
	// - - - register on cache [processing]
	// - - - [goroutine] compute interest
	// - - - [goroutine] save interest to db
	// - - - [goroutine] register on cache [completed|failed]
	for offset < total {
		wallets := walletStore.FindHasInterest(from, to, offset, limit)
		rate := config.Float("INTEREST_RATE")

		for _, w := range wallets {
			fmt.Println(w.WalletId, w.UserId, w.Amount, w.AvailableAmount, w.Currency)
			cache := cache.Default()
			status, _ := cache.Get(fmt.Sprintf(SETTLEMENT_STATUS_CACHE_KEY, w.UserId, w.WalletId), "")

			if !utils.InSlice(status, []string{"processing", "completed"}) {
				continue
			}

			_ = cache.Set(fmt.Sprintf(SETTLEMENT_STATUS_CACHE_KEY, w.UserId, w.WalletId), "processing", 0)

			go func(w walletStore.Wallet, rate float64) {
				settlementStore.Create(&settlementStore.Settlement{
					Amount:       w.Amount,
					Currency:     w.Currency,
					UserId:       w.UserId,
					WalletId:     w.WalletId,
					Date:         w.Date,
					Interest:     rate * w.Amount,
					InterestRate: rate,
				})

				_ = cache.Set(fmt.Sprintf(SETTLEMENT_STATUS_CACHE_KEY, w.UserId, w.WalletId), "completed", 0)
			}(w, rate)
		}

		offset = offset + limit
	}
}

func (s *settlement) noInterestRun(period string, from string, to string, chunk int64) {

	//1. count
	total := walletStore.CountNoInterest(from, to)
	var offset int64 = 0
	limit := chunk

	//2. for each iteration per chunk
	// - fetch walletStore records
	// - - [goroutine] batch insert interest = 0 to db, ignore if exist
	for offset < total {
		walletStores := walletStore.FindNoInterest(from, to, offset, limit)
		rate := config.Float("INTEREST_RATE")

		for _, w := range walletStores {
			fmt.Println(w.WalletId, w.UserId, w.Amount, w.AvailableAmount, w.Currency)
			cache := cache.Default()
			status, _ := cache.Get(fmt.Sprintf(SETTLEMENT_STATUS_CACHE_KEY, w.UserId, w.WalletId), "")

			if !utils.InSlice(status, []string{"processing", "completed"}) {
				continue
			}

			_ = cache.Set(fmt.Sprintf(SETTLEMENT_STATUS_CACHE_KEY, w.UserId, w.WalletId), "processing", 0)

			go func(w walletStore.Wallet, rate float64) {
				settlementStore.Create(&settlementStore.Settlement{
					Amount:       w.Amount,
					Currency:     w.Currency,
					UserId:       w.UserId,
					WalletId:     w.WalletId,
					Date:         w.Date,
					Interest:     0,
					InterestRate: rate,
				})

				_ = cache.Set(fmt.Sprintf(SETTLEMENT_STATUS_CACHE_KEY, w.UserId, w.WalletId), "completed", 0)
			}(w, rate)
		}

		offset = offset + limit
	}
}

var SettlementService = &settlement{}
