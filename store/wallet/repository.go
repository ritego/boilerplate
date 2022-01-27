package wallet

import (
	"fmt"

	"github.com/aellacredit/jara/store"
)

func Create(wallet *Wallet) (*Wallet, error) {
	result := store.WalletDb.Create(wallet)
	return wallet, result.Error
}

func Save(wallet *Wallet) error {
	result := store.WalletDb.Create(wallet)
	return result.Error
}

func FindOne(id string) (*Wallet, error) {
	wallet := &Wallet{}
	result := store.WalletDb.First(wallet, "id = ?", id)
	return wallet, result.Error
}

func FindOneBy(wallet *Wallet) (*Wallet, error) {
	u := &Wallet{}
	result := store.WalletDb.Where(wallet).First(u)
	return u, result.Error
}

func FindMany(wallet *Wallet) (*[]Wallet, error) {
	wallets := &[]Wallet{}
	result := store.WalletDb.Where(wallet).Find(wallets)
	return wallets, result.Error
}

func Update(id string, wallet *Wallet) (*Wallet, error) {
	wallet.Id = id
	result := store.WalletDb.Updates(wallet)
	return wallet, result.Error
}

func Delete(id string) error {
	wallet := &Wallet{}
	wallet.Id = id
	result := store.WalletDb.Delete(&wallet)
	return result.Error
}

func CountHasInterest(from string, to string) int64 {
	q := fmt.Sprintf(`
		SELECT
			count(*) as total
		FROM
			Wallets w, Transactions t
		WHERE
			w.id = t.wallet_id
			AND t.transaction_type_id != 32
			AND t.status = 1
			AND t.status_text = 'available'
			AND t.date BETWEEN '%s' AND '%s'
			AND w.created_at <= '%s'
			AND t.amount >= 0
		GROUP BY w.id
		HAVING transactions > 0
		ORDER BY w.created_at ASC
	`, from, to, to)

	var total int64
	store.WalletDb.Raw(q).Scan(&total)

	return total
}

func FindHasInterest(from string, to string, offset int64, limit int64) []Wallet {
	q := fmt.Sprintf(`
		SELECT
			w.id as wallet_id, w.user_id, w.amount, w.available_amount, w.currency,
			count(t.id) as transactions
		FROM
			Wallets w, Transactions t
		WHERE
			w.id = t.wallet_id
			AND t.transaction_type_id != 32
			AND t.status = 1
			AND t.status_text = 'available'
			AND t.date BETWEEN '%s' AND '%s'
			AND w.created_at <= '%s'
			AND t.amount >= 0
		GROUP BY w.id
		HAVING transactions > 0
		ORDER BY w.created_at ASC
		LIMIT %d
		OFFSET %d
	`, from, to, to, limit, offset)

	var wallets []Wallet
	store.WalletDb.Raw(q).Scan(&wallets)

	return wallets
}

func CountNoInterest(from string, to string) int64 {
	q := fmt.Sprintf(`
		SELECT
			w.id as wallet_id, w.user_id, w.amount, w.available_amount, w.currency,
			count(t.id) as transactions
		FROM
			Wallets w, Transactions t
		WHERE
			w.id = t.wallet_id
			AND t.transaction_type_id != 32
			AND t.status = 1
			AND t.status_text = 'available'
			AND t.date BETWEEN '%s' AND '%s'
			AND w.created_at <= '%s'
			AND t.amount >= 0
		GROUP BY w.id
		HAVING transactions < 1
		ORDER BY w.created_at ASC
	`, from, to, to)

	var total int64
	store.WalletDb.Raw(q).Scan(&total)

	return total
}

func FindNoInterest(from string, to string, offset int64, limit int64) []Wallet {
	q := fmt.Sprintf(`
		SELECT
			w.id as wallet_id, w.user_id, w.amount, w.available_amount, w.currency,
			count(t.id) as transactions
		FROM
			Wallets w, Transactions t
		WHERE
			w.id = t.wallet_id
			AND t.transaction_type_id != 32
			AND t.status = 1
			AND t.status_text = 'available'
			AND t.date BETWEEN '%s' AND '%s'
			AND w.created_at <= '%s'
			AND t.amount >= 0
		GROUP BY w.id
		HAVING transactions < 1
		ORDER BY w.created_at ASC
		LIMIT %d
		OFFSET %d
	`, from, to, to, limit, offset)

	var wallets []Wallet
	store.WalletDb.Raw(q).Scan(&wallets)

	return wallets
}
