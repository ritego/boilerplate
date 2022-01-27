package settlement

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Settlement struct {
	Id           string         `gorm:"primaryKey" json:"id"`
	Amount       float64        `gorm:"type:decimal(18, 6)" json:"amount"`
	Currency     string         `gorm:"type:varchar(3)" json:"currency"`
	UserId       string         `gorm:"type:char(36)" json:"user_id"`
	WalletId     string         `gorm:"type:char(36)" json:"wallet_id"`
	Date         time.Time      `gorm:"type:datetime" json:"date"`
	SettledAt    time.Time      `gorm:"type:datetime" json:"settled_at"`
	ClosedAt     time.Time      `gorm:"type:datetime" json:"closed_at"`
	Interest     float64        `gorm:"type:decimal(18, 6)" json:"interest"`
	InterestRate float64        `gorm:"type:decimal(18, 6)" json:"interest_rate"`
	CreatedAt    time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"type:datetime" json:"deleted_at"`
}

func (settlement *Settlement) BeforeCreate(tx *gorm.DB) error {
	settlement.Id = uuid.NewString()
	return nil
}
