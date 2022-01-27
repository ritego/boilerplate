package service

import (
	"fmt"
	"time"

	"github.com/aellacredit/jara/cache"
	"github.com/aellacredit/jara/config"
	payoutStore "github.com/aellacredit/jara/store/payout"
	settlementStore "github.com/aellacredit/jara/store/settlement"
	"github.com/aellacredit/jara/utils"
	"github.com/aellacredit/jara/utils/sec"
)

type payout struct{}

func (p *payout) Run(class string, period string, from string, to string, chunk int64) {

	//1. count
	fromTime, _ := time.Parse("2006-01-02 15:04:05", from)
	total, _ := settlementStore.CountMany(&settlementStore.Settlement{Date: fromTime})
	var offset int64 = 0
	limit := chunk

	//1. fetch transactions
	//1. skip if in cache [processing|wallet|completed|wallet-completed]
	//2. calculate payout
	//3. register payout on cache [processing]
	//4. [goroutine] send payout to wallet
	// [goroutine] register payout on cache [wallet-completed|wallet-failed]
	// [goroutine] save balance
	// [goroutine] register payout on cache [completed|failed]
	for offset < total {
		settlements, _ := settlementStore.FindMany(&settlementStore.Settlement{Date: fromTime})

		for _, s := range *settlements {
			fmt.Println(s)
			cache := cache.Default()
			status, _ := cache.Get(fmt.Sprintf(PAYOUT_STATUS_CACHE_KEY, s.UserId, s.WalletId), "")

			if !utils.InSlice(status, []string{"processing", "wallet-completed", "completed"}) {
				continue
			}

			_ = cache.Set(fmt.Sprintf(PAYOUT_STATUS_CACHE_KEY, s.UserId, s.WalletId), "processing", 0)

			go func(s settlementStore.Settlement) {
				ref := sec.GenerateReference()
				amount := s.Amount
				tax := config.Float("TAX_RATE")
				amountAfterTax := amount * (1 - tax)

				WalletClient.Credit(s.WalletId, ref, amount)

				payoutStore.Create(&payoutStore.Payout{
					Amount:         amount,
					AmountAfterTax: amountAfterTax,
					Tax:            tax,
					Currency:       s.Currency,
					UserId:         s.UserId,
					WalletId:       s.WalletId,
					Date:           s.Date,
					SettledAt:      time.Now(),
					Reference:      ref,
					Status:         "",
				})

				_ = cache.Set(fmt.Sprintf(SETTLEMENT_STATUS_CACHE_KEY, s.UserId, s.WalletId), "completed", 0)
			}(s)
		}

		offset = offset + limit
	}
}

var PayoutService = &payout{}
