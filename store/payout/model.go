package payout

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payout struct {
	Id             string         `gorm:"primaryKey" json:"id"`
	Amount         float64        `gorm:"type:decimal(18, 6)" json:"amount"`
	AmountAfterTax float64        `gorm:"type:decimal(18, 6)" json:"amount_after_tax"`
	Tax            float64        `gorm:"type:decimal(18, 6)" json:"tax"`
	Currency       string         `gorm:"type:varchar(3)" json:"currency"`
	UserId         string         `gorm:"type:char(36)" json:"user_id"`
	WalletId       string         `gorm:"type:char(36)" json:"wallet_id"`
	Date           time.Time      `gorm:"type:datetime" json:"date"`
	SettledAt      time.Time      `gorm:"type:datetime" json:"settled_at"`
	Reference      string         `gorm:"type:varchar(60)" json:"reference"`
	Status         string         `gorm:"type:char(26)" json:"status"`
	CreatedAt      time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"type:datetime" json:"deleted_at"`
}

func (payout *Payout) BeforeCreate(tx *gorm.DB) error {
	payout.Id = uuid.NewString()
	return nil
}
