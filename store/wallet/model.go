package wallet

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	Id              string         `gorm:"primaryKey" json:"id"`
	Amount          float64        `gorm:"type:decimal(18, 6)" json:"amount"`
	AvailableAmount float64        `gorm:"type:decimal(18, 6)" json:"available_amount"`
	Currency        string         `gorm:"type:varchar(3)" json:"currency"`
	Date            time.Time      `gorm:"type:datetime" json:"date"`
	UserId          string         `gorm:"type:char(36)" json:"user_id"`
	WalletId        string         `gorm:"type:char(36)" json:"wallet_id"`
	CreatedAt       time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"type:datetime" json:"deleted_at"`
}

func (wallet *Wallet) BeforeCreate(tx *gorm.DB) error {
	wallet.Id = uuid.NewString()
	return nil
}
